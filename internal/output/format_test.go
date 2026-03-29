package output

import "testing"

func TestFormatIntHuman(t *testing.T) {
	tests := []struct {
		n    int
		want string
	}{
		{0, "0"},
		{12, "12"},
		{1234, "1,234"},
		{-1234, "-1,234"},
		{1000000, "1,000,000"},
	}
	for _, tc := range tests {
		if got := formatIntHuman(tc.n); got != tc.want {
			t.Fatalf("formatIntHuman(%d) = %q want %q", tc.n, got, tc.want)
		}
	}
}
