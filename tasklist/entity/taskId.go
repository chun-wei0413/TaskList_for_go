package entity

// ValueObject
type TaskId struct {
	id int64
}

func NewTaskId(id int64) *TaskId {
	return &TaskId{
		id: id,
	}
}

func (t *TaskId) GetId() int64 {
	return t.id
}
