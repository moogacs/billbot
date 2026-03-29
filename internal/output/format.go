package output

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

// formatIntHuman renders n with thousands separators.
func formatIntHuman(n int) string {
	neg := n < 0
	if neg {
		n = -n
	}
	s := strconv.FormatInt(int64(n), 10)
	if len(s) <= 3 {
		if neg {
			return "-" + s
		}
		return s
	}
	var b strings.Builder
	lead := len(s) % 3
	if lead == 0 {
		lead = 3
	}
	b.WriteString(s[:lead])
	for i := lead; i < len(s); i += 3 {
		b.WriteByte(',')
		b.WriteString(s[i : i+3])
	}
	out := b.String()
	if neg {
		return "-" + out
	}
	return out
}

// formatUSD picks a sensible precision for terminal display.
func formatUSD(usd float64) string {
	switch {
	case usd == 0:
		return "$0.00"
	case usd < 0.0001:
		return fmt.Sprintf("$%.6f", usd)
	case usd < 0.01:
		return fmt.Sprintf("$%.5f", usd)
	case usd < 1:
		return fmt.Sprintf("$%.4f", usd)
	case usd < 1000:
		return fmt.Sprintf("$%.2f", usd)
	default:
		return fmt.Sprintf("$%.2f", usd)
	}
}

func truncatePath(s string, maxRunes int) string {
	if maxRunes < 4 {
		return s
	}
	if utf8.RuneCountInString(s) <= maxRunes {
		return s
	}
	r := []rune(s)
	return "…" + string(r[len(r)-maxRunes+1:])
}
