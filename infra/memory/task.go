package memory

import (
	"context"
	"fmt"

	"github.com/tomocy/go-todo"
	"github.com/tomocy/go-todo/infra/rand"
)

type TaskRepo struct {
	tasks map[todo.TaskID]*todo.Task
}

func (r *TaskRepo) NextID(context.Context) (todo.TaskID, error) {
	return todo.TaskID(rand.GenerateString(50)), nil
}

func (r *TaskRepo) Get(_ context.Context, uid todo.UserID) ([]*todo.Task, error) {
	r.initIfNecessary()

	ts := make([]*todo.Task, 0, len(r.tasks))
	for _, t := range r.tasks {
		if t.UserID() != uid {
			continue
		}

		ts = append(ts, t)
	}

	return ts, nil
}

func (r *TaskRepo) Find(_ context.Context, id todo.TaskID) (*todo.Task, error) {
	r.initIfNecessary()

	if t, ok := r.tasks[id]; ok {
		return t, nil
	}

	return nil, fmt.Errorf("no such task")
}

func (r *TaskRepo) Save(_ context.Context, t *todo.Task) error {
	r.initIfNecessary()

	r.tasks[t.ID()] = t

	return nil
}

func (r *TaskRepo) initIfNecessary() {
	if r.tasks == nil {
		r.tasks = make(map[todo.TaskID]*todo.Task)
	}
}
