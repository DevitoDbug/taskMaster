package models

type Task struct {
	Id        int
	Name      string
	Completed bool
}

func CreateTask(id int, name string) *Task {
	createdTask := Task{
		Id:        id,
		Name:      name,
		Completed: false,
	}
	return &createdTask
}
