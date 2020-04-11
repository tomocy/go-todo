package todo

import (
	"context"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type UserRepo interface {
	NextID(context.Context) (UserID, error)
	FindByEmail(context.Context, string) (*User, error)
	Save(context.Context, *User) error
}

func NewUser(id UserID, name, email string, password password) (*User, error) {
	u := new(User)

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

type User struct {
	id       UserID
	name     string
	email    string
	password password
	status   userStatus
}

func (u *User) ID() UserID {
	return u.id
}

func (u *User) setID(id UserID) error {
	if id == "" {
		return fmt.Errorf("empty id")
	}

	u.id = id

	return nil
}

func (u *User) setName(name string) error {
	if name == "" {
		return fmt.Errorf("empty name")
	}

	u.name = name

	return nil
}

func (u *User) Email() string {
	return u.email
}

func (u *User) setEmail(email string) error {
	if email == "" {
		return fmt.Errorf("empty email")
	}

	u.email = email

	return nil
}

func (u *User) Password() password {
	return u.password
}

func (u *User) setPassword(pass password) error {
	if pass == "" {
		return fmt.Errorf("empty password")
	}

	u.password = pass

	return nil
}

type UserID string

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
