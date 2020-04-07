package todo

type user struct {
	id      userID
	name    string
	status  userStatus
	profile profile
}

type userID string

type userStatus int

const (
	userActive userStatus = iota
	userInactive
)

type profile struct {
	email string
}

type task struct {
	id taskID
}

type taskID string
