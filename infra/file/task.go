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
		tasks: make(map[todo.TaskID]*todo.Task),
	}
}

type taskRepo struct {
	fname string
	tasks map[todo.TaskID]*todo.Task
}

func (r *taskRepo) NextID(context.Context) (todo.TaskID, error) {
	return todo.TaskID(rand.GenerateString(50)), nil
}

func (r *taskRepo) Get(_ context.Context, uid todo.UserID) ([]*todo.Task, error) {
	if err := r.load(); err != nil {
		return nil, fmt.Errorf("failed to load tasks: %w", err)
	}

	ts := make([]*todo.Task, 0, len(r.tasks))
	for _, t := range r.tasks {
		if t.UserID() != uid {
			continue
		}

		ts = append(ts, t)
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
	s, err := load(r.fname)
	if err != nil {
		return err
	}

	for _, t := range s.Tasks {
		adapted := t.adapt()
		r.tasks[adapted.ID()] = adapted
	}

	return nil
}

func (r *taskRepo) Save(_ context.Context, t *todo.Task) error {
	return r.save(t)
}

func (r *taskRepo) save(t *todo.Task) error {
	s, err := load(r.fname)
	if err != nil {
		return fmt.Errorf("failed to load status: %w", err)
	}

	converted := convertTask(t)
	s.addTask(converted)

	return save(r.fname, s)
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
