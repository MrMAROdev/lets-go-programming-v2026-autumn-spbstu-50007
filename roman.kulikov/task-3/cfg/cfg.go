package cfg

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	In  string `yaml:"input-file"`
	Out string `yaml:"output-file"`
}

func Load(path string) Config {
	raw, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	var cfg Config
	if err := yaml.Unmarshal(raw, &cfg); err != nil {
		panic(err)
	}

	if cfg.In == "" || cfg.Out == "" {
		panic("config: input-file and output-file must be set")
	}

	return cfg
}