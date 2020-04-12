package cli

import (
	"fmt"
	"time"

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

func (a *app) createTask(ctx *cli.Context) error {
	u := usecase.NewCreateTask(a.taskRepo(), a.sessionRepo())

	var (
		name       = ctx.String("name")
		rawDueDate = ctx.String("due-date")
	)
	var (
		dueDate time.Time
		err     error
	)
	if rawDueDate != "" {
		dueDate, err = time.Parse("2006/01/02", rawDueDate)
		if err != nil {
			return fmt.Errorf("failed to parse due date: %w", err)
		}
	}

	task, err := u.Do(name, dueDate)
	if err != nil {
		return err
	}

	a.printf("%v\n", task)

	return nil
}
