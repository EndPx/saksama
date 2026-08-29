// Package llm is a minimal hand-written client for the Anthropic Messages API.
// It uses net/http directly (no SDK) so the dependency surface stays empty and
// the build is reproducible. Model and endpoint come from the environment.
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

// Client talks to the Messages API. Construct it with New.
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
// SAKSAMA_MODEL are required. SAKSAMA_PRICE_IN and SAKSAMA_PRICE_OUT (USD per
// million tokens) are optional; without them CostUSD is reported as 0.
func New() (*Client, error) {
	base := os.Getenv("SAKSAMA_API_BASE")
	key := os.Getenv("SAKSAMA_API_KEY")
	model := os.Getenv("SAKSAMA_MODEL")
	if base == "" || key == "" || model == "" {
		return nil, errors.New("llm: SAKSAMA_API_BASE, SAKSAMA_API_KEY, and SAKSAMA_MODEL must all be set")
	}
	c := &Client{
		base:       base,
		key:        key,
		model:      model,
		priceIn:    envFloat("SAKSAMA_PRICE_IN"),
		priceOut:   envFloat("SAKSAMA_PRICE_OUT"),
		http:       &http.Client{Timeout: 120 * time.Second},
		maxRetries: 5,
	}
	return c, nil
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
	System      string       `json:"system,omitempty"`
	Messages    []apiMessage `json:"messages"`
}

type apiResponse struct {
	Content []struct {
		Type string `json:"type"`
		Text string `json:"text"`
	} `json:"content"`
	Usage struct {
		InputTokens  int `json:"input_tokens"`
		OutputTokens int `json:"output_tokens"`
	} `json:"usage"`
	Error *struct {
		Type    string `json:"type"`
		Message string `json:"message"`
	} `json:"error"`
}

// Complete sends one request, retrying transient failures with exponential
// backoff. It returns the concatenated text blocks and the usage/cost.
func (c *Client) Complete(ctx context.Context, req Request) (Response, error) {
	if req.MaxTokens == 0 {
		req.MaxTokens = 4096
	}
	body, err := json.Marshal(apiRequest{
		Model:       c.model,
		MaxTokens:   req.MaxTokens,
		Temperature: req.Temperature,
		System:      req.System,
		Messages:    toAPIMessages(req.Messages),
	})
	if err != nil {
		return Response{}, fmt.Errorf("llm: marshal: %w", err)
	}

	url := c.base + "/v1/messages"
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
		httpReq.Header.Set("x-api-key", c.key)
		httpReq.Header.Set("anthropic-version", "2023-06-01")

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
			return Response{}, fmt.Errorf("llm: decode: %w", err)
		}
		if ar.Error != nil {
			return Response{}, fmt.Errorf("llm: api error %s: %s", ar.Error.Type, ar.Error.Message)
		}
		var text string
		for _, blk := range ar.Content {
			if blk.Type == "text" {
				text += blk.Text
			}
		}
		return Response{
			Text: text,
			Usage: Usage{
				InputTokens:  ar.Usage.InputTokens,
				OutputTokens: ar.Usage.OutputTokens,
				CostUSD:      c.costUSD(ar.Usage.InputTokens, ar.Usage.OutputTokens),
			},
		}, nil
	}
	return Response{}, fmt.Errorf("llm: exhausted retries: %w", lastErr)
}

func toAPIMessages(ms []Message) []apiMessage {
	out := make([]apiMessage, len(ms))
	for i, m := range ms {
		out[i] = apiMessage{Role: m.Role, Content: m.Content}
	}
	return out
}

func truncate(b []byte) string {
	const max = 500
	if len(b) > max {
		return string(b[:max]) + "..."
	}
	return string(b)
}
