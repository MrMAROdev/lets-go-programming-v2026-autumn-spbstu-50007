package main

import (
	"flag"

	"github.com/MrMAROdev/task-3/cfg"
	"github.com/MrMAROdev/task-3/parse"
	"github.com/MrMAROdev/task-3/save"
)

func main() {
	cfgPath := flag.String("config", "config.yaml", "")
	flag.Parse()

	config := cfg.Load(*cfgPath)
	parsed := parse.LoadAndSort(config.In)

	currencies := make([]save.Currency, len(parsed))
	for i, c := range parsed {
		currencies[i] = save.Currency{
			NumCode:  c.NumCode,
			CharCode: c.CharCode,
			Value:    c.Value,
		}
	}

	save.ToJSON(currencies, config.Out)
}