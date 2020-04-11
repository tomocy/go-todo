package file

import (
	"time"

	"github.com/tomocy/go-todo"
)

func NewTaskRepo(fname string) *taskRepo {
	return &taskRepo{
		fname: fname,
		tasks: make(map[todo.TaskID]*todo.Task),
	}
}

type taskRepo struct {
	fname string
	tasks map[todo.TaskID]*todo.Task
}

func (r *taskRepo) load() error {
	var ts []*task
	if err := load(r.fname, ts); err != nil {
		return err
	}

	for _, t := range ts {
		adapted := t.adapt()
		r.tasks[adapted.ID()] = adapted
	}

	return nil
}

type task struct {
	ID             todo.TaskID     `json:"id"`
	UserID         todo.UserID     `json:"user_id"`
	Name           string          `json:"name"`
	Status         todo.TaskStatus `json:"status"`
	DueDate        time.Time       `json:"due_date"`
	PostponedTimes int             `json:"postponed_times"`
}

func (t *task) adapt() *todo.Task {
	return todo.RecoverTask(
		t.ID, t.UserID, t.Name,
		t.Status, t.DueDate, t.PostponedTimes,
	)
}
