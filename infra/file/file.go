package file

import (
	"encoding/json"
	"fmt"
	"os"
)

func load(fname string, dst interface{}) error {
	src, err := os.Open(fname)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}

	if err := json.NewDecoder(src).Decode(dst); err != nil {
		return fmt.Errorf("failed to decode: %w", err)
	}

	return nil
}

func save(fname string, v interface{}) error {
	dst, err := os.OpenFile(fname, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}

	if err := json.NewEncoder(dst).Encode(v); err != nil {
		return fmt.Errorf("failed to encode: %w", err)
	}

	return nil
}

type status struct {
	Tasks []*task `json:"tasks"`
}
