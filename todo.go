package todo

import (
	"fmt"
	"time"
)

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

func (t *task) postpone() error {
	if t.postponedTimes > postponedMaxTimes {
		return fmt.Errorf("postponed times exceeded: task can be postponed up to %d", postponedMaxTimes)
	}

	t.dueDate.Add(24 * time.Hour)
	t.postponedTimes++

	return nil
}

type taskID string

type taskStatus int

const (
	taskUndone taskStatus = iota
	taskDone
)
