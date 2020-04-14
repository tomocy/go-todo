package html

import (
	"flag"
	"fmt"
	"net/http"
)

type app struct {
	*http.ServeMux
	addr string
}

func (a *app) parseFlags(args []string) error {
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
