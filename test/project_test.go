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
		*entity.NewTask(1, "Test Task1", false),
		*entity.NewTask(2, "Test Task2", false),
	}
	project := entity.NewProject(*projectName, taskList)

	if len(project.GetTasks()) != 2 {
		t.Errorf("Expected 2 task, got %d", len(project.GetTasks()))
	}

	if project.GetTasks()[0].GetDescription() != "Test Task1" {
		t.Errorf("Expected 'Test Task1', got '%s'", project.GetTasks()[0].GetDescription())
	}

	if project.GetTasks()[1].GetDescription() != "Test Task2" {
		t.Errorf("Expected 'Test Task2', got '%s'", project.GetTasks()[1].GetDescription())
	}
}
