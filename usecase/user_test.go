package usecase

import (
	"fmt"
	"testing"

	"github.com/tomocy/go-todo"
	"github.com/tomocy/go-todo/infra/memory"
)

func TestCreateUser(t *testing.T) {
	repo := new(memory.UserRepo)
	u := createUser{
		repo: repo,
	}

	name, email, pass := "name", "email", "pass"

	user, err := u.do(name, email, pass)
	if err != nil {
		t.Errorf("should have created user: %s", err)
		return
	}

	if err := assertUser(user, name, email, pass); err != nil {
		t.Errorf("should have returned the create user: %s", err)
		return
	}
}

func TestAuthenticateUser(t *testing.T) {
	repo := new(memory.UserRepo)
	createUsecase := createUser{
		repo: repo,
	}

	name, email, pass := "name", "email", "pass"
	createUsecase.do(name, email, pass)

	u := authenticateUser{
		repo: repo,
	}

	user, err := u.do(email, pass)
	if err != nil {
		t.Errorf("should have authenticated user: %s", err)
		return
	}

	if err := assertUser(user, name, email, pass); err != nil {
		t.Errorf("should have returned the create user: %s", err)
		return
	}
}

func assertUser(u *todo.User, name, email, pass string) error {
	if u.Name() != name {
		return reportUnexpected("name", u.Name(), name)
	}
	if u.Email() != email {
		return reportUnexpected("email", u.Email(), email)
	}
	if !u.Password().IsSame(pass) {
		return fmt.Errorf("unexpected password: password is not correct")
	}

	return nil
}
