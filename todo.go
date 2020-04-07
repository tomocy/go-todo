package todo

import "time"

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

const postponedMaxTimes = 3

type task struct {
	id             taskID
	userID         userID
	name           string
	status         taskStatus
	dueDate        time.Time
	postponedTimes int
}

type taskID string

type taskStatus int

const (
	taskUndone taskStatus = iota
	taskDone
)
