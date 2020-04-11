package todo

import (
	"context"
	"fmt"
	"time"
)

type TaskRepo interface {
	NextID(context.Context) (TaskID, error)
	Save(context.Context, *Task) error
}

func NewTask(id TaskID, userID UserID, name string, dueDate time.Time) (*Task, error) {
	t := new(Task)

	if err := t.setID(id); err != nil {
		return nil, err
	}
	if err := t.setUserID(userID); err != nil {
		return nil, err
	}
	if err := t.setName(name); err != nil {
		return nil, err
	}
	if err := t.setDueDate(dueDate); err != nil {
		return nil, err
	}

	return t, nil
}

type Task struct {
	id             TaskID
	userID         UserID
	name           string
	status         taskStatus
	dueDate        time.Time
	postponedTimes int
}

func (t *Task) ID() TaskID {
	return t.id
}

func (t *Task) setID(id TaskID) error {
	if id == "" {
		return fmt.Errorf("empty id")
	}

	t.id = id

	return nil
}

func (t *Task) UserID() UserID {
	return t.userID
}

func (t *Task) setUserID(id UserID) error {
	if id == "" {
		return fmt.Errorf("empty user id")
	}

	t.userID = id

	return nil
}

func (t *Task) setName(n string) error {
	if n == "" {
		return fmt.Errorf("empty name")
	}

	t.name = n

	return nil
}

func (t *Task) setDueDate(d time.Time) error {
	t.dueDate = d

	return nil
}

const postponedMaxTimes = 3

func (t *Task) postpone() error {
	if t.postponedTimes > postponedMaxTimes {
		return fmt.Errorf("postponed times exceeded: Task can be postponed up to %d", postponedMaxTimes)
	}

	t.dueDate.Add(24 * time.Hour)
	t.postponedTimes++

	return nil
}

type TaskID string

type taskStatus int

const (
	taskUndone taskStatus = iota
	taskDone
)
