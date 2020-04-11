package cli

import (
	"io"

	"github.com/tomocy/go-todo"
	"github.com/tomocy/go-todo/infra/memory"
	"github.com/urfave/cli"
)

type app struct {
	*cli.App
	w io.Writer
}

func (a *app) init() {
	a.App = cli.NewApp()
	a.Name = "todo"
	a.Commands = []cli.Command{}
}

func (a *app) taskRepo() todo.TaskRepo {
	return new(memory.TaskRepo)
}
