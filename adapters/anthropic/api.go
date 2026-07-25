package anthropic

// apiRequest is the shape of the request body sent to the API
type apiRequest struct {
	Model     string    `json:"model"`
	MaxTokens int       `json:"max_tokens"`
	Tools     []AITool  `json:"tools,omitempty"`
	Messages  []Message `json:"messages"`
}

// apiResponse is the shape of the response body returned by the API
type apiResponse struct {
	ID         string         `json:"id"`
	StopReason string         `json:"stop_reason"`
	Content    []ContentBlock `json:"content"`
}
