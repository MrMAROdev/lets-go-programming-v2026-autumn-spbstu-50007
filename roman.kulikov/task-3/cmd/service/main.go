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

	c := cfg.Load(*cfgPath)
	parsed := parse.LoadAndSort(c.In)

	currencies := make([]save.Currency, len(parsed))
	for i, p := range parsed {
		currencies[i] = save.Currency{NumCode: p.NumCode, CharCode: p.CharCode, Value: p.Value}
	}

	save.ToJSON(currencies, c.Out)
}