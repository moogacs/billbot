package output

import "testing"

func TestFormatPromptCell(t *testing.T) {
	if got := FormatPromptCell("  hello\nworld\t", 100); got != "hello world" {
		t.Fatalf("got %q", got)
	}
	long := "abcdefghijklmnop"
	if got := FormatPromptCell(long, 8); got != "abcdefgh…" {
		t.Fatalf("got %q", got)
	}
}
