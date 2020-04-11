package memory

import (
	"context"
	"fmt"

	"github.com/tomocy/go-todo"
)

type UserRepo struct {
	users map[todo.UserID]*todo.User
}

func (r *UserRepo) NextID(context.Context) (todo.UserID, error) {
	return todo.UserID(generateRandomString(30)), nil
}

func (r *UserRepo) FindByEmail(_ context.Context, email string) (*todo.User, error) {
	r.initIfNecessary()

	for _, u := range r.users {
		if u.Email() == email {
			return u, nil
		}
	}

	return nil, fmt.Errorf("no such user")
}

func (u *UserRepo) initIfNecessary() {
	if u.users == nil {
		u.users = make(map[todo.UserID]*todo.User)
	}
}
