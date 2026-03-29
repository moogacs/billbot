package model

// Vendor identifies the upstream API family used for pricing rules.
type Vendor string

const (
	VendorAnthropic Vendor = "anthropic"
	VendorOpenAI    Vendor = "openai"
	VendorCursor    Vendor = "cursor"
)

// SessionRef points at one analyzable session on disk.
type SessionRef struct {
	Vendor Vendor `json:"vendor"`
	Path   string `json:"path"`
	ID     string `json:"id,omitempty"`
}

// Money is a USD amount for estimates only (not an invoice).
type Money struct {
	USD float64 `json:"usd"`
}

// UsageBreakdown is token counts for one API completion (or aggregated).
type UsageBreakdown struct {
	InputTokens              int `json:"input_tokens"`
	OutputTokens             int `json:"output_tokens"`
	CacheReadInputTokens     int `json:"cache_read_input_tokens,omitempty"`
	CacheCreationInputTokens int `json:"cache_creation_input_tokens,omitempty"`
}

// Add returns the component-wise sum of u and o.
func (u UsageBreakdown) Add(o UsageBreakdown) UsageBreakdown {
	return UsageBreakdown{
		InputTokens:              u.InputTokens + o.InputTokens,
		OutputTokens:             u.OutputTokens + o.OutputTokens,
		CacheReadInputTokens:     u.CacheReadInputTokens + o.CacheReadInputTokens,
		CacheCreationInputTokens: u.CacheCreationInputTokens + o.CacheCreationInputTokens,
	}
}

// UserPrompt is a top-level human question in the conversation.
type UserPrompt struct {
	Vendor    Vendor `json:"vendor"`
	Text      string `json:"text"`
	Timestamp string `json:"timestamp,omitempty"`
	SourceRef string `json:"source_ref,omitempty"`
}

// AssistantCompletion is one model response (one API completion after dedup).
type AssistantCompletion struct {
	Vendor    Vendor         `json:"vendor"`
	Model     string         `json:"model"`
	Timestamp string         `json:"timestamp,omitempty"`
	DedupKey  string         `json:"dedup_key,omitempty"`
	Usage     UsageBreakdown `json:"usage"`
	SourceRef string         `json:"source_ref,omitempty"`
}

// Turn groups one user prompt with all assistant completions until the next prompt.
type Turn struct {
	Index       int
	Prompt      UserPrompt
	Completions []AssistantCompletion
}

// NormalizedEvent is a vendor-neutral timeline event for turn building.
type NormalizedEvent interface {
	isNormalizedEvent()
}

type EventUser struct {
	Prompt UserPrompt
}

func (EventUser) isNormalizedEvent() {}

type EventAssistant struct {
	Completion AssistantCompletion
}

func (EventAssistant) isNormalizedEvent() {}

// AnalyzeReport is JSON output for --format json.
type AnalyzeReport struct {
	Meta    ReportMeta     `json:"meta"`
	Turns   []TurnReport   `json:"turns"`
	Answers []AnswerReport `json:"answers"`
}

// ReportMeta carries run metadata and disclaimers.
type ReportMeta struct {
	Vendor      Vendor `json:"vendor"`
	SessionPath string `json:"session_path"`
	Disclaimer  string `json:"disclaimer"`
}

// TurnReport is one user question with rollups and nested answers.
type TurnReport struct {
	Index           int            `json:"index"`
	Prompt          string         `json:"prompt"`
	PromptTimestamp string         `json:"prompt_timestamp,omitempty"`
	Usage           UsageBreakdown `json:"usage"`
	CostUSD         float64        `json:"cost_usd"`
	Answers         []AnswerReport `json:"answers"`
}

// AnswerReport is one assistant/API completion row.
type AnswerReport struct {
	TurnIndex    int            `json:"turn_index"`
	Model        string         `json:"model"`
	Timestamp    string         `json:"timestamp,omitempty"`
	DedupKey     string         `json:"dedup_key,omitempty"`
	Usage        UsageBreakdown `json:"usage"`
	CostUSD      float64        `json:"cost_usd"`
	UnknownModel bool           `json:"unknown_model"`
	SourceRef    string         `json:"source_ref,omitempty"`
}

// AggregateByProviderReport is JSON for directory mode (all sessions, split by vendor).
type AggregateByProviderReport struct {
	Meta struct {
		Disclaimer string `json:"disclaimer"`
	} `json:"meta"`
	Providers   []ProviderAggregateSection `json:"providers"`
	GrandTotals AggregateTotalsRow         `json:"grand_totals"`
}

// ProviderAggregateSection is one vendor's sessions plus subtotals.
type ProviderAggregateSection struct {
	Vendor       Vendor                `json:"vendor"`
	SessionFiles int                   `json:"session_files"`
	Sessions     []AggregateSessionRow `json:"sessions"`
	Totals       AggregateTotalsRow    `json:"totals"`
}

// AggregateSessionRow is one analyzed session file in an aggregate run.
type AggregateSessionRow struct {
	SessionPath         string         `json:"session_path"`
	FirstPrompt         string         `json:"first_prompt,omitempty"`
	CostUSD             float64        `json:"cost_usd"`
	Usage               UsageBreakdown `json:"usage"`
	Turns               int            `json:"turns"`
	Answers             int            `json:"answers"`
	UnknownModelAnswers int            `json:"unknown_model_answers"`
}

// AggregateTotalsRow sums all sessions in an aggregate run.
type AggregateTotalsRow struct {
	CostUSD             float64        `json:"cost_usd"`
	Usage               UsageBreakdown `json:"usage"`
	Turns               int            `json:"turns"`
	Answers             int            `json:"answers"`
	UnknownModelAnswers int            `json:"unknown_model_answers"`
}
