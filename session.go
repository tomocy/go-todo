package todo

type Session struct {
	userID UserID
}

func (s *Session) UserID() UserID {
	return s.userID
}
