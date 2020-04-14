package html

import "net/http"

type app struct {
	*http.ServeMux
}
