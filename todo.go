package todo

type user struct {
	id   userID
	name string
}

type userID string

type userStatus int

const (
	userActive userStatus = iota
	userInactive
)
