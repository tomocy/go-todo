package cli

import (
	"github.com/tomocy/go-todo"
	"github.com/tomocy/go-todo/infra/memory"
	"github.com/urfave/cli"
)

type app struct {
	*cli.App
}

func (a *app) init() {
	a.App = cli.NewApp()
	a.Name = "todo"
	a.Commands = []cli.Command{}
}

func (a *app) taskRepo() todo.TaskRepo {
	return new(memory.TaskRepo)
}
