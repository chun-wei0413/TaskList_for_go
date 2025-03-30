package entity

type ToDoList struct {
	projects []Project
}

func NewToDoList() *ToDoList {
	return &ToDoList{
		projects: []Project{},
	}
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
