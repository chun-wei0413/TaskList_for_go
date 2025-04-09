package entity

// AggregateRoot
type ToDoList struct {
	id         ToDoListId
	projects   []Project
	lastTaskId int64
}

func NewToDoList(id ToDoListId) *ToDoList {
	return &ToDoList{
		id:       id,
		projects: []Project{},
	}
}

func (t *ToDoList) GetId() ToDoListId {
	return t.id
}

func (t *ToDoList) GetProjects() []Project {
	projectsCopy := make([]Project, len(t.projects))
	copy(projectsCopy, t.projects)
	return projectsCopy
}

func (t *ToDoList) AddProject(name ProjectName) {
	for _, p := range t.projects {
		if p.GetName().ToString() == name.ToString() {
			return // Project already exists
		}
	}
	t.projects = append(t.projects, *NewEmptyProject(name))
}

func (t *ToDoList) GetTasks(projectName ProjectName) []Task {
	for _, p := range t.projects {
		if p.name == projectName {
			return p.tasks
		}
	}
	return nil
}

// AddTask 向指定項目添加任務，使用下一個任務 ID
func (t *ToDoList) AddTask(name ProjectName, description string, done bool) {
	project := t.GetProject(name)
	if project != nil {
		taskID := t.nextTaskId()
		project.AddTask(*NewTask(*NewTaskId(taskID), description, done))
	}
}

// GetProject 返回指定名稱的項目，若不存在返回 nil
func (t *ToDoList) GetProject(projectName ProjectName) *Project {
	for i, p := range t.projects {
		if p.GetName().ToString() == projectName.ToString() {
			return &t.projects[i]
		}
	}
	return nil
}

// SetDone 設置指定任務的完成狀態
func (t *ToDoList) SetDone(taskId TaskId, done bool) {
	for _, p := range t.projects {
		if p.ContainTask(taskId) {
			p.SetTaskDone(taskId, done)
			return
		}
	}
}

// nextTaskId 生成下一個任務 ID
func (t *ToDoList) nextTaskId() int64 {
	t.lastTaskId++
	return t.lastTaskId
}
