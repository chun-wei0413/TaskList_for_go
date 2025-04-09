package test

import (
	"TaskList_for_go/tasklist/entity"
	"testing"
)

func TestTask(t *testing.T) {
	task := entity.NewTask(*entity.NewTaskId(1), "Test Description", false)

	id := task.GetId()
	if id.GetId() != 1 {
		t.Error("Expected 1, got ", id.GetId())
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
