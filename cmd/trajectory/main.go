// Command trajectory converts Claude Code session JSONL files into readable
// markdown, so the agent's own working process is part of the submission.
package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func main() {
	home, _ := os.UserHomeDir()
	projects := flag.String("projects", filepath.Join(home, ".claude", "projects"), "Claude Code projects dir")
	out := flag.String("out", "trajectories", "output directory")
	limit := flag.Int("limit", 5, "max most-recent sessions to convert")
	flag.Parse()

	files, err := findSessions(*projects)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
	if len(files) > *limit {
		files = files[:*limit]
	}
	if len(files) == 0 {
		fmt.Fprintln(os.Stderr, "no session files found under", *projects)
		return
	}
	for _, f := range files {
		md, err := convert(f)
		if err != nil {
			fmt.Fprintf(os.Stderr, "skip %s: %v\n", f, err)
			continue
		}
		name := "session_" + strings.TrimSuffix(filepath.Base(f), ".jsonl") + ".md"
		dst := filepath.Join(*out, name)
		if err := os.WriteFile(dst, []byte(md), 0o644); err != nil {
			fmt.Fprintf(os.Stderr, "write %s: %v\n", dst, err)
			continue
		}
		fmt.Fprintf(os.Stderr, "wrote %s\n", dst)
	}
}

// findSessions returns *.jsonl paths, most recently modified first.
func findSessions(dir string) ([]string, error) {
	var files []string
	err := filepath.Walk(dir, func(p string, info os.FileInfo, err error) error {
		if err != nil {
			return nil // ignore unreadable entries
		}
		if !info.IsDir() && strings.HasSuffix(p, ".jsonl") {
			files = append(files, p)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	sort.Slice(files, func(i, j int) bool {
		fi, _ := os.Stat(files[i])
		fj, _ := os.Stat(files[j])
		return fi.ModTime().After(fj.ModTime())
	})
	return files, nil
}

type entry struct {
	Type    string `json:"type"`
	Message *struct {
		Role    string          `json:"role"`
		Content json.RawMessage `json:"content"`
	} `json:"message"`
}

type block struct {
	Type    string          `json:"type"`
	Text    string          `json:"text"`
	Name    string          `json:"name"`
	Input   json.RawMessage `json:"input"`
	Content json.RawMessage `json:"content"`
}

func convert(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	var b strings.Builder
	fmt.Fprintf(&b, "# Trajectory — %s\n\n", filepath.Base(path))

	sc := bufio.NewScanner(f)
	sc.Buffer(make([]byte, 1024*1024), 8*1024*1024)
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if line == "" {
			continue
		}
		var e entry
		if err := json.Unmarshal([]byte(line), &e); err != nil {
			continue
		}
		if e.Message == nil {
			continue
		}
		role := e.Message.Role
		if role == "" {
			role = e.Type
		}
		fmt.Fprintf(&b, "## %s\n\n", titleCase(role))
		writeContent(&b, e.Message.Content)
		b.WriteString("\n")
	}
	return b.String(), sc.Err()
}

func titleCase(s string) string {
	if s == "" {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

func writeContent(b *strings.Builder, raw json.RawMessage) {
	if len(raw) == 0 {
		return
	}
	// content may be a plain string...
	var s string
	if json.Unmarshal(raw, &s) == nil {
		b.WriteString(strings.TrimSpace(s) + "\n")
		return
	}
	// ...or an array of typed blocks.
	var blocks []block
	if json.Unmarshal(raw, &blocks) != nil {
		return
	}
	for _, blk := range blocks {
		switch blk.Type {
		case "text":
			if t := strings.TrimSpace(blk.Text); t != "" {
				b.WriteString(t + "\n\n")
			}
		case "tool_use":
			fmt.Fprintf(b, "**→ tool call: %s**\n\n```json\n%s\n```\n\n", blk.Name, compact(blk.Input))
		case "tool_result":
			fmt.Fprintf(b, "**← tool result**\n\n```\n%s\n```\n\n", truncate(string(blk.Content), 2000))
		}
	}
}

func compact(raw json.RawMessage) string {
	if len(raw) == 0 {
		return "{}"
	}
	var buf bytes.Buffer
	if err := json.Indent(&buf, raw, "", "  "); err != nil {
		return string(raw)
	}
	return truncate(buf.String(), 2000)
}

func truncate(s string, max int) string {
	s = strings.TrimSpace(s)
	if len(s) > max {
		return s[:max] + "\n... (truncated)"
	}
	return s
}
