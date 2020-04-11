package todo

import (
	"context"
	"fmt"

	"golang.org/x/crypto/bcrypt"
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

func HashPassword(p string) (password, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return password(hashed), nil
}

type password string

func (p password) isSame(other string) bool {
	return bcrypt.CompareHashAndPassword([]byte(p), []byte(other)) == nil
}
