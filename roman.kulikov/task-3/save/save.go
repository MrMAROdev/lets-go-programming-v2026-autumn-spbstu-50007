package save

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/MrMAROdev/task-3/parse"
)

const (
	dirPerm  = 0o755
	filePerm = 0o644
)

func WriteJSON(path string, rates []parse.Currency) error {
	data, err := json.MarshalIndent(rates, "", "    ")
	if err != nil {
		return fmt.Errorf("save.WriteJSON: %w", err)
	}
	if err := os.MkdirAll(filepath.Dir(path), dirPerm); err != nil {
		return fmt.Errorf("save.WriteJSON: %w", err)
	}
	if err := os.WriteFile(path, data, filePerm); err != nil {
		return fmt.Errorf("save.WriteJSON: %w", err)
	}
	return nil
}
