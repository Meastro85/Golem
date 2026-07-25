package anthropic

import (
	"Golem/tool"
	"encoding/json"
)

// AITool is the shape the API expects in the "tools" array
type AITool struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	InputSchema any    `json:"input_schema"`
}

// Message is a single turn. Content is either plain string or []ContentBlock.
type Message struct {
	Roles   string `json:"roles"`
	Content any    `json:"content"`
}

// ContentBlock covers the block types the loop needs to read or write:
// text, tool_use (from Claude), and tool_result (sent back to Claude)
type ContentBlock struct {
	Type string `json:"type"`
	Text string `json:"text,omitempty"`

	// tool_use
	ID    string          `json:"id,omitempty"`
	Name  string          `json:"name,omitempty"`
	Input json.RawMessage `json:"input,omitempty"`

	// tool_result
	ToolUseID string `json:"tool_use_id,omitempty"`
	Content   string `json:"content,omitempty"`
	IsError   bool   `json:"is_error,omitempty"`
}

func ToolsFor(tools []tool.Tool) []AITool {
	aiTools := make([]AITool, len(tools))
	for i, t := range tools {
		aiTools[i] = AITool{
			Name:        t.Name,
			Description: t.Description,
			InputSchema: t.Schema,
		}
	}
	return aiTools
}
