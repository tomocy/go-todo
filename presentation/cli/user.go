package cli

import (
	"fmt"
	"strings"

	"github.com/tomocy/go-todo"
	"github.com/tomocy/go-todo/usecase"
	"github.com/urfave/cli"
)

func (a *app) createUser(ctx *cli.Context) error {
	u := usecase.NewCreateUser(a.userRepo())

	var (
		name  = ctx.String("name")
		email = ctx.String("email")
		pass  = ctx.String("password")
	)

	raw, err := u.Do(name, email, pass)
	if err != nil {
		return err
	}

	a.printf("User is successfully created.\n\n")
	a.printf("%v\n", user(*raw))

	return nil
}

func (a *app) authenticateUser(ctx *cli.Context) error {
	u := usecase.NewAuthenticateUser(a.userRepo(), a.sessionRepo())

	var (
		email = ctx.String("email")
		pass  = ctx.String("password")
	)

	raw, _, err := u.Do(email, pass)
	if err != nil {
		return err
	}

	a.printf("User is successfully authenticated.\n\n")
	a.printf("%v\n", user(*raw))

	return nil
}

type user todo.User

func (u user) String() string {
	var w strings.Builder
	raw := todo.User(u)

	fmt.Fprintf(&w, "ID: %s\n", raw.ID())
	fmt.Fprintf(&w, "Name: %s\n", raw.Name())
	fmt.Fprintf(&w, "Email: %s", raw.Email())

	return w.String()
}
