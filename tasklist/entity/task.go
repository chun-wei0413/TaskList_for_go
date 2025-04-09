package entity

// Entity
type Task struct {
	id          TaskId
	description string
	done        bool
}

func NewTask(id TaskId, description string, done bool) *Task {
	return &Task{
		id:          id,
		description: description,
		done:        done,
	}
}

func (task *Task) GetId() TaskId {
	return task.id
}

func (task *Task) GetDescription() string {
	return task.description
}

func (task *Task) GetDone() bool {
	return task.done
}

func (task *Task) SetDone(done bool) {
	task.done = done
}
