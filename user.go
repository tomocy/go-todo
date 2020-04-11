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
	u := new(user)

	if err := u.setID(id); err != nil {
		return nil, err
	}
	if name == "" {
		name = string(id)
	}
	if err := u.setName(name); err != nil {
		return nil, err
	}
	if err := u.setEmail(email); err != nil {
		return nil, err
	}
	if err := u.setPassword(password); err != nil {
		return nil, err
	}

	return u, nil
}

type user struct {
	id       userID
	name     string
	email    string
	password password
	status   userStatus
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

	u.email = email

	return nil
}

func (u *user) Password() password {
	return u.password
}

func (u *user) setPassword(pass password) error {
	if pass == "" {
		return fmt.Errorf("empty password")
	}

	u.password = pass

	return nil
}

type userID string

type userStatus int

const (
	userActive userStatus = iota
	userInactive
)

func HashPassword(p string) (password, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return password(hashed), nil
}

type password string

func (p password) IsSame(other string) bool {
	return bcrypt.CompareHashAndPassword([]byte(p), []byte(other)) == nil
}
