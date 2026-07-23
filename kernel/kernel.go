package kernel

import (
	"Golem/tool"
	"context"
	"encoding/json"
	"fmt"
	"reflect"
)

// Kernel is the main struct that represents the kernel. It holds a registry of tools and provides methods to register new tools and retrieve existing ones.
// It serves as the central point for managing and executing tools and interaction with the LLM.
type Kernel struct {
	registry *tool.Registry
}

func NewKernel() *Kernel {
	return &Kernel{
		registry: tool.NewRegistry(),
	}
}

// RegisterTool registers a new tool with the kernel. It takes the name, description, and handler function of the tool, constructs a Tool object, and stores it in the kernel's registry. If the tool cannot be constructed, it returns an error.
func (k *Kernel) RegisterTool(name, description string, fn any) error {
	return k.registry.Register(name, description, fn)
}

// Execute executes a registered tool by its name with the provided raw JSON arguments. It retrieves the tool from the registry, decodes the arguments into the expected struct type, and calls the tool's handler function. It returns the result of the execution or an error if the tool is not found or if there is an issue with argument decoding or execution.
func (k *Kernel) Execute(ctx context.Context, name string, rawArgs json.RawMessage) (any, error) {
	tool, ok := k.registry.Get(name)

	if !ok {
		return nil, fmt.Errorf("kernel: no function registered with name %q", name)
	}

	argsPtr := reflect.New(tool.ArgsType)
	if len(rawArgs) > 0 {
		if err := json.Unmarshal(rawArgs, argsPtr.Interface()); err != nil {
			return nil, fmt.Errorf("kernel: decoding args for %q: %w", name, err)
		}
	}

	results := tool.Handler.Call([]reflect.Value{
		reflect.ValueOf(ctx),
		argsPtr.Elem(),
	})

	if errVal := results[1].Interface(); errVal != nil {
		return nil, errVal.(error)
	}
	return results[0].Interface(), nil
}
