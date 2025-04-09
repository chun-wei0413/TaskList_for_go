package test

import (
	"TaskList_for_go/tasklist/entity"
	"testing"
)

// setUp 設置一個帶有初始數據的 ToDoList 用於測試
func setUp() *entity.ToDoList {
	taskList := []entity.Task{
		*entity.NewTask(*entity.NewTaskId(1), "Task 1", false),
		*entity.NewTask(*entity.NewTaskId(2), "Task 2", true),
	}

	projectName := entity.NewProjectName("Project 1")
	id := *entity.NewToDoListId(1)
	toDoList := entity.NewToDoList(id)
	toDoList.AddProject(*projectName)
	for _, task := range taskList {
		toDoList.AddTask(*projectName, task.GetDescription(), task.GetDone())
	}
	return toDoList
}

// TestNewToDoList 測試創建新的 ToDoList
func TestNewToDoList(t *testing.T) {
	id := *entity.NewToDoListId(1)
	toDoList := entity.NewToDoList(id)
	if toDoList == nil {
		t.Fatal("ToDoList should not return nil")
	}

	projects := toDoList.GetProjects()
	if len(projects) != 0 {
		t.Errorf("Expected empty projects list, got %d projects", len(projects))
	}

	getId := toDoList.GetId()
	if getId.GetId() != 1 {
		t.Errorf("Expected ToDoList ID 1, got %d", getId.GetId())
	}
}

// TestAddProject 測試添加項目
func TestAddProject(t *testing.T) {
	toDoList := entity.NewToDoList(*entity.NewToDoListId(1))
	projectName := entity.NewProjectName("Project 1")

	toDoList.AddProject(*projectName)
	projects := toDoList.GetProjects()
	if len(projects) != 1 {
		t.Fatalf("Expected 1 project, got %d", len(projects))
	}

	if projects[0].GetName().ToString() != "Project 1" {
		t.Errorf("Expected project name 'Project 1', got '%s'", projects[0].GetName().ToString())
	}

	// 測試添加重複項目
	toDoList.AddProject(*projectName)
	if len(toDoList.GetProjects()) != 1 {
		t.Errorf("Expected 1 project after adding duplicate, got %d", len(toDoList.GetProjects()))
	}
}

// TestGetProject 測試獲取特定項目
func TestGetProject(t *testing.T) {
	toDoList := setUp()
	projectName := entity.NewProjectName("Project 1")
	nonExistent := entity.NewProjectName("Project 2")

	// 測試存在的項目
	project := toDoList.GetProject(*projectName)
	if project == nil {
		t.Fatal("Expected project 'Project 1' to be found, got nil")
	}
	if project.GetName().ToString() != "Project 1" {
		t.Errorf("Expected project name 'Project 1', got '%s'", project.GetName().ToString())
	}

	// 測試不存在的項目
	project = toDoList.GetProject(*nonExistent)
	if project != nil {
		t.Errorf("Expected nil for non-existent project 'Project 2', got %v", project)
	}
}

// TestAddTask 測試添加任務
func TestAddTask(t *testing.T) {
	toDoList := entity.NewToDoList(*entity.NewToDoListId(1))
	projectName := entity.NewProjectName("Project 1")
	toDoList.AddProject(*projectName)

	// 添加任務
	toDoList.AddTask(*projectName, "New Task", false)
	tasks := toDoList.GetTasks(*projectName)
	if len(tasks) != 1 {
		t.Fatalf("Expected 1 task, got %d", len(tasks))
	}

	id := tasks[0].GetId()
	if id.GetId() != 1 {
		t.Errorf("Expected task ID 1, got %d", id.GetId())
	}
	if tasks[0].GetDescription() != "New Task" {
		t.Errorf("Expected task description 'New Task', got '%s'", tasks[0].GetDescription())
	}
	if tasks[0].GetDone() {
		t.Errorf("Expected task not done, got done")
	}

	// 添加第二個任務，檢查 ID 遞增
	toDoList.AddTask(*projectName, "Another Task", true)
	tasks = toDoList.GetTasks(*projectName)
	if len(tasks) != 2 {
		t.Fatalf("Expected 2 tasks, got %d", len(tasks))
	}
	taskId := tasks[1].GetId()
	if taskId.GetId() != 2 {
		t.Errorf("Expected task ID 2, got %d", taskId.GetId())
	}
}

