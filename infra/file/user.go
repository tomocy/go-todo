package file

import (
	"context"

	"github.com/tomocy/go-todo"
)

type userRepo struct {
	fname string
	users map[todo.UserID]*todo.User
}

func (r *userRepo) NextID(context.Context) (todo.UserID, error) {
	return todo.UserID(generateRandomString(30)), nil
}
