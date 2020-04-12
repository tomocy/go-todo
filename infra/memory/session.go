package memory

import (
	"context"

	"github.com/tomocy/go-todo"
	"github.com/tomocy/go-todo/infra/rand"
)

func NewSessionRepo() *sessionRepo {
	return new(sessionRepo)
}

type sessionRepo struct {
	sess *todo.Session
}

func (r *sessionRepo) NextID(context.Context) (todo.SessionID, error) {
	return todo.SessionID(rand.GenerateString(30)), nil
}

func (r *sessionRepo) Pull(context.Context) (*todo.Session, error) {
	return r.sess, nil
}

func (r *sessionRepo) Push(_ context.Context, s *todo.Session) error {
	r.sess = s

	return nil
}

func (r *sessionRepo) Delete(context.Context) error {
	r.sess = nil

	return nil
}
