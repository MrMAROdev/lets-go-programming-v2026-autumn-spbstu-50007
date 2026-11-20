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
	currencies := parse.LoadAndSort(config.In)
	save.ToJSON(currencies, config.Out)
}