// TestGetTasksByProjectName 測試根據項目名稱獲取任務
func TestGetTasksByProjectName(t *testing.T) {
	toDoList := setUp()
	targetProjectName := entity.NewProjectName("Project 1")
	nonExistent := entity.NewProjectName("Project 2")

	tasks := toDoList.GetTasks(*targetProjectName)
	if len(tasks) != 2 {
		t.Fatalf("Expected 2 tasks, got %d", len(tasks))
	}

	expectedTasks := []struct {
		id   int64
		desc string
		done bool
	}{
		{1, "Task 1", false},
		{2, "Task 2", true},
	}
	for i, task := range tasks {
		id := task.GetId()
		if id.GetId() != expectedTasks[i].id {
			t.Errorf("Expected task ID %d, got %d", expectedTasks[i].id, id.GetId())
		}
		if task.GetDescription() != expectedTasks[i].desc {
			t.Errorf("Expected description '%s', got '%s'", expectedTasks[i].desc, task.GetDescription())
		}
		if task.GetDone() != expectedTasks[i].done {
			t.Errorf("Expected done %v, got %v", expectedTasks[i].done, task.GetDone())
		}
	}

	// 測試不存在的項目
	tasks = toDoList.GetTasks(*nonExistent)
	if tasks != nil {
		t.Errorf("Expected nil tasks for non-existent project, got %v", tasks)
	}
}

// TestSetDone 測試設置任務完成狀態
func TestSetDone(t *testing.T) {
	toDoList := setUp()
	taskId1 := entity.NewTaskId(1)
	taskId2 := entity.NewTaskId(2)
	nonExistent := entity.NewTaskId(3)

	// 設置任務 1 為完成
	toDoList.SetDone(*taskId1, true)
	tasks := toDoList.GetTasks(*entity.NewProjectName("Project 1"))
	if !tasks[0].GetDone() {
		t.Errorf("Expected task 1 to be done, got not done")
	}

	// 設置任務 2 為未完成
	toDoList.SetDone(*taskId2, false)
	if tasks[1].GetDone() {
		t.Errorf("Expected task 2 to be not done, got done")
	}

	// 測試不存在的任務 ID，不應改變任何狀態
	toDoList.SetDone(*nonExistent, true)
	if tasks[0].GetDone() != true || tasks[1].GetDone() != false {
		t.Errorf("Setting non-existent task ID should not affect existing tasks")
	}
}

// TestProjectContainTask 測試 Project 的 ContainTask 方法
func TestProjectContainTask(t *testing.T) {
	toDoList := setUp()
	project := toDoList.GetProject(*entity.NewProjectName("Project 1"))
	if project == nil {
		t.Fatal("Project should not be nil")
	}

	taskId1 := entity.NewTaskId(1)
	taskId3 := entity.NewTaskId(3)

	if !project.ContainTask(*taskId1) {
		t.Errorf("Expected project to contain task ID 1")
	}
	if project.ContainTask(*taskId3) {
		t.Errorf("Expected project not to contain task ID 3")
	}
}

// TestProjectSetTaskDone 測試 Project 的 SetTaskDone 方法
func TestProjectSetTaskDone(t *testing.T) {
	toDoList := setUp()
	project := toDoList.GetProject(*entity.NewProjectName("Project 1"))
	if project == nil {
		t.Fatal("Project should not be nil")
	}

	taskId1 := entity.NewTaskId(1)
	taskId3 := entity.NewTaskId(3)

	// 設置任務 1 為完成
	project.SetTaskDone(*taskId1, true)
	tasks := project.GetTasks()
	if !tasks[0].GetDone() {
		t.Errorf("Expected task 1 to be done, got not done")
	}

	// 設置不存在的任務 ID，不應影響現有任務
	project.SetTaskDone(*taskId3, true)
	if tasks[0].GetDone() != true || tasks[1].GetDone() != true {
		t.Errorf("Setting non-existent task ID should not affect existing tasks")
	}
}
