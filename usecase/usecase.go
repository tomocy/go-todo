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

func (u *createUser) createUser(name, email, password string) error {
	ctx := context.Background()

	id, err := u.repo.NextID(ctx)
	if err != nil {
		return fmt.Errorf("failed to generate user id: %w", err)
	}
	user, err := todo.NewUser(id, name, email, password)
	if err != nil {
		return fmt.Errorf("failed to generate user: %w", err)
	}

	if err := u.repo.Save(ctx, user); err != nil {
		return fmt.Errorf("failed to save user: %w", err)
	}

	return nil
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
