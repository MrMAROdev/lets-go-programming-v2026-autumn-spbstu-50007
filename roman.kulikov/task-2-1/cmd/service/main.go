package main

import "fmt"

const (
	low = 15
	high = 30
)

func process() {
	var workers int

	if _, err := fmt.Scan(&workers); err != nil {
		return
	}

	minTemp := low
	maxTemp := high

	for range workers {
		var operator string

		var value int

		if _, err := fmt.Scan(&operator, &value); err != nil {
			return
		}

		switch operator {
		case ">=":
			if value > minTemp {
				minTemp = value
			}
		case "<=":
			if value < maxTemp {
				maxTemp = value
			}
		}

		if minTemp <= maxTemp {
			fmt.Println(minTemp)
		} else {
			fmt.Println(-1)
		}
	}
}

func main() {
	var departments int

	if _, err := fmt.Scan(&departments); err != nil {
		return
	}

	for range departments {
		process()
	}
}