package cbr

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
	ErrNoRates     = errors.New("no currencies found")
	ErrUnsupported = errors.New("unsupported xml encoding")
)

type Rate struct {
	NumCode  int     `json:"num_code" xml:"NumCode"`
	CharCode string  `json:"char_code" xml:"CharCode"`
	Value    float64 `json:"value" xml:"Value"`
}

type xmlRoot struct {
	XMLName xml.Name `xml:"ValCurs"`
	Valutes []Rate   `xml:"Valute"`
}

func windows1251Reader(charset string, input io.Reader) (io.Reader, error) {
	if strings.EqualFold(charset, "windows-1251") {
		return charmap.Windows1251.NewDecoder().Reader(input), nil
	}

	return nil, fmt.Errorf("%w: %s", ErrUnsupported, charset)
}

func Parse(path string) ([]Rate, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("cbr.Parse: %w", err)
	}

	defer func() { _ = file.Close() }()

	decoder := xml.NewDecoder(file)
	decoder.CharsetReader = windows1251Reader

	var root xmlRoot
	if err := decoder.Decode(&root); err != nil {
		return nil, fmt.Errorf("cbr.Parse: %w", err)
	}

	if len(root.Valutes) == 0 {
		return nil, ErrNoRates
	}

	rates := make([]Rate, 0, len(root.Valutes))
	for _, item := range root.Valutes {
		rates = append(rates, Rate{
			NumCode:  toInt(item.NumCode),
			CharCode: strings.TrimSpace(item.CharCode),
			Value:    toFloat(item.Value),
		})
	}

	if len(rates) == 0 {
		return nil, ErrNoRates
	}

	return rates, nil
}

func toInt(raw string) int {
	trimmed := strings.TrimSpace(raw)
	if trimmed == "" {
		return 0
	}

	if n, err := strconv.Atoi(trimmed); err == nil {
		return n
	}

	return 0
}

func toFloat(raw string) float64 {
	trimmed := strings.TrimSpace(raw)
	if trimmed == "" {
		return 0
	}

	if f, err := strconv.ParseFloat(strings.ReplaceAll(trimmed, ",", "."), 64); err == nil {
		return f
	}

	return 0
}
