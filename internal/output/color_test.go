package output

import "testing"

func TestDisplayOptions_noColorPassthrough(t *testing.T) {
	o := DisplayOptions{Color: false}
	if got := o.Dim("hello"); got != "hello" {
		t.Fatalf("Dim: %q", got)
	}
	if got := o.MoneyUSD(1.5); got != formatUSD(1.5) {
		t.Fatalf("MoneyUSD: %q", got)
	}
}

func TestDisplayOptions_colorWraps(t *testing.T) {
	o := DisplayOptions{Color: true}
	if got := o.Dim("x"); got == "x" || len(got) < len("x") {
		t.Fatalf("expected ansi wrap, got %q", got)
	}
}
