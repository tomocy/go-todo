package todo

import (
	"context"
	"fmt"
)

type SessionRepo interface {
	NextID(context.Context) (SessionID, error)
	Pull(context.Context) (*Session, error)
	Push(context.Context, *Session) error
}

func NewSession(id SessionID, userID UserID) *Session {
	return &Session{
		userID: userID,
	}
}

type Session struct {
	id     SessionID
	userID UserID
}

func (s *Session) UserID() UserID {
	return s.userID
}

func (s *Session) setUserID(id UserID) error {
	if id == "" {
		return fmt.Errorf("empty user id")
	}

	s.userID = id

	return nil
}

type SessionID string
