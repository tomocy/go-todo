package cli

import (
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

	user, err := u.Do(name, email, pass)
	if err != nil {
		return err
	}

	a.printf("%v\n", user)

	return nil
}
