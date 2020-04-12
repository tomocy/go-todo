package file

import (
	"context"
	"fmt"

	"github.com/tomocy/go-todo"
	"github.com/tomocy/go-todo/infra/rand"
)

type sessionRepo struct {
	fname string
}

func (r *sessionRepo) NextID(context.Context) (todo.SessionID, error) {
	return todo.SessionID(rand.GenerateString(30)), nil
}

func (r *sessionRepo) Pull(context.Context) (*todo.Session, error) {
	s, err := load(r.fname)
	if err != nil {
		return nil, fmt.Errorf("failed to load session: %w", err)
	}

	if s.Session == nil {
		return nil, fmt.Errorf("no session")
	}

	return s.Session.adapt(), nil
}

func convertSession(src *todo.Session) *session {
	return &session{
		ID:     src.ID(),
		UserID: src.UserID(),
	}
}

type session struct {
	ID     todo.SessionID `json:"id"`
	UserID todo.UserID    `json:"user_id"`
}

func (s *session) adapt() *todo.Session {
	return todo.RecoverSession(s.ID, s.UserID)
}
