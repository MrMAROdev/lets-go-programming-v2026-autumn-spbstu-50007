package save

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Currency struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

func ToJSON(currencies []Currency, path string) {
	out, err := json.MarshalIndent(currencies, "", "    ")
	if err != nil {
		panic(err)
	}

	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		panic(err)
	}

	if err := os.WriteFile(path, out, 0o644); err != nil {
		panic(err)
	}
}