package cfg

import (
	"errors"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

var errEmptyPaths = errors.New("input-file and output-file must be non-empty")

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func Load(path string) (Config, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return Config{}, fmt.Errorf("cfg.Load: %w", err)
	}

	var config Config
	if err := yaml.Unmarshal(raw, &config); err != nil {
		return Config{}, fmt.Errorf("cfg.Load: %w", err)
	}

	if config.InputFile == "" || config.OutputFile == "" {
		return Config{}, errEmptyPaths
	}

	return config, nil
}