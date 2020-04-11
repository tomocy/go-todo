package cli

import (
	"github.com/tomocy/go-todo/usecase"
	"github.com/urfave/cli"
)

func (a *app) getTasks(ctx *cli.Context) error {
	u := usecase.NewGetTasks(a.taskRepo())

	tasks, err := u.Do()
	if err != nil {
		return err
	}

	if len(tasks) < 1 {
		a.printf("Nothing to do\n")
		return nil
	}

	a.printf("TODOs\n")
	for _, t := range tasks {
		a.printf("%v\n", t)
	}

	return nil
}
