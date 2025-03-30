package test

import (
	"TaskList_for_go/tasklist/entity"
	"testing"
)

func setUp() (toDoList *entity.ToDoList) {
	taskList := []entity.Task{
		*entity.NewTask(1, "Task 1", false),
		*entity.NewTask(2, "Task 2", true),
	}

	projectName := entity.NewProjectName("Project 1")
	toDoList = entity.NewToDoList()
	toDoList.AddProjects(*projectName, taskList)

	return toDoList
}

func TestNewToDoList(t *testing.T) {
	projects := entity.NewToDoList()
	if projects == nil {
		t.Fatal("projects should not return nil")
	}

	entries := projects.GetProjects()
	if len(entries) != 0 {
		t.Errorf("Expected empty projects list, got %d projects", len(entries))
	}
}

func TestGetProject(t *testing.T) {
	toDoList := setUp()

	projects := toDoList.GetProjects()

	if len(projects) == 0 {
		t.Fatal("Projects list is empty")
	}

	if projects[0].GetName() == nil {
		t.Fatal("Project name should not be nil")
	}

	if projects[0].GetName().ToString() != "Project 1" {
		t.Errorf("Expected project name 'Project 1', got '%s'", projects[0].GetName().ToString())
	}
}

func TestAddProject(t *testing.T) {
	toDoList := setUp()

	projects := toDoList.GetProjects()
	if len(projects) != 1 {
		t.Fatalf("Expected 1 project, got %d", len(projects))
	}

	if (projects)[0].GetName() == nil {
		t.Fatal("Project name should not be nil")
	}

	if (projects)[0].GetName().ToString() != "Project 1" {
		t.Errorf("Expected project name 'Project 1', got '%s'", (projects)[0].GetName().ToString())
	}
}

func TestGetTasksByProjectName(t *testing.T) {
	taskList := []entity.Task{
		*entity.NewTask(1, "Task 1", false),
		*entity.NewTask(2, "Task 2", true),
	}
	targetProjectName := entity.NewProjectName("Project 1")
	toDoList := setUp()

	tasks := toDoList.GetTasks(*targetProjectName)
	if len(tasks) != len(taskList) {
		t.Fatalf("Expected %d tasks, got %d", len(taskList), len(tasks))
	}

	for i, task := range tasks {
		if task.GetId() != taskList[i].GetId() {
			t.Errorf("Expected task ID %d, got %d", taskList[i].GetId(), task.GetId())
		}
	}
}
