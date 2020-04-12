package file

import (
	"context"

	"github.com/tomocy/go-todo"
	"github.com/tomocy/go-todo/infra/rand"
)

type sessionRepo struct {
	fname string
}

func (r *sessionRepo) NextID(context.Context) (todo.SessionID, error) {
	return todo.SessionID(rand.GenerateString(30)), nil
}
