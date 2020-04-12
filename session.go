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

func NewSession(id SessionID, userID UserID) (*Session, error) {
	s := new(Session)

	if err := s.setID(id); err != nil {
		return nil, err
	}
	if err := s.setUserID(userID); err != nil {
		return nil, err
	}

	return s, nil
}

type Session struct {
	id     SessionID
	userID UserID
}

func (s *Session) setID(id SessionID) error {
	if id == "" {
		return fmt.Errorf("empty id")
	}

	s.id = id

	return nil
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
