package cfg

import (
	"os"

	"go.yaml.in/yaml/v3"
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
	var c Config
	if err := yaml.Unmarshal(raw, &c); err != nil {
		panic(err)
	}
	if c.In == "" || c.Out == "" {
		panic("config: input-file and output-file must be set")
	}
	return c
}