package memory

import (
	"context"
	"fmt"

	"github.com/tomocy/go-todo"
	"github.com/tomocy/go-todo/infra/rand"
)

func NewTaskRepo() *taskRepo {
	return &taskRepo{
		tasks: make(map[todo.TaskID]*todo.Task),
	}
}

type taskRepo struct {
	tasks map[todo.TaskID]*todo.Task
}

func (r *taskRepo) NextID(context.Context) (todo.TaskID, error) {
	return todo.TaskID(rand.GenerateString(50)), nil
}

func (r *taskRepo) Get(_ context.Context, uid todo.UserID) ([]*todo.Task, error) {
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
	if t, ok := r.tasks[id]; ok {
		return t, nil
	}

	return nil, fmt.Errorf("no such task")
}

func (r *taskRepo) Save(_ context.Context, t *todo.Task) error {
	r.tasks[t.ID()] = t

	return nil
}

func (r *taskRepo) Delete(_ context.Context, id todo.TaskID) error {
	delete(r.tasks, id)

	return nil
}
