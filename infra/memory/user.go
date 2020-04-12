package memory

import (
	"context"
	"fmt"

	"github.com/tomocy/go-todo"
	"github.com/tomocy/go-todo/infra/rand"
)

func NewUserRepo() *userRepo {
	return &userRepo{
		users: make(map[todo.UserID]*todo.User),
	}
}

type userRepo struct {
	users map[todo.UserID]*todo.User
}

func (r *userRepo) NextID(context.Context) (todo.UserID, error) {
	return todo.UserID(rand.GenerateString(30)), nil
}

func (r *userRepo) FindByEmail(_ context.Context, email string) (*todo.User, error) {
	for _, u := range r.users {
		if u.Email() == email {
			return u, nil
		}
	}

	return nil, fmt.Errorf("no such user")
}

func (r *userRepo) Save(_ context.Context, u *todo.User) error {
	r.users[u.ID()] = u

	return nil
}

func (r *userRepo) Delete(_ context.Context, id todo.UserID) error {
	delete(r.users, id)

	return nil
}
