package anthropic

import "testing"

func TestEncodeProjectPath_unix(t *testing.T) {
	got := EncodeProjectPath("/Users/foo/bar")
	want := "-Users-foo-bar"
	if got != want {
		t.Fatalf("got %q want %q", got, want)
	}
}
