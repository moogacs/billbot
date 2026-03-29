package output

// DisplayOptions controls optional columns, prompt truncation, and ANSI colors in text tables.
type DisplayOptions struct {
	// ShowPrompts adds prompt text to the per-turn table, per-completion table (turn prompt), and aggregate rows (first prompt in session).
	ShowPrompts bool
	// PromptMaxRunes is the maximum runes per prompt cell; if <= 0, 72 is used.
	PromptMaxRunes int
	// Color enables ANSI colors when true (TTY / --color always; off for pipes, NO_COLOR, or --no-color).
	Color bool
}

// PromptLimit returns the effective max runes for prompt cells.
func (o DisplayOptions) PromptLimit() int {
	if o.PromptMaxRunes <= 0 {
		return 72
	}
	return o.PromptMaxRunes
}
