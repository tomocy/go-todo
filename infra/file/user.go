package file

import (
	"context"
	"fmt"

	"github.com/tomocy/go-todo"
	"github.com/tomocy/go-todo/infra/rand"
)

func NewUserRepo(fname string) *userRepo {
	return &userRepo{
		fname: fname,
	}
}

type userRepo struct {
	fname string
}

func (r *userRepo) NextID(context.Context) (todo.UserID, error) {
	return todo.UserID(rand.GenerateString(30)), nil
}

func (r *userRepo) FindByEmail(_ context.Context, email string) (*todo.User, error) {
	s, err := load(r.fname)
	if err != nil {
		return nil, fmt.Errorf("failed to load tasks: %w", err)
	}

	for _, u := range s.Users {
		if u.Email == email {
			return u.adapt(), nil
		}
	}

	return nil, fmt.Errorf("no such user")
}

func (r *userRepo) Save(_ context.Context, u *todo.User) error {
	s, err := load(r.fname)
	if err != nil {
		return fmt.Errorf("failed to load tasks: %w", err)
	}

	converted := convertUser(u)
	s.addUser(converted)

	return save(r.fname, s)
}

func (r *userRepo) Delete(_ context.Context, id todo.UserID) error {
	s, err := load(r.fname)
	if err != nil {
		return fmt.Errorf("failed to load users: %w", err)
	}

	for i, u := range s.Users {
		if u.ID != id {
			continue
		}

		s.Users = append(s.Users[i:], s.Users[i+1:]...)
	}

	if err := save(r.fname, s); err != nil {
		return fmt.Errorf("failed to save users: %w", err)
	}

	return nil
}

func convertUser(src *todo.User) *user {
	return &user{
		ID:       src.ID(),
		Name:     src.Name(),
		Email:    src.Email(),
		Password: src.Password(),
		Status:   src.Status(),
	}
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
