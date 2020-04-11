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

func NewUser(id userID, name, email string, password password) (*user, error) {
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

func (u *user) setID(id userID) error {
	if id == "" {
		return fmt.Errorf("empty id")
	}

	u.id = id

	return nil
}

func (u *user) setName(name string) error {
	if name == "" {
		return fmt.Errorf("empty name")
	}

	u.name = name

	return nil
}

func (u *user) setEmail(email string) error {
	if email == "" {
		return fmt.Errorf("empty email")
	}

	u.profile.email = email

	return nil
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
	password password
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
