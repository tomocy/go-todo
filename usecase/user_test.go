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

	user, err := u.Do(name, email, pass)
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
	userRepo := new(memory.UserRepo)
	createUsecase := createUser{
		repo: userRepo,
	}

	name, email, pass := "name", "email", "pass"
	createUsecase.Do(name, email, pass)

	sessRepo := memory.NewSessionRepo()
	u := authenticateUser{
		userRepo: userRepo,
		sessRepo: sessRepo,
	}

	user, sess, err := u.Do(email, pass)
	if err != nil {
		t.Errorf("should have authenticated user: %s", err)
		return
	}

	if err := assertUser(user, name, email, pass); err != nil {
		t.Errorf("should have returned the create user: %s", err)
		return
	}

	if err := assertSession(sess, user.ID()); err != nil {
		t.Errorf("should have returned the created session: %s", err)
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

func assertSession(s *todo.Session, userID todo.UserID) error {
	if s.UserID() != userID {
		return reportUnexpected("user id", s.UserID(), userID)
	}

	return nil
}
