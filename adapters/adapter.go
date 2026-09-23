package adapters

import (
	"context"
	"fmt"
	"os"
)

type Adapter interface {
	Send(ctx context.Context, req Request) (*Response, error)
}

// CreateAdapter selects and configures an Adapter from environment variables:
//
//	AI_PROVIDER         "anthropic" or "groq" (required)
//	AI_MODEL            model name for the selected provider (required)
//	ANTHROPIC_API_KEY   required when AI_PROVIDER=anthropic
//	GROQ_API_KEY        required when AI_PROVIDER=groq
func CreateAdapter() (Adapter, error) {
	model := os.Getenv("AI_MODEL")
	if model == "" {
		return nil, fmt.Errorf("adapters: AI_MODEL is required")
	}

	switch provider := os.Getenv("AI_PROVIDER"); provider {
	case "anthropic":
		apiKey := os.Getenv("ANTHROPIC_API_KEY")
		if apiKey == "" {
			return nil, fmt.Errorf("adapters: ANTHROPIC_API_KEY is required when AI_PROVIDER=anthropic")
		}
		return nil, nil //newAnthropicAdapter(apiKey, model), nil

	case "groq":
		apiKey := os.Getenv("GROQ_API_KEY")
		if apiKey == "" {
			return nil, fmt.Errorf("adapters: GROQ_API_KEY is required when AI_PROVIDER=groq")
		}
		return newGroqAdapter(apiKey, model), nil

	case "":
		return nil, fmt.Errorf(`adapters: AI_PROVIDER is required (e.g. "anthropic" or "groq")`)

	default:
		return nil, fmt.Errorf("adapters: unknown AI_PROVIDER %q (supported: \"anthropic\", \"groq\")", provider)
	}
}
