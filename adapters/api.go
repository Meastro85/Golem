package adapters

import "encoding/json"

// Request is a canonical completion request.
type Request struct {
	Messages  []Message
	Tools     []Tool
	MaxTokens int
}

// Response is what any adapter returns after a completion.
type Response struct {
	Text       string
	ToolCalls  []ToolCall
	StopReason StopReason
}

// Message is a canonical chat message, independent of any provider's wire format.
type Message struct {
	Role    string // "user", "assistant", "system"
	Content string
}

// Tool is a canonical function description handed to any provider.
type Tool struct {
	Name        string
	Description string
	Parameters  json.RawMessage
}

// ToolCall is a single tool invocation the model requested.
type ToolCall struct {
	ID        string
	Name      string
	Arguments json.RawMessage
}

// StopReason tells why the model stopped running.
type StopReason int

const (
	StopEndTurn StopReason = iota
	StopToolUse
)
