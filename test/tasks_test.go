package test

import (
	"TaskList_for_go/tasklist/entity"
	"testing"
)

import (
	"reflect"
)

func setUp() (taskList []entity.Task) {
	taskList = []entity.Task{
		*entity.NewTask(1, "Task 1", false),
		*entity.NewTask(2, "Task 2", true),
	}
	return
}

// TestNewTasks 測試 NewTasks 構造函數
func TestNewTasks(t *testing.T) {
	tasks := entity.NewTasks()
	if tasks == nil {
		t.Fatal("NewTasks should not return nil")
	}
	// 通過 EntrySet 間接檢查 map 是否初始化且為空
	entries := tasks.EntrySet()
	if len(entries) != 0 {
		t.Errorf("NewTasks should initialize an empty map, got %d entries", len(entries))
	}
}

// TestPut 測試 Put 方法
func TestTasksPut(t *testing.T) {
	tasks := entity.NewTasks()
	taskList := setUp()

	project1 := entity.NewProjectName("project1")
	// 測試正常情況
	tasks.Put(*project1, taskList)
	entries := tasks.EntrySet()
	if len(entries) != 1 {
		t.Errorf("Expected 1 entry in map, got %d", len(entries))
	}
	if got := tasks.Get(*project1); !reflect.DeepEqual(got, taskList) {
		t.Errorf("Expected tasks %v, got %v", taskList, got)
	}

	// 測試覆蓋現有鍵
	newTaskList := []entity.Task{*entity.NewTask(3, "Task 3", false)}
	tasks.Put(*project1, newTaskList)
	entries = tasks.EntrySet()
	if len(entries) != 1 {
		t.Errorf("Expected 1 entry in map after overwrite, got %d", len(entries))
	}
	if got := tasks.Get(*project1); !reflect.DeepEqual(got, newTaskList) {
		t.Errorf("Expected tasks %v after overwrite, got %v", newTaskList, got)
	}

	project2 := entity.NewProjectName("project2")
	// 測試空任務列表
	tasks.Put(*project2, []entity.Task{})
	entries = tasks.EntrySet()
	if len(entries) != 2 {
		t.Errorf("Expected 2 entries in map, got %d", len(entries))
	}
	if got := tasks.Get(*project2); len(got) != 0 {
		t.Errorf("Expected empty task list for project2, got %v", got)
	}
}

func TestTasksGet(t *testing.T) {
	tasks := entity.NewTasks()
	taskList := setUp()
	project1 := entity.NewProjectName("project1")
	tasks.Put(*project1, taskList)

	// 測試正常情況
	if got := tasks.Get(*project1); !reflect.DeepEqual(got, taskList) {
		t.Errorf("Expected tasks %v, got %v", taskList, got)
	}

	// 測試不存在的鍵
	project2 := entity.NewProjectName("project2")
	if got := tasks.Get(*project2); len(got) != 0 {
		t.Errorf("Expected empty task list for project2, got %v", got)
	}
}

func TestTasksEntrySet(t *testing.T) {
	tasks := entity.NewTasks()
	taskList := setUp()
	project1 := entity.NewProjectName("project1")
	tasks.Put(*project1, taskList)

	// 測試正常情況
	entries := tasks.EntrySet()
	if len(entries) != 1 {
		t.Errorf("Expected 1 entry in map, got %d", len(entries))
	}
	if entries[0].Key != *project1 {
		t.Errorf("Expected key %v, got %v", project1, entries[0].Key)
	}
	if !reflect.DeepEqual(entries[0].Value, taskList) {
		t.Errorf("Expected tasks %v, got %v", taskList, entries[0].Value)
	}

}
