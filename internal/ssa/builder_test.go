package ssa

import (
	"fmt"
	"os"
	"testing"
)

func TestParseAndBuildSSA(t *testing.T) {
	tests := []string{
		"simpleIf",
		"ifElse",
		"nestedIf",
		"simpleLoop",
		"whileLoop",
		"nestedLoops",
		"loopWithBreakContinue",
		"switchExample",
		"multipleReturns",
		"withPanic",
		"factorial",
		"withGoto",
		"withDefer",
		"complexConditions",
	}

	bytes, err := os.ReadFile("../../homework1/examples/test_functions.go")
	if err != nil {
		t.Fatalf("test functions not found")
	}
	src := string(bytes)

	for _, tt := range tests {
		fn, err := NewBuilder().ParseAndBuildSSA(src, tt)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if fn == nil {
			t.Fatalf("function %s not found", tt)
		}

		fmt.Println(SprintSSA(fn))
	}
}
