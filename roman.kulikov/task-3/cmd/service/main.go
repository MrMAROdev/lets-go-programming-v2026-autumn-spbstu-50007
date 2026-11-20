package main

import (
	"flag"
	"sort"

	"github.com/MrMAROdev/lets-go-programming-v2026-autumn-spbstu-50007/roman.kulikov/task-3/cbr"
	"github.com/MrMAROdev/lets-go-programming-v2026-autumn-spbstu-50007/roman.kulikov/task-3/config"
	"github.com/MrMAROdev/lets-go-programming-v2026-autumn-spbstu-50007/roman.kulikov/task-3/output"
)

func run() error {
	configPath := flag.String("config", "", "path to the YAML configuration file")
	flag.Parse()

	if *configPath == "" {
		return config.ErrEmptyPaths
	}

	cfg, err := config.Load(*configPath)
	if err != nil {
		return err
	}

	rates, err := cbr.Parse(cfg.InputFile)
	if err != nil {
		return err
	}

	sort.SliceStable(rates, func(i, j int) bool {
		return rates[i].Value > rates[j].Value
	})

	return output.Write(cfg.OutputFile, rates)
}

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}
