package tool

import (
	"fmt"
)

// Registry is a collection of tools that can be registered and retrieved by name. It provides methods to register new tools and retrieve existing ones.
type Registry struct {
	tools map[string]*Tool
}

func NewRegistry() *Registry {
	return &Registry{
		tools: make(map[string]*Tool),
	}
}

// Register adds a new tool to the registry. It takes the name, description, and handler function of the tool, constructs a Tool object, and stores it in the registry. If the tool cannot be constructed, it returns an error.
func (r *Registry) Register(name, description string, fn any) error {
	f, err := BuildTool(name, description, fn)
	if err != nil {
		return fmt.Errorf("failed to register tool %s: %w", name, err)
	}
	r.tools[name] = f
	return nil
}

// List returns a slice of all registered tools in the registry.
func (r *Registry) List() []*Tool {
	out := make([]*Tool, 0, len(r.tools))
	for _, tool := range r.tools {
		out = append(out, tool)
	}
	return out
}

// Get retrieves a tool from the registry by its name. It returns the tool and a boolean indicating whether the tool was found in the registry.
func (r *Registry) Get(name string) (*Tool, bool) {
	tool, ok := r.tools[name]
	return tool, ok
}
