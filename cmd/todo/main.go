package main

import (
	"fmt"
	"os"

	"github.com/tomocy/go-todo/presentation/cli"
	"github.com/tomocy/go-todo/presentation/html"
)

func main() {
	app := newRunner()
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

const (
	modeCLI  = "cli"
	modeHTML = "html"
)

func newRunner() runner {
	mode, ok := os.LookupEnv("TODO_MODE")
	if !ok {
		mode = modeCLI
	}

	switch mode {
	case modeHTML:
		return html.New(os.Stdout)
	default:
		return cli.New(os.Stdout)
	}
}

type runner interface {
	Run([]string) error
}
