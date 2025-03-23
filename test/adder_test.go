package test

import (
	"TaskList_for_go/tasklist"
	"testing"
)

func TestAdder_Sum(t *testing.T) {
	a := tasklist.NewAdder(1, 2)
	if a.Sum() != 3 {
		t.Error("Expected 3, got ", a.Sum())
	}
}
