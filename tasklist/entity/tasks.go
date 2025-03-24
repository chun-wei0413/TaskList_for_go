package entity

// Entry 結構體用於儲存鍵值對
type Entry struct {
	Key   ProjectName
	Value []Task
}

// Tasks 結構體
type Tasks struct {
	tasks map[ProjectName][]Task
}

// NewTasks 構造函數，初始化 Tasks
func NewTasks() *Tasks {
	return &Tasks{
		tasks: make(map[ProjectName][]Task),
	}
}

// EntrySet 返回所有鍵值對的列表
func (t *Tasks) EntrySet() []Entry {
	entries := make([]Entry, 0, len(t.tasks))
	for k, v := range t.tasks {
		entries = append(entries, Entry{Key: k, Value: v})
	}
	return entries
}

// Put 方法，將任務列表與專案名稱關聯
func (t *Tasks) Put(name ProjectName, tasks []Task) {
	t.tasks[name] = tasks
}

// Get 方法，根據專案名稱返回任務列表
func (t *Tasks) Get(project ProjectName) []Task {
	return t.tasks[project]
}
