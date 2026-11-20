package parse

import (
	"bytes"
	"encoding/xml"
	"os"
	"sort"
	"strconv"
	"strings"
)

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

func LoadAndSort(path string) []Currency {
	raw, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	raw = bytes.ReplaceAll(raw, []byte("windows-1251"), []byte("utf-8"))

	ascii := make([]byte, 0, len(raw))
	for _, b := range raw {
		if b < 0x80 {
			ascii = append(ascii, b)
		}
	}

	var data ValuteData
	if err := xml.Unmarshal(ascii, &data); err != nil {
		panic(err)
	}

	currencies := make([]Currency, 0, len(data.Valutes))
	for _, v := range data.Valutes {
		num, err := strconv.Atoi(strings.TrimSpace(v.NumCode))
		if err != nil {
			num = 0
		}
		val, err := strconv.ParseFloat(strings.ReplaceAll(strings.TrimSpace(v.Value), ",", "."), 64)
		if err != nil {
			val = 0
		}
		currencies = append(currencies, Currency{num, strings.TrimSpace(v.CharCode), val})
	}

	sort.Slice(currencies, func(i, j int) bool {
		return currencies[i].Value > currencies[j].Value
	})

	return currencies
}