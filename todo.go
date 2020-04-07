package todo

type user struct {
	id     userID
	name   string
	status userStatus
}

type userID string

type userStatus int

const (
	userActive userStatus = iota
	userInactive
)
