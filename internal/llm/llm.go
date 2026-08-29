// Package llm is a minimal hand-written client for OpenAI-compatible Chat
// Completions APIs (OpenRouter, and any other /chat/completions endpoint). It
// uses net/http directly (no SDK) so the dependency surface stays empty and the
// build is reproducible. Model and endpoint come from the environment.
package llm

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

// Usage is the token usage and derived cost of one call.
type Usage struct {
	InputTokens  int
	OutputTokens int
	CostUSD      float64
}

// Message is one turn in a conversation.
type Message struct {
	Role    string // "user" or "assistant"
	Content string
}

// Request is a single completion request.
type Request struct {
	System      string
	Messages    []Message
	MaxTokens   int
	Temperature float64
}

// Response is the text and usage returned by the model.
type Response struct {
	Text  string
	Usage Usage
}

// Client talks to an OpenAI-compatible Chat Completions API. Construct with New.
type Client struct {
	base       string
	key        string
	model      string
	priceIn    float64 // USD per million input tokens
	priceOut   float64 // USD per million output tokens
	http       *http.Client
	maxRetries int
}

// New builds a Client from the environment: SAKSAMA_API_BASE, SAKSAMA_API_KEY,
// SAKSAMA_MODEL are required (base is the origin, e.g. https://openrouter.ai/api/v1).
// SAKSAMA_PRICE_IN and SAKSAMA_PRICE_OUT (USD per million tokens) are optional;
// without them CostUSD is reported as 0.
func New() (*Client, error) {
	base := strings.TrimRight(os.Getenv("SAKSAMA_API_BASE"), "/")
	key := os.Getenv("SAKSAMA_API_KEY")
	model := os.Getenv("SAKSAMA_MODEL")
	if base == "" || key == "" || model == "" {
		return nil, errors.New("llm: SAKSAMA_API_BASE, SAKSAMA_API_KEY, and SAKSAMA_MODEL must all be set")
	}
	return &Client{
		base:       base,
		key:        key,
		model:      model,
		priceIn:    envFloat("SAKSAMA_PRICE_IN"),
		priceOut:   envFloat("SAKSAMA_PRICE_OUT"),
		http:       &http.Client{Timeout: 180 * time.Second},
		maxRetries: 5,
	}, nil
}

// Model returns the configured model id.
func (c *Client) Model() string { return c.model }

func envFloat(k string) float64 {
	f, _ := strconv.ParseFloat(os.Getenv(k), 64)
	return f
}

// costUSD converts token counts to dollars using the configured prices.
func (c *Client) costUSD(in, out int) float64 {
	return float64(in)/1e6*c.priceIn + float64(out)/1e6*c.priceOut
}

type apiMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type apiRequest struct {
	Model       string       `json:"model"`
	MaxTokens   int          `json:"max_tokens"`
	Temperature float64      `json:"temperature"`
	Messages    []apiMessage `json:"messages"`
}

type apiResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
	} `json:"usage"`
	Error *struct {
		Message string `json:"message"`
		Code    any    `json:"code"`
	} `json:"error"`
}

// Complete sends one Chat Completions request, retrying transient failures with
// exponential backoff. It returns the assistant text and the usage/cost.
func (c *Client) Complete(ctx context.Context, req Request) (Response, error) {
	if req.MaxTokens == 0 {
		req.MaxTokens = 4096
	}
	msgs := make([]apiMessage, 0, len(req.Messages)+1)
	if req.System != "" {
		msgs = append(msgs, apiMessage{Role: "system", Content: req.System})
	}
	for _, m := range req.Messages {
		msgs = append(msgs, apiMessage{Role: m.Role, Content: m.Content})
	}
	body, err := json.Marshal(apiRequest{
		Model:       c.model,
		MaxTokens:   req.MaxTokens,
		Temperature: req.Temperature,
		Messages:    msgs,
	})
	if err != nil {
		return Response{}, fmt.Errorf("llm: marshal: %w", err)
	}

	url := c.base + "/chat/completions"
	var lastErr error
	for attempt := 0; attempt <= c.maxRetries; attempt++ {
		if attempt > 0 {
			backoff := time.Duration(1<<uint(attempt-1)) * time.Second
			select {
			case <-ctx.Done():
				return Response{}, ctx.Err()
			case <-time.After(backoff):
			}
		}

		httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(body))
		if err != nil {
			return Response{}, fmt.Errorf("llm: new request: %w", err)
		}
		httpReq.Header.Set("content-type", "application/json")
		httpReq.Header.Set("authorization", "Bearer "+c.key)
		// Optional OpenRouter attribution headers (harmless elsewhere).
		httpReq.Header.Set("HTTP-Referer", "https://github.com/EndPx/saksama")
		httpReq.Header.Set("X-Title", "Saksama")

		resp, err := c.http.Do(httpReq)
		if err != nil {
			lastErr = err
			continue // network error: retry
		}
		respBody, _ := io.ReadAll(resp.Body)
		resp.Body.Close()

		if resp.StatusCode == http.StatusTooManyRequests || resp.StatusCode >= 500 {
			lastErr = fmt.Errorf("llm: http %d: %s", resp.StatusCode, truncate(respBody))
			continue // transient: retry
		}
		if resp.StatusCode != http.StatusOK {
			return Response{}, fmt.Errorf("llm: http %d: %s", resp.StatusCode, truncate(respBody))
		}

		var ar apiResponse
		if err := json.Unmarshal(respBody, &ar); err != nil {
			return Response{}, fmt.Errorf("llm: decode: %w (body: %s)", err, truncate(respBody))
		}
		if ar.Error != nil {
			// Free models often return a rate-limit inside a 200 body; retry those.
			if isRateLimit(ar.Error.Code, ar.Error.Message) {
				lastErr = fmt.Errorf("llm: rate limited: %s", ar.Error.Message)
				continue
			}
			return Response{}, fmt.Errorf("llm: api error: %s", ar.Error.Message)
		}
		if len(ar.Choices) == 0 {
			return Response{}, fmt.Errorf("llm: no choices in response: %s", truncate(respBody))
		}
		return Response{
			Text: ar.Choices[0].Message.Content,
			Usage: Usage{
				InputTokens:  ar.Usage.PromptTokens,
				OutputTokens: ar.Usage.CompletionTokens,
				CostUSD:      c.costUSD(ar.Usage.PromptTokens, ar.Usage.CompletionTokens),
			},
		}, nil
	}
	return Response{}, fmt.Errorf("llm: exhausted retries: %w", lastErr)
}

func truncate(b []byte) string {
	const max = 500
	if len(b) > max {
		return string(b[:max]) + "..."
	}
	return string(b)
}

// isRateLimit reports whether an in-body error looks like a rate limit (HTTP
// 429), so the caller can retry instead of failing hard.
func isRateLimit(code any, msg string) bool {
	switch v := code.(type) {
	case float64:
		if int(v) == 429 {
			return true
		}
	case string:
		if v == "429" || strings.Contains(v, "429") {
			return true
		}
	}
	m := strings.ToLower(msg)
	return strings.Contains(m, "rate-limit") || strings.Contains(m, "rate limit")
}
