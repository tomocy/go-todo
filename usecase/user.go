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
	hashed, err := todo.HashPassword(password)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}
	user, err := todo.NewUser(id, name, email, hashed)
	if err != nil {
		return fmt.Errorf("failed to generate user: %w", err)
	}

	if err := u.repo.Save(ctx, user); err != nil {
		return fmt.Errorf("failed to save user: %w", err)
	}

	return nil
}

type authenticateUser struct {
	repo todo.UserRepo
}

func (u *authenticateUser) do(email, password string) error {
	ctx := context.TODO()

	user, err := u.repo.FindByEmail(ctx, email)
	if err != nil {
		return fmt.Errorf("failed to find user by email: %w", err)
	}

	if !user.Password().IsSame(password) {
		return fmt.Errorf("incorrect password")
	}

	return nil
}
