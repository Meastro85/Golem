package adapters

import (
	"Golem/adapters/groq"
	"context"
	"encoding/json"
)

type groqAdapter struct {
	apiKey string
	model  string
}

func newGroqAdapter(apiKey, model string) *groqAdapter {
	return &groqAdapter{apiKey: apiKey, model: model}
}

func (a *groqAdapter) Send(ctx context.Context, req Request) (*Response, error) {
	tools := make([]groq.AITool, len(req.Tools))
	for i, t := range req.Tools {
		var params *groq.Parameters
		err := json.Unmarshal(t.Parameters, params)

		if err != nil {

			continue
		}

		tools[i] = groq.AITool{
			Type: "function",
			Function: groq.Function{
				Name:        t.Name,
				Description: t.Description,
				Parameters:  *params,
			},
		}
	}

	return nil, nil
}
