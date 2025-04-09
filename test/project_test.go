package test

import (
	"TaskList_for_go/tasklist/entity"
	"testing"
)

func TestEmptyProject(t *testing.T) {
	projectName := entity.NewProjectName("test")
	taskList := []entity.Task{}
	project := entity.NewProject(*projectName, taskList)

	if project.GetName().ToString() != "test" {
		t.Errorf("projectNameOf() = %v; want test", project.GetName().ToString())
	}

	if len(project.GetTasks()) != 0 {
		t.Errorf("Expected 0 task, got %d", len(project.GetTasks()))
	}
}

func TestGetName(t *testing.T) {
	projectName := entity.NewProjectName("test")
	taskList := []entity.Task{}
	project := entity.NewProject(*projectName, taskList)
	if project.GetName().ToString() != "test" {
		t.Errorf("projectNameOf() = %v; want test", project.GetName().ToString())
	}
}

func TestGetTasks(t *testing.T) {
	projectName := entity.NewProjectName("test")
	taskList := []entity.Task{
		*entity.NewTask(*entity.NewTaskId(1), "Test Task1", false),
		*entity.NewTask(*entity.NewTaskId(2), "Test Task2", false),
	}
	project := entity.NewProject(*projectName, taskList)

	if len(project.GetTasks()) != 2 {
		t.Errorf("Expected 2 task, got %d", len(project.GetTasks()))
	}

	if project.ContainTask(*entity.NewTaskId(1)) != true {
		t.Errorf("Expected task with ID 1 to be present")
	}

	if project.ContainTask(*entity.NewTaskId(2)) != true {
		t.Errorf("Expected task with ID 3 to not be present")
	}

	project.SetTaskDone(*entity.NewTaskId(1), true)

	if project.GetTasks()[0].GetDone() != true {
		t.Errorf("Expected task with ID 1 to be done")
	}

	if project.GetTasks()[0].GetDescription() != "Test Task1" {
		t.Errorf("Expected 'Test Task1', got '%s'", project.GetTasks()[0].GetDescription())
	}

	if project.GetTasks()[1].GetDescription() != "Test Task2" {
		t.Errorf("Expected 'Test Task2', got '%s'", project.GetTasks()[1].GetDescription())
	}

	project.AddTask(*entity.NewTask(*entity.NewTaskId(3), "Test Task2", false))

	if len(project.GetTasks()) != 3 {
		t.Errorf("Expected 3 task, got %d", len(project.GetTasks()))
	}

	if project.ContainTask(*entity.NewTaskId(3)) != true {
		t.Errorf("Expected task with ID 3 to be present")
	}
}
