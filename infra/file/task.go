package file

import (
	"time"

	"github.com/tomocy/go-todo"
)

type taskRepo struct {
	fname string
	tasks map[todo.TaskID]*todo.Task
}

type task struct {
	ID             todo.TaskID     `json:"id"`
	UserID         todo.UserID     `json:"user_id"`
	Name           string          `json:"name"`
	Status         todo.TaskStatus `json:"status"`
	DueDate        time.Time       `json:"due_date"`
	PostponedTimes int             `json:"postponed_times"`
}
