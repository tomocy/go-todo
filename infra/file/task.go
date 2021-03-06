package file

import (
	"context"
	"fmt"
	"time"

	"github.com/tomocy/go-todo"
	"github.com/tomocy/go-todo/infra/rand"
)

func NewTaskRepo(fname string) *taskRepo {
	return &taskRepo{
		fname: fname,
	}
}

type taskRepo struct {
	fname string
}

func (r *taskRepo) NextID(context.Context) (todo.TaskID, error) {
	return todo.TaskID(rand.GenerateString(50)), nil
}

func (r *taskRepo) Get(_ context.Context, userID todo.UserID) ([]*todo.Task, error) {
	s, err := load(r.fname)
	if err != nil {
		return nil, fmt.Errorf("failed to load tasks: %w", err)
	}

	ts := make([]*todo.Task, 0, len(s.Tasks))
	for _, t := range s.Tasks {
		if t.UserID != userID {
			continue
		}

		ts = append(ts, t.adapt())
	}

	return ts, nil
}

func (r *taskRepo) Find(_ context.Context, id todo.TaskID) (*todo.Task, error) {
	s, err := load(r.fname)
	if err != nil {
		return nil, fmt.Errorf("failed to load tasks: %w", err)
	}

	for _, t := range s.Tasks {
		if t.ID == id {
			return t.adapt(), nil
		}
	}

	return nil, fmt.Errorf("no such task")
}

func (r *taskRepo) Save(_ context.Context, t *todo.Task) error {
	s, err := load(r.fname)
	if err != nil {
		return fmt.Errorf("failed to load tasks: %w", err)
	}

	converted := convertTask(t)
	s.addTask(converted)

	return save(r.fname, s)
}

func (r *taskRepo) Delete(_ context.Context, id todo.TaskID) error {
	s, err := load(r.fname)
	if err != nil {
		return fmt.Errorf("failed to load tasks: %w", err)
	}

	for i, t := range s.Tasks {
		if t.ID != id {
			continue
		}

		s.Tasks = append(s.Tasks[:i], s.Tasks[i+1:]...)
	}

	if err := save(r.fname, s); err != nil {
		return fmt.Errorf("failed to save tasks: %w", err)
	}

	return nil
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
