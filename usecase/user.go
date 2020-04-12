package usecase

import (
	"context"
	"fmt"

	"github.com/tomocy/go-todo"
)

func NewCreateUser(repo todo.UserRepo) *createUser {
	return &createUser{
		repo: repo,
	}
}

type createUser struct {
	repo todo.UserRepo
}

func (u *createUser) Do(name, email, password string) (*todo.User, error) {
	ctx := context.TODO()

	id, err := u.repo.NextID(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to generate user id: %w", err)
	}
	hashed, err := todo.HashPassword(password)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}
	user, err := todo.NewUser(id, name, email, hashed)
	if err != nil {
		return nil, fmt.Errorf("failed to generate user: %w", err)
	}

	if err := u.repo.Save(ctx, user); err != nil {
		return nil, fmt.Errorf("failed to save user: %w", err)
	}

	return user, nil
}

type deleteUser struct {
	userRepo todo.UserRepo
}

func NewAuthenticateUser(userRepo todo.UserRepo, sessRepo todo.SessionRepo) *authenticateUser {
	return &authenticateUser{
		userRepo: userRepo,
		sessRepo: sessRepo,
	}
}

type authenticateUser struct {
	userRepo todo.UserRepo
	sessRepo todo.SessionRepo
}

func (u *authenticateUser) Do(email, password string) (*todo.User, *todo.Session, error) {
	ctx := context.TODO()

	user, err := u.userRepo.FindByEmail(ctx, email)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to find user by email: %w", err)
	}

	if !user.Password().IsSame(password) {
		return nil, nil, fmt.Errorf("incorrect password")
	}

	sessID, err := u.sessRepo.NextID(ctx)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to generate session id: %w", err)
	}
	sess, err := todo.NewSession(sessID, user.ID())
	if err != nil {
		return nil, nil, fmt.Errorf("failed to generate session: %w", err)
	}

	if err := u.sessRepo.Push(ctx, sess); err != nil {
		return nil, nil, fmt.Errorf("failed to save session: %w", err)
	}

	return user, sess, nil
}

func NewDeauthenticateUser(repo todo.SessionRepo) *deauthenticateUser {
	return &deauthenticateUser{
		repo: repo,
	}
}

type deauthenticateUser struct {
	repo todo.SessionRepo
}

func (u *deauthenticateUser) Do() error {
	ctx := context.TODO()

	return u.repo.Delete(ctx)
}
