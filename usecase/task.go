package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/tomocy/go-todo"
)

type createTask struct {
	repo todo.TaskRepo
}

func (u *createTask) do(userID todo.UserID, name string, dueDate time.Time) (*todo.Task, error) {
	ctx := context.Background()

	id, err := u.repo.NextID(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to generate task id: %w", err)
	}
	task, err := todo.NewTask(id, userID, name, dueDate)
	if err != nil {
		return nil, fmt.Errorf("failed to generate task: %w", err)
	}

	if err := u.repo.Save(ctx, task); err != nil {
		return nil, fmt.Errorf("failed to save task: %w", err)
	}

	return task, nil
}
