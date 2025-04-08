package entity

// ValueObject
type ToDoListId struct {
	id int64
}

func NewToDoListId(id int64) *ToDoListId {
	return &ToDoListId{
		id: id,
	}
}

func (t *ToDoListId) GetId() int64 {
	return t.id
}
