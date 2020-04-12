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
			Name: "user",
			Subcommands: []cli.Command{
				{
					Name: "create",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name: "name",
						},
						cli.StringFlag{
							Name:     "email",
							Required: true,
						},
						cli.StringFlag{
							Name:     "password",
							Required: true,
						},
					},
					Action: a.createUser,
				},
				{
					Name:      "authenticate",
					ShortName: "authn",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:     "email",
							Required: true,
						},
						cli.StringFlag{
							Name:     "password",
							Required: true,
						},
					},
					Action: a.authenticateUser,
				},
			},
		},
		{
			Name: "task",
			Subcommands: []cli.Command{
				{
					Name:   "get",
					Action: a.getTasks,
				},
			},
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

func (a *app) userRepo() todo.UserRepo {
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

		return file.NewUserRepo(fname)
	default:
		return new(memory.UserRepo)
	}
}

func (a *app) sessionRepo() todo.SessionRepo {
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

		return file.NewSessionRepo(fname)
	default:
		return memory.NewSessionRepo()
	}
}

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
