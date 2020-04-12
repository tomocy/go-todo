package cli

import (
	"fmt"
	"strings"
	"time"

	"github.com/tomocy/go-todo"
	"github.com/tomocy/go-todo/usecase"
	"github.com/urfave/cli"
)

func (a *app) getTasks(ctx *cli.Context) error {
	u := usecase.NewGetTasks(a.taskRepo(), a.sessionRepo())

	tasks, err := u.Do()
	if err != nil {
		return err
	}

	if len(tasks) < 1 {
		a.printf("Nothing to do\n")
		return nil
	}

	a.printf("TODOs\n\n")
	for i, t := range tasks {
		a.printf("%v\n", task(*t))
		if i != len(tasks)-1 {
			a.printf("----------\n")
		}
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

	raw, err := u.Do(name, dueDate)
	if err != nil {
		return err
	}

	a.printf("Task is successfully created.\n\n")
	a.printf("%v\n", task(*raw))

	return nil
}

func (a *app) changeDueDate(ctx *cli.Context) error {
	u := usecase.NewChangeDueDate(a.taskRepo(), a.sessionRepo())

	var (
		id         = todo.TaskID(ctx.String("id"))
		rawDueDate = ctx.String("due-date")
	)
	dueDate, err := time.Parse("2006/01/02", rawDueDate)
	if err != nil {
		return fmt.Errorf("failed to parse due date: %w", err)
	}

	raw, err := u.Do(id, dueDate)
	if err != nil {
		return err
	}

	a.printf("Task is successfully configured.\n\n")
	a.printf("%v\n", task(*raw))

	return nil
}

func (a *app) cancelDueDate(ctx *cli.Context) error {
	u := usecase.NewCancelDueDate(a.taskRepo(), a.sessionRepo())

	id := todo.TaskID(ctx.String("id"))

	raw, err := u.Do(id)
	if err != nil {
		return err
	}

	a.printf("Task is successfully configured.\n\n")
	a.printf("%v\n", task(*raw))

	return nil
}

func (a *app) postponeTask(ctx *cli.Context) error {
	u := usecase.NewPostponeTask(a.taskRepo(), a.sessionRepo())

	id := todo.TaskID(ctx.String("id"))

	raw, err := u.Do(id)
	if err != nil {
		return err
	}

	a.printf("Task is successfully postponed.\n\n")
	a.printf("%v\n", task(*raw))

	return nil
}

func (a *app) deleteTask(ctx *cli.Context) error {
	u := usecase.NewDeleteTask(a.taskRepo(), a.sessionRepo())

	id := todo.TaskID(ctx.String("id"))

	if err := u.Do(id); err != nil {
		return err
	}

	a.printf("Task is successfully deleted.\n")

	return nil
}

type task todo.Task

func (t task) String() string {
	var w strings.Builder
	raw := todo.Task(t)

	fmt.Fprintf(&w, "ID: %s\n", raw.ID())
	fmt.Fprintf(&w, "Name: %s\n", raw.ID())
	if !raw.DueDate().IsZero() {
		fmt.Fprintf(&w, "DueDate: %s\n", raw.DueDate().Format("2006/01/02"))
	}
	fmt.Fprintf(&w, "PostponedTimes: %s", raw.ID())

	return w.String()
}
