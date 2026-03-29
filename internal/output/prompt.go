package output

import (
	"strings"
	"unicode/utf8"
)

// FormatPromptCell turns a prompt into a single-line table cell, truncated by runes.
func FormatPromptCell(text string, maxRunes int) string {
	s := strings.TrimSpace(text)
	s = strings.ReplaceAll(s, "\n", " ")
	s = strings.ReplaceAll(s, "\r", " ")
	s = strings.ReplaceAll(s, "\t", " ")
	for strings.Contains(s, "  ") {
		s = strings.ReplaceAll(s, "  ", " ")
	}
	if maxRunes <= 0 || utf8.RuneCountInString(s) <= maxRunes {
		return s
	}
	r := []rune(s)
	return string(r[:maxRunes]) + "…"
}
