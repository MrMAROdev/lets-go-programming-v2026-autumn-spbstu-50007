package main

import (
	"encoding/json"
	"encoding/xml"
	"flag"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/net/html/charset"
	"gopkg.in/yaml.v3"
)

type Config struct {
	In  string `yaml:"input-file"`
	Out string `yaml:"output-file"`
}

type Valute struct {
	NumCode  string `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Value    string `xml:"Value"`
}

type ValuteData struct {
	XMLName xml.Name `xml:"ValCurs"`
	Valutes []Valute `xml:"Valute"`
}

type Currency struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	cfgPath := flag.String("config", "config.yaml", "")
	flag.Parse()

	raw, err := os.ReadFile(*cfgPath)
	check(err)

	var cfg Config
	check(yaml.Unmarshal(raw, &cfg))

	f, err := os.Open(cfg.In)
	check(err)
	defer f.Close()

	dec := xml.NewDecoder(f)
	dec.CharsetReader = charset.NewReaderLabel

	var data ValuteData
	check(dec.Decode(&data))

	currencies := make([]Currency, 0, len(data.Valutes))
	for _, v := range data.Valutes {
		num, err := strconv.Atoi(v.NumCode)
		check(err)
		val, err := strconv.ParseFloat(strings.ReplaceAll(v.Value, ",", "."), 64)
		check(err)
		currencies = append(currencies, Currency{num, v.CharCode, val})
	}

	sort.Slice(currencies, func(i, j int) bool {
		return currencies[i].Value > currencies[j].Value
	})

	out, err := json.MarshalIndent(currencies, "", "    ")
	check(err)

	check(os.MkdirAll(filepath.Dir(cfg.Out), 0o755))
	check(os.WriteFile(cfg.Out, out, 0o644))
}