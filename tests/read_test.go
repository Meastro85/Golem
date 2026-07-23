package tests

import (
	"Golem/tool"
	"testing"
)

func TestList(t *testing.T) {
	r := tool.NewRegistry()
	if err := r.Register("add", "adds two numbers", add); err != nil {
		t.Fatal(err)
	}
	fns := r.List()
	if len(fns) != 1 || fns[0].Name != "add" {
		t.Errorf("List() = %+v, want a single \"add\" function", fns)
	}
}

func TestGet(t *testing.T) {
	r := tool.NewRegistry()
	if err := r.Register("add", "adds two numbers", add); err != nil {
		t.Fatal(err)
	}

	fn, ok := r.Get("add")
	if !ok {
		t.Fatal("expected function to be registered")
	}
	if fn.Name != "add" || fn.Description != "adds two numbers" {
		t.Errorf("Get() = %+v, want name=add, description=adds two numbers", fn)
	}

	if _, ok := r.Get("missing"); ok {
		t.Error("expected missing function to not be found")
	}
}
