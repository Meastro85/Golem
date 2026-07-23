package tests

import (
	"context"
)

type addArgs struct {
	A int `json:"a" description:"first number"`
	B int `json:"b" description:"second number"`
}

func add(_ context.Context, args addArgs) (int, error) {
	return args.A + args.B, nil
}
