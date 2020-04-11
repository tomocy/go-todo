package cli

import (
	"fmt"
	"io"

	"github.com/tomocy/go-todo"
	"github.com/tomocy/go-todo/infra/memory"
	"github.com/urfave/cli"
)

func NewApp(w io.Writer) *app {
	a := &app{
		w: w,
	}
	a.init()

	return a
}

type app struct {
	*cli.App
	w io.Writer
}

func (a *app) init() {
	a.App = cli.NewApp()
	a.Name = "todo"
	a.Commands = []cli.Command{
		{
			Name:   "get",
			Action: a.getTasks,
		},
	}
}

func (a *app) printf(format string, as ...interface{}) {
	fmt.Fprintf(a.w, format, as...)
}

func (a *app) taskRepo() todo.TaskRepo {
	return new(memory.TaskRepo)
}
