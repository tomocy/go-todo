package file

import (
	"encoding/json"
	"fmt"
	"os"
)

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
