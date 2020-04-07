package usecase

import (
	"time"

	"github.com/tomocy/go-todo"
)

type createTask struct {
	repo todo.TaskRepo
}

func (u *createTask) createTask(name string, dueDate time.Time) error {

}
