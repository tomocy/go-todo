package file

import (
	"context"
	"fmt"
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

func (r *taskRepo) Get(context.Context) ([]*todo.Task, error) {
	if err := r.load(); err != nil {
		return nil, fmt.Errorf("failed to load tasks: %w", err)
	}

	ts := make([]*todo.Task, len(r.tasks))
	var i int
	for _, t := range r.tasks {
		ts[i] = t
		i++
	}

	return ts, nil
}

func (r *taskRepo) Find(_ context.Context, id todo.TaskID) (*todo.Task, error) {
	if err := r.load(); err != nil {
		return nil, fmt.Errorf("failed to load tasks: %w", err)
	}

	if t, ok := r.tasks[id]; ok {
		return t, nil
	}

	return nil, fmt.Errorf("no such task")
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

func (r *taskRepo) save(t *todo.Task) error {
	converted := convertTask(t)

	return save(r.fname, converted)
}

func convertTask(src *todo.Task) *task {
	return &task{
		ID:             src.ID(),
		UserID:         src.UserID(),
		Name:           src.Name(),
		Status:         src.Status(),
		DueDate:        src.DueDate(),
		PostponedTimes: src.PostponedTimes(),
	}
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
