package file

import "github.com/tomocy/go-todo"

type userRepo struct {
	fname string
	users map[todo.UserID]*todo.User
}
