package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/tomocy/go-todo"
)

type createUser struct {
	repo todo.UserRepo
}

type createTask struct {
	repo todo.TaskRepo
}

func (u *createTask) createTask(name string, dueDate time.Time) error {
	ctx := context.Background()

	id, err := u.repo.NextID(ctx)
	if err != nil {
		return fmt.Errorf("failed to generate task id: %w", err)
	}
	t, err := todo.NewTask(id, name, dueDate)
	if err != nil {
		return fmt.Errorf("failed to generate task: %w", err)
	}

	if err := u.repo.Save(ctx, t); err != nil {
		return fmt.Errorf("failed to save task: %w", err)
	}

	return nil
}
