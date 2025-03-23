package test

import (
	"TaskList_for_go/tasklist/entity"
	"testing"
)

import (
	"reflect"
)

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
func TestPut(t *testing.T) {
	tasks := entity.NewTasks()
	taskList := []entity.Task{
		*entity.NewTask(1, "Task 1", false),
		*entity.NewTask(2, "Task 2", true),
	}

	// 測試正常情況
	tasks.Put("project1", taskList)

	entries := tasks.EntrySet()

	if len(entries) != 1 {
		t.Errorf("Expected 1 entry in map, got %d", len(entries))
	}
	if got := tasks.Get("project1"); !reflect.DeepEqual(got, taskList) {
		t.Errorf("Expected tasks %v, got %v", taskList, got)
	}

	// 測試覆蓋現有鍵
	newTaskList := []entity.Task{*entity.NewTask(3, "Task 3", false)}
	tasks.Put("project1", newTaskList)
	entries = tasks.EntrySet()
	if len(entries) != 1 {
		t.Errorf("Expected 1 entry in map after overwrite, got %d", len(entries))
	}
	if got := tasks.Get("project1"); !reflect.DeepEqual(got, newTaskList) {
		t.Errorf("Expected tasks %v after overwrite, got %v", newTaskList, got)
	}

	// 測試空任務列表
	tasks.Put("project2", []entity.Task{})
	entries = tasks.EntrySet()
	if len(entries) != 2 {
		t.Errorf("Expected 2 entries in map, got %d", len(entries))
	}
	if got := tasks.Get("project2"); len(got) != 0 {
		t.Errorf("Expected empty task list for project2, got %v", got)
	}
}

// TestGet 測試 Get 方法
func TestGet(t *testing.T) {
	tasks := entity.NewTasks()
	taskList := []entity.Task{
		*entity.NewTask(1, "Task 1", false),
		*entity.NewTask(2, "Task 2", true),
	}
	tasks.Put("project1", taskList)

	// 測試正常情況
	got := tasks.Get("project1")
	if !reflect.DeepEqual(got, taskList) {
		t.Errorf("Expected tasks %v, got %v", taskList, got)
	}

	// 測試不存在的鍵
	got = tasks.Get("project2")
	if got != nil {
		t.Errorf("Expected nil for non-existent key, got %v", got)
	}
}

// TestEntrySet 測試 EntrySet 方法
func TestEntrySet(t *testing.T) {
	tasks := entity.NewTasks()
	taskList1 := []entity.Task{*entity.NewTask(1, "Task 1", false)}
	taskList2 := []entity.Task{*entity.NewTask(2, "Task 2", true)}
	tasks.Put("project1", taskList1)
	tasks.Put("project2", taskList2)

	// 測試正常情況
	entries := tasks.EntrySet()
	if len(entries) != 2 {
		t.Errorf("Expected 2 entries, got %d", len(entries))
	}

	// 檢查 entries 是否包含所有鍵值對
	expected := map[string][]entity.Task{
		"project1": taskList1,
		"project2": taskList2,
	}
	for _, entry := range entries {
		if expectedEntry, exists := expected[entry.Key]; !exists {
			t.Errorf("Unexpected key in EntrySet: %s", entry.Key)
		} else if !reflect.DeepEqual(entry.Value, expectedEntry) {
			t.Errorf("For key %s, expected tasks %v, got %v", entry.Key, expectedEntry, entry.Value)
		}
	}

	// 測試空 map
	emptyTasks := entity.NewTasks()
	entries = emptyTasks.EntrySet()
	if len(entries) != 0 {
		t.Errorf("Expected 0 entries for empty map, got %d", len(entries))
	}
}
