package main

import (
	"fmt"
	"os"

	"github.com/tomocy/go-todo/presentation/cli"
)

func main() {
	app := cli.NewApp(os.Stdout)
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}