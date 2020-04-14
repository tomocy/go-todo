package html

import (
	"fmt"
	"net/http"

	"github.com/tomocy/go-todo/usecase"
)

func (a *app) users(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		a.createUser(w, r)
	default:
		http.Error(w, fmt.Sprint(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

func (a *app) createUser(w http.ResponseWriter, r *http.Request) {
	u := usecase.NewCreateUser(a.userRepo())

	if err := r.ParseForm(); err != nil {
		a.printf("failed to parse form: %s\n", err)
		http.Error(w, fmt.Sprintf("failed to parse form: %s", err), http.StatusBadRequest)
		return
	}
	var (
		name  = r.FormValue("name")
		email = r.FormValue("email")
		pass  = r.FormValue("password")
	)

	user, err := u.Do(name, email, pass)
	if err != nil {
		a.printf("failed to create user: %s\n", err)
		http.Error(w, fmt.Sprintf("failed to create user: %s", err), http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, user)
}
