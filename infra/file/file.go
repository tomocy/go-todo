package file

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func generateRandomString(length int) string {
	s := make([]byte, length)
	for i := range s {
		s[i] = letters[rand.Int63n(int64(length))]
	}

	return string(s)
}

func load(fname string) (status, error) {
	if err := initIfNotExist(fname); err != nil {
		return status{}, fmt.Errorf("failed to init file: %w", err)
	}

	src, err := os.OpenFile(fname, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		return status{}, fmt.Errorf("failed to open file: %w", err)
	}

	var dst status
	if err := json.NewDecoder(src).Decode(&dst); err != nil {
		return status{}, fmt.Errorf("failed to decode: %w", err)
	}

	return dst, nil
}

func initIfNotExist(fname string) error {
	if _, err := os.Stat(fname); err == nil {
		return nil
	}

	return save(fname, status{})
}

func save(fname string, src status) error {
	dst, err := os.OpenFile(fname, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}

	if err := json.NewEncoder(dst).Encode(src); err != nil {
		return fmt.Errorf("failed to encode: %w", err)
	}

	return nil
}

type status struct {
	Users []*user `json:"users"`
	Tasks []*task `json:"tasks"`
}

func (s *status) addUser(u *user) {
	for i, added := range s.Users {
		if added.ID == u.ID {
			s.Users[i] = u
			return
		}
	}

	s.Users = append(s.Users, u)
}

func (s *status) addTask(t *task) {
	for i, added := range s.Tasks {
		if added.ID == t.ID {
			s.Tasks[i] = t
			return
		}
	}

	s.Tasks = append(s.Tasks, t)
}
