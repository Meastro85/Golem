package tests

import (
	"Golem/tool"
	"context"
	"testing"
)

func TestRegister(t *testing.T) {
	r := tool.NewRegistry()
	if err := r.Register("add", "adds two numbers", add); err != nil {
		t.Fatalf("RegisterFunc failed: %v", err)
	}

	fn, ok := r.Get("add")
	if !ok {
		t.Fatal("expected function to be registered")
	}
	if fn.Description != "adds two numbers" {
		t.Errorf("description = %q", fn.Description)
	}
	if fn.Schema.Type != "object" {
		t.Errorf("schema type = %q, want object", fn.Schema.Type)
	}
	if len(fn.Schema.Properties) != 2 {
		t.Errorf("expected 2 properties, got %d", len(fn.Schema.Properties))
	}
	if got := fn.Schema.Properties["a"].Description; got != "first number" {
		t.Errorf("property a description = %q", got)
	}
	if got := fn.Schema.Properties["a"].Type; got != "number" {
		t.Errorf("property a type = %q, want number", got)
	}
}

func TestRegister_RejectsBadSignatures(t *testing.T) {
	cases := map[string]any{
		"no context arg":  func(a, b int) (int, error) { return a + b, nil },
		"no error return": func(_ context.Context, args addArgs) int { return args.A },
		"non-struct args": func(_ context.Context, n int) (int, error) { return n, nil },
		"not a function":  42,
	}

	for name, fn := range cases {
		t.Run(name, func(t *testing.T) {
			r := tool.NewRegistry()
			if err := r.Register("x", "desc", fn); err == nil {
				t.Errorf("expected error for %s, got nil", name)
			}
		})
	}
}

func TestRegisterFunc_RequiresNameAndDescription(t *testing.T) {
	r := tool.NewRegistry()
	if err := r.Register("", "desc", add); err == nil {
		t.Error("expected error for empty name")
	}
	if err := r.Register("add", "", add); err == nil {
		t.Error("expected error for empty description")
	}
}

// Might change this later to throw an error instead of overwriting, but for now this is the behavior.
func TestRegisterFunc_OverwritesOnDuplicateName(t *testing.T) {
	r := tool.NewRegistry()
	if err := r.Register("add", "first version", add); err != nil {
		t.Fatal(err)
	}
	replacement := func(_ context.Context, args addArgs) (int, error) {
		return args.A + args.B + 100, nil
	}
	if err := r.Register("add", "second version", replacement); err != nil {
		t.Fatal(err)
	}

	fn, _ := r.Get("add")
	if fn.Description != "second version" {
		t.Errorf("expected second registration to win, description = %q", fn.Description)
	}
	if len(r.List()) != 1 {
		t.Errorf("expected 1 function after re-registering the same name, got %d", len(r.List()))
	}
}
