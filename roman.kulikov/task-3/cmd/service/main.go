package main

import (
	"flag"
	"fmt"
	"sort"

	"github.com/MrMAROdev/task-3/cfg"
	"github.com/MrMAROdev/task-3/parse"
	"github.com/MrMAROdev/task-3/save"
)

func main() {
	configPath := flag.String("config", "", "path to the YAML configuration file")
	flag.Parse()

	if *configPath == "" {
		panic("main: -config flag is required")
	}

	config, err := cfg.Load(*configPath)
	if err != nil {
		panic(fmt.Sprintf("main: %v", err))
	}

	rates, err := parse.DecodeXML(config.InputFile)
	if err != nil {
		panic(fmt.Sprintf("main: %v", err))
	}

	sort.Slice(rates, func(i, j int) bool { return rates[i].Value > rates[j].Value })

	if err := save.WriteJSON(config.OutputFile, rates); err != nil {
		panic(fmt.Sprintf("main: %v", err))
	}
}