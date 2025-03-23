package test

import (
	"TaskList_for_go/tasklist/entity"
	"testing"
)

func TestTask(t *testing.T) {
	task := entity.NewTask(1, "Test Description", false)

	if task.GetId() != 1 {
		t.Error("Expected 1, got ", task.GetId())
	}

	if task.GetDescription() != "Test Description" {
		t.Error("Expected Test Description, got ", task.GetDescription())
	}

	if task.GetDone() != false {
		t.Error("Expected false, got ", task.GetDone())
	}

	task.SetDone(true)

	if task.GetDone() != true {
		t.Error("Expected true, got ", task.GetDone())
	}
}
