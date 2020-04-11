package cli

import (
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
