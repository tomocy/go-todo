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

func NewUser(id UserID, name, email string, password Password) (*User, error) {
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

func RecoverUser(
	id UserID, name, email string,
	password Password, status UserStatus,
) *User {
	return &User{
		id:       id,
		name:     name,
		email:    email,
		password: password,
		status:   status,
	}
}

type User struct {
	id       UserID
	name     string
	email    string
	password Password
	status   UserStatus
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

func (u *User) Name() string {
	return u.name
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

func (u *User) Password() Password {
	return u.password
}

func (u *User) setPassword(pass Password) error {
	if pass == "" {
		return fmt.Errorf("empty Password")
	}

	u.password = pass

	return nil
}

func (u *User) Status() UserStatus {
	return u.status
}

type UserID string

type UserStatus int

const (
	userActive UserStatus = iota
	userInactive
)

func HashPassword(p string) (Password, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return Password(hashed), nil
}

type Password string

func (p Password) IsSame(other string) bool {
	return bcrypt.CompareHashAndPassword([]byte(p), []byte(other)) == nil
}
