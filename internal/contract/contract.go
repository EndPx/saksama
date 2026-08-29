// Package contract splits an Indonesian employment contract (markdown) into its
// numbered articles (Pasal), so a stage can reason about one section at a time
// instead of the whole document.
package contract

import (
	"regexp"
	"strings"
)

// Section is one article of a contract.
type Section struct {
	Number  string // e.g. "4"
	Title   string // e.g. "Masa Percobaan"
	Heading string // the raw heading line
	Body    string // text until the next article heading
}

// Label returns the canonical citation label, e.g. "Pasal 4".
func (s Section) Label() string { return "Pasal " + s.Number }

// headingRe matches a markdown "Pasal N" article heading and captures the
// number and the (optional) title after a dash or colon separator.
var headingRe = regexp.MustCompile(`(?mi)^#{1,6}\s*Pasal\s+(\d+)\s*[—\-:]?\s*(.*?)\s*$`)

// Parse splits md into the text before the first article (the preamble, which
// carries the parties' identities) and the ordered list of articles.
func Parse(md string) (preamble string, sections []Section) {
	locs := headingRe.FindAllStringSubmatchIndex(md, -1)
	if len(locs) == 0 {
		return strings.TrimSpace(md), nil
	}
	preamble = strings.TrimSpace(md[:locs[0][0]])
	for i, loc := range locs {
		number := md[loc[2]:loc[3]]
		title := strings.TrimSpace(md[loc[4]:loc[5]])
		heading := strings.TrimSpace(md[loc[0]:loc[1]])
		bodyStart := loc[1]
		bodyEnd := len(md)
		if i+1 < len(locs) {
			bodyEnd = locs[i+1][0]
		}
		sections = append(sections, Section{
			Number:  number,
			Title:   title,
			Heading: heading,
			Body:    strings.TrimSpace(md[bodyStart:bodyEnd]),
		})
	}
	return preamble, sections
}

// wsRe collapses any run of whitespace to a single space.
var wsRe = regexp.MustCompile(`\s+`)

// NormalizeWhitespace collapses all whitespace runs to single spaces and trims.
// The citation gate uses it so a quotation matches regardless of line wrapping.
func NormalizeWhitespace(s string) string {
	return strings.TrimSpace(wsRe.ReplaceAllString(s, " "))
}

// ContainsQuote reports whether quote appears in text after both are
// whitespace-normalised. Used by the S5 citation gate.
func ContainsQuote(text, quote string) bool {
	q := NormalizeWhitespace(quote)
	if q == "" {
		return false
	}
	return strings.Contains(NormalizeWhitespace(text), q)
}
