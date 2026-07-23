package tests

import (
	"Golem/kernel"
	"context"
	"encoding/json"
	"errors"
	"testing"
)

func TestExecute(t *testing.T) {
	k := kernel.NewKernel()
	if err := k.RegisterTool("add", "adds two numbers", add); err != nil {
		t.Fatal(err)
	}

	result, err := k.Execute(context.Background(), "add", json.RawMessage(`{"a": 2, "b": 3}`))
	if err != nil {
		t.Fatalf("Execute failed: %v", err)
	}
	if got, ok := result.(int); !ok || got != 5 {
		t.Errorf("Execute result = %v, want 5", result)
	}
}

func TestExecute_UnknownFunction(t *testing.T) {
	k := kernel.NewKernel()
	if _, err := k.Execute(context.Background(), "missing", nil); err == nil {
		t.Error("expected error for unknown function name")
	}
}

func TestExecute_MalformedArgs(t *testing.T) {
	k := kernel.NewKernel()
	if err := k.RegisterTool("add", "adds two numbers", add); err != nil {
		t.Fatal(err)
	}
	if _, err := k.Execute(context.Background(), "add", json.RawMessage(`not json`)); err == nil {
		t.Error("expected error for malformed JSON args")
	}
}

func TestExecute_HandlerErrorPropagates(t *testing.T) {
	k := kernel.NewKernel()
	boom := func(_ context.Context, _ addArgs) (float64, error) {
		return 0, errors.New("boom")
	}
	if err := k.RegisterTool("boom", "always fails", boom); err != nil {
		t.Fatal(err)
	}
	if _, err := k.Execute(context.Background(), "boom", json.RawMessage(`{}`)); err == nil {
		t.Error("expected handler error to propagate")
	}
}

func TestExecute_MissingArgsDefaultsZeroValue(t *testing.T) {
	k := kernel.NewKernel()
	if err := k.RegisterTool("add", "adds two numbers", add); err != nil {
		t.Fatal(err)
	}
	// No args at all -> zero-value struct -> 0 + 0.
	result, err := k.Execute(context.Background(), "add", nil)
	if err != nil {
		t.Fatalf("Execute failed: %v", err)
	}
	if got := result.(int); got != 0 {
		t.Errorf("result = %v, want 0", got)
	}
}
