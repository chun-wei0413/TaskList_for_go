package entity

// Entity
type Project struct {
	name  ProjectName
	tasks []Task
}

func NewEmptyProject(name ProjectName) *Project {
	return &Project{
		name:  name,
		tasks: []Task{},
	}
}

func NewProject(name ProjectName, tasks []Task) *Project {
	return &Project{
		name:  name,
		tasks: tasks,
	}
}

func (p *Project) GetName() *ProjectName {
	return &p.name
}

func (p *Project) GetTasks() []Task {
	return p.tasks
}
