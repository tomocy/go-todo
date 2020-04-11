package memory

import (
	"context"
	"fmt"

	"github.com/tomocy/go-todo"
)

type TaskRepo struct {
	tasks map[todo.TaskID]*todo.Task
}

func (r *TaskRepo) NextID(context.Context) (todo.TaskID, error) {
	return todo.TaskID(generateRandomString(50)), nil
}

func (r *TaskRepo) Get(context.Context) ([]*todo.Task, error) {
	r.initIfNecessary()

	ts := make([]*todo.Task, len(r.tasks))
	var i int
	for _, t := range r.tasks {
		ts[i] = t
		i++
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
