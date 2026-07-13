package tool

import (
	"context"
	"fmt"
	"reflect"
)

var (
	ctxType = reflect.TypeFor[context.Context]()
	errType = reflect.TypeFor[error]()
)

// Tool is the "command" the kernel exposes: a name, a description
// and a schema for the parameters it accepts, as well as a handler function
type Tool struct {
	Name        string
	Description string
	Schema      ParamSchema

	Handler  reflect.Value
	ArgsType reflect.Type
}

// BuildTool constructs a Tool from the given name, description, and handler function. It validates the handler function's signature and generates a parameter schema based on the second argument of the handler function.
func BuildTool(name, description string, fn any) (*Tool, error) {
	if name == "" {
		return nil, fmt.Errorf("tool name cannot be empty")
	}

	if description == "" {
		return nil, fmt.Errorf("tool description cannot be empty")
	}

	fv := reflect.ValueOf(fn)
	ft := fv.Type()
	if ft.Kind() != reflect.Func {
		return nil, fmt.Errorf("tool handler must be a function")
	}

	if ft.NumIn() != 2 || !ft.In(0).Implements(ctxType) {
		return nil, fmt.Errorf("tool handler must have the signature func(context.Context, ArgsStruct) (Result, error)")
	}

	if ft.NumOut() != 2 || !ft.Out(1).Implements(errType) {
		return nil, fmt.Errorf("tool handler must have the signature func(context.Context, ArgsStruct) (Result, error)")
	}

	argsType := ft.In(1)
	if argsType.Kind() != reflect.Struct {
		return nil, fmt.Errorf("tool handler's second argument must be a struct, got %s", argsType.Kind())
	}

	return &Tool{
		Name:        name,
		Description: description,
		Schema:      SchemaFor(argsType),
		Handler:     fv,
		ArgsType:    argsType,
	}, nil
}
