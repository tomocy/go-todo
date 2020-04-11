package file

import (
	"context"
	"fmt"

	"github.com/tomocy/go-todo"
)

type userRepo struct {
	fname string
	users map[todo.UserID]*todo.User
}

func (r *userRepo) NextID(context.Context) (todo.UserID, error) {
	return todo.UserID(generateRandomString(30)), nil
}

func (r *userRepo) FindByEmail(_ context.Context, email string) (*todo.User, error) {
	if err := r.load(); err != nil {
		return nil, fmt.Errorf("failed to load users: %w", err)
	}

	for _, u := range r.users {
		if u.Email() == email {
			return u, nil
		}
	}

	return nil, fmt.Errorf("no such user")
}

func (r *userRepo) load() error {
	s, err := load(r.fname)
	if err != nil {
		return err
	}

	for _, u := range s.Users {
		adapted := u.adapt()
		r.users[adapted.ID()] = adapted
	}

	return nil
}

type user struct {
	ID       todo.UserID     `json:"id"`
	Name     string          `json:"name"`
	Email    string          `json:"email"`
	Password todo.Password   `json:"password"`
	Status   todo.UserStatus `json:"status"`
}

func (u *user) adapt() *todo.User {
	return todo.RecoverUser(
		u.ID, u.Name, u.Email,
		u.Password, u.Status,
	)
}
