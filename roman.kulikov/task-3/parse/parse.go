package parse

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"golang.org/x/text/encoding/charmap"
)

var (
	errNoRates      = errors.New("no currencies found")
	errUnsupported  = errors.New("unsupported xml encoding")
)

type Currency struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

type xmlValute struct {
	NumCode  string `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Value    string `xml:"Value"`
}

type xmlRoot struct {
	XMLName xml.Name    `xml:"ValCurs"`
	Valutes []xmlValute `xml:"Valute"`
}

func DecodeXML(path string) ([]Currency, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("parse.DecodeXML: %w", err)
	}
	defer func() { _ = f.Close() }()

	dec := xml.NewDecoder(f)
	dec.CharsetReader = func(cs string, r io.Reader) (io.Reader, error) {
		if strings.EqualFold(cs, "windows-1251") {
			return charmap.Windows1251.NewDecoder().Reader(r), nil
		}
		return nil, fmt.Errorf("%w: %s", errUnsupported, cs)
	}

	var root xmlRoot
	if err := dec.Decode(&root); err != nil {
		return nil, fmt.Errorf("parse.DecodeXML: %w", err)
	}
	if len(root.Valutes) == 0 {
		return nil, errNoRates
	}

	rates := make([]Currency, 0, len(root.Valutes))
	for _, v := range root.Valutes {
		num := 0
		if s := strings.TrimSpace(v.NumCode); s != "" {
			if n, e := strconv.Atoi(s); e == nil {
				num = n
			}
		}
		val := 0.0
		if s := strings.TrimSpace(v.Value); s != "" {
			if x, e := strconv.ParseFloat(strings.ReplaceAll(s, ",", "."), 64); e == nil {
				val = x
			}
		}
		rates = append(rates, Currency{NumCode: num, CharCode: strings.TrimSpace(v.CharCode), Value: val})
	}
	if len(rates) == 0 {
		return nil, errNoRates
	}
	return rates, nil
}
