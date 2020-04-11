package file

import "github.com/tomocy/go-todo"

type taskRepo struct {
	fname string
	tasks map[todo.TaskID]*todo.Task
}
