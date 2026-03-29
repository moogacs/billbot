package output

import (
	"fmt"
	"io"
	"strings"
)

const csi = "\033["

func (o DisplayOptions) wrap(seq, s string) string {
	if !o.Color {
		return s
	}
	return csi + seq + "m" + s + csi + "0m"
}

// Dim renders faint text (disclaimers, rules, labels).
func (o DisplayOptions) Dim(s string) string { return o.wrap("2", s) }

// Bold renders strong emphasis.
func (o DisplayOptions) Bold(s string) string { return o.wrap("1", s) }

// Cyan, BrightCyan, Green, Yellow, Magenta, Blue are accent colors.
func (o DisplayOptions) Cyan(s string) string { return o.wrap("36", s) }

func (o DisplayOptions) BrightCyan(s string) string { return o.wrap("1;96", s) }

func (o DisplayOptions) Green(s string) string { return o.wrap("32", s) }

func (o DisplayOptions) Yellow(s string) string { return o.wrap("33", s) }

func (o DisplayOptions) Magenta(s string) string { return o.wrap("35", s) }

func (o DisplayOptions) Blue(s string) string { return o.wrap("94", s) }

// MoneyUSD styles estimated dollar amounts.
func (o DisplayOptions) MoneyUSD(usd float64) string {
	s := formatUSD(usd)
	if !o.Color {
		return s
	}
	if usd == 0 {
		return o.Dim(s)
	}
	return o.Green(s)
}

// PrintSectionBanner prints a titled block divider.
func (o DisplayOptions) PrintSectionBanner(w io.Writer, width int, title string) {
	if width < 20 {
		width = 88
	}
	line := strings.Repeat("=", width)
	fmt.Fprintln(w, o.Dim(line))
	fmt.Fprintf(w, "  %s\n", o.BrightCyan(o.Bold(title)))
	fmt.Fprintln(w, o.Dim(line))
	fmt.Fprintln(w)
}
