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

func (p *Project) ContainTask(taskId TaskId) bool {
	for _, task := range p.tasks {
		id := task.GetId()
		if id.GetId() == taskId.GetId() {
			return true
		}
	}
	return false
}

func (p *Project) SetTaskDone(taskId TaskId, done bool) {
	for i, task := range p.tasks {
		id := task.GetId()
		if id.GetId() == taskId.GetId() {
			p.tasks[i].SetDone(done)
			return
		}
	}
}

func (p *Project) AddTask(task Task) {
	p.tasks = append(p.tasks, task)
}
