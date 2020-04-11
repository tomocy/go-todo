package usecase

import (
	"context"
	"fmt"

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