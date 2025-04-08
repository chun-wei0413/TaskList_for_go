package entity

// AggregateRoot
type ToDoList struct {
	id       ToDoListId
	projects []Project
}

func NewToDoList(id ToDoListId) *ToDoList {
	return &ToDoList{
		id:       id,
		projects: []Project{},
	}
}

func (t *ToDoList) GetToDoListId() ToDoListId {
	return t.id
}

func (t *ToDoList) GetProjects() []Project {
	return t.projects
}

func (t *ToDoList) AddProjects(name ProjectName, tasks []Task) {
	t.projects = append(t.projects, Project{name, tasks})
}

func (t *ToDoList) GetTasks(project ProjectName) []Task {
	for _, p := range t.projects {
		if p.name == project {
			return p.tasks
		}
	}
	return nil
}
