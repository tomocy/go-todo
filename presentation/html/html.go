package html

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/tomocy/go-todo"
	"github.com/tomocy/go-todo/infra/file"
	"github.com/tomocy/go-todo/infra/memory"
)

func New(w io.Writer) *app {
	return &app{
		ServeMux: http.NewServeMux(),
		w:        w,
	}
}

type app struct {
	*http.ServeMux
	w    io.Writer
	addr string
}

func (a *app) Run(args []string) error {
	if err := a.parse(args); err != nil {
		return fmt.Errorf("failed to parse: %w", err)
	}

	a.register()

	a.printf("listen and serve on %s\n", a.addr)
	if err := http.ListenAndServe(a.addr, a); err != nil {
		return fmt.Errorf("failed to listen and serve: %w", err)
	}

	return nil
}

func (a *app) parse(args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("too less arguments")
	}

	flags := flag.NewFlagSet(args[0], flag.ContinueOnError)
	addr := flags.String("addr", ":80", "the address to listen and serve")

	if err := flags.Parse(args[1:]); err != nil {
		return err
	}

	a.addr = *addr

	return nil
}

func (a *app) register() {
	a.HandleFunc("/users", a.users)
}

func (a *app) printf(format string, as ...interface{}) {
	fmt.Fprintf(a.w, format, as...)
}

const (
	envRepo           = "TODO_REPO"
	envStatusFilename = "TODO_STATUS_FILENAME"

	repoFile   = "file"
	repoMemory = "memory"

	defaultStatusFilename = "./status.json"
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
			fname = defaultStatusFilename
		}

		return file.NewUserRepo(fname)
	default:
		return memory.NewUserRepo()
	}
}
