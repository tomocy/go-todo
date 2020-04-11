package todo

import (
	"context"
	"fmt"
	"time"
)

type TaskRepo interface {
	NextID(context.Context) (taskID, error)
	Save(context.Context, *task) error
}

func NewTask(id taskID, name string, dueDate time.Time) (*task, error) {
	if id == "" {
		return nil, fmt.Errorf("empty id")
	}
	if name == "" {
		return nil, fmt.Errorf("empty name")
	}

	return &task{
		id:      id,
		name:    name,
		dueDate: dueDate,
	}, nil
}

type task struct {
	id             taskID
	userID         userID
	name           string
	status         taskStatus
	dueDate        time.Time
	postponedTimes int
}

func (t *task) setID(id taskID) error {
	if id == "" {
		return fmt.Errorf("empty id")
	}

	t.id = id

	return nil
}

func (t *task) setName(n string) error {
	if n == "" {
		return fmt.Errorf("empty name")
	}

	t.name = n

	return nil
}

func (t *task) setDueDate(d time.Time) error {
	t.dueDate = d

	return nil
}

const postponedMaxTimes = 3

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
