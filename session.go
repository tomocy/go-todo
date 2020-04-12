package todo

func NewSession(userID UserID) *Session {
	return &Session{
		userID: userID,
	}
}

type Session struct {
	userID UserID
}

func (s *Session) UserID() UserID {
	return s.userID
}

type SessionID string
