package parse

import (
	"encoding/xml"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/text/encoding/charmap"
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
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	dec := xml.NewDecoder(f)
	dec.CharsetReader = func(cs string, r io.Reader) (io.Reader, error) {
		if strings.EqualFold(cs, "windows-1251") {
			return charmap.Windows1251.NewDecoder().Reader(r), nil
		}
		return r, nil
	}

	var data ValuteData
	if err := dec.Decode(&data); err != nil {
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