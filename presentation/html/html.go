package html

import (
	"flag"
	"fmt"
	"io"
	"net/http"
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

func (a *app) printf(format string, as ...interface{}) {
	fmt.Fprintf(a.w, format, as...)
}
