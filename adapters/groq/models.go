package groq

// AITool is the shape the API expects in the "tools" array
// Valid type defitions are: "function", "browser_search", "code_interpreter"
type AITool struct {
	Type     string   `json:"type"`
	Function Function `json:"function,omitempty"`
}

type Message struct {
	Roles   string `json:"roles"`
	Content string `json:"content"`
}

// Function is the definition of the self defined function for the AI
type Function struct {
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Parameters  Parameters `json:"parameters"`
}

// Parameters is a JSON Schema object describing a function's arguments.
type Parameters struct {
	Type       string              `json:"type"`
	Properties map[string]Property `json:"properties"`
	Required   []string            `json:"required,omitempty"`
}

// Property describes a single parameter.
type Property struct {
	Type        string   `json:"type"`
	Description string   `json:"description,omitempty"`
	Enum        []string `json:"required,omitempty"`
}

type toolCall struct {
	Id       string           `json:"id"`
	Type     string           `json:"type"`
	Function toolCallFunction `json:"function"`
}

type toolCallFunction struct {
	Name      string `json:"name"`
	Arguments string `json:"arguments"`
}
