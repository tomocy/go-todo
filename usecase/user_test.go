package usecase

import (
	"context"
	"fmt"
	"testing"

	"github.com/tomocy/go-todo"
	"github.com/tomocy/go-todo/infra/memory"
)

func TestCreateUser(t *testing.T) {
	repo := memory.NewUserRepo()
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

func TestDeleteUser(t *testing.T) {
	userRepo := memory.NewUserRepo()
	sessRepo := memory.NewSessionRepo()

	name, email, pass := "name", "email", "pass"

	createUsecase := createUser{
		repo: userRepo,
	}
	user, _ := createUsecase.Do(name, email, pass)

	authenticateUsecase := authenticateUser{
		userRepo: userRepo,
		sessRepo: sessRepo,
	}
	authenticateUsecase.Do(email, pass)

	u := deleteUser{
		userRepo: userRepo,
		sessRepo: sessRepo,
	}
	if err := u.Do(user.ID()); err != nil {
		t.Errorf("should have deleted user: %s", err)
		return
	}

	if _, err := sessRepo.Pull(context.Background()); err == nil {
		t.Errorf("should have deleted session: %s", err)
		return
	}
}

func TestAuthenticateUser(t *testing.T) {
	userRepo := memory.NewUserRepo()
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

func TestDeauthenticateUser(t *testing.T) {
	sessRepo := memory.NewSessionRepo()
	sess, _ := todo.NewSession(todo.SessionID("session id"), todo.UserID("user id"))
	sessRepo.Push(context.Background(), sess)

	u := deauthenticateUser{
		repo: sessRepo,
	}

	if err := u.Do(); err != nil {
		t.Errorf("should have deauthenticate user: %s", err)
		return
	}

	if _, err := sessRepo.Pull(context.Background()); err == nil {
		t.Errorf("should have deleted session")
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
