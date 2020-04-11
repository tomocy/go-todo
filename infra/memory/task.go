package memory

import (
	"context"

	"github.com/tomocy/go-todo"
)

type TaskRepo struct {
	tasks map[todo.TaskID]*todo.Task
}

func (r *TaskRepo) NextID(context.Context) (todo.TaskID, error) {
	return todo.TaskID(generateRandomString(50)), nil
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
