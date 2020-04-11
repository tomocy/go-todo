package memory

import (
	"github.com/tomocy/go-todo"
)

type TaskRepo struct {
	tasks map[todo.TaskID]*todo.Task
}

func (r *TaskRepo) initIfNecessary() {
	if r.tasks == nil {
		r.tasks = make(map[todo.TaskID]*todo.Task)
	}
}
