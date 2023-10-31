package models

type Task struct {
	ID        int
	Name      string
	Completed bool
}

func CreateTask(id int, name string) *Task {
	createdTask := Task{
		ID:        id,
		Name:      name,
		Completed: false,
	}
	return &createdTask
}
