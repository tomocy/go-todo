package todo

import "context"

type SessionRepo interface {
	NextID(context.Context) (SessionID, error)
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

type SessionID string
