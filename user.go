package todo

import (
	"context"
	"fmt"
)

type UserRepo interface {
	NextID(context.Context) (userID, error)
	FindByEmail(context.Context, string) (*user, error)
	Save(context.Context, *user) error
}

func NewUser(id userID, name, email, password string) (*user, error) {
	if name == "" {
		name = string(id)
	}
	if email == "" {
		return nil, fmt.Errorf("empty email")
	}
	if password == "" {
		return nil, fmt.Errorf("empty password")
	}

	return &user{
		id:   id,
		name: name,
		profile: profile{
			email: email,
		},
		cred: cred{
			password: password,
		},
	}, nil
}

type user struct {
	id      userID
	name    string
	status  userStatus
	profile profile
	cred    cred
}

type userID string

type userStatus int

const (
	userActive userStatus = iota
	userInactive
)

type profile struct {
	email string
}

type cred struct {
	password string
}

type password string
