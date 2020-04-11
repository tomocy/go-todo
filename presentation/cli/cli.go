package cli

import (
	"fmt"
	"io"
	"os"

	"github.com/tomocy/go-todo/infra/memory"

	"github.com/tomocy/go-todo"
	"github.com/tomocy/go-todo/infra/file"
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

const (
	envRepo           = "TODO_REPO"
	envStatusFilename = "TODO_STATUS_FILENAME"

	repoFile   = "file"
	repoMemory = "memory"
)

func (a *app) taskRepo() todo.TaskRepo {
	repo, ok := os.LookupEnv(envRepo)
	if !ok {
		repo = repoMemory
	}

	switch repo {
	case repoFile:
		fname, ok := os.LookupEnv(envStatusFilename)
		if !ok {
			fname = "./"
		}

		return file.NewTaskRepo(fname)
	default:
		return new(memory.TaskRepo)
	}
}
