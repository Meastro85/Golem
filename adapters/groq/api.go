package groq

// apiRequest is the shape of the request body sent to the API
type apiRequest struct {
	Model     string    `json:"model"`
	MaxTokens int       `json:"max_tokens"`
	Tools     []AITool  `json:"tools,omitempty"`
	Messages  []Message `json:"messages"`
}

type apiResponse struct {
	Role      string     `json:"role"`
	ToolCalls []toolCall `json:"tool_calls"`
}
