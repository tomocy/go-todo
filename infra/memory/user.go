package memory

import (
	"context"

	"github.com/tomocy/go-todo"
)

type UserRepo struct {
	users []*todo.User
}

func (r *UserRepo) NextID(context.Context) (todo.UserID, error) {
	return todo.UserID(generateRandomString(30)), nil
}
