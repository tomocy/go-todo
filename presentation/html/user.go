package html

import (
	"fmt"
	"net/http"
)

func (a *app) users(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	default:
		http.Error(w, fmt.Sprint(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}
