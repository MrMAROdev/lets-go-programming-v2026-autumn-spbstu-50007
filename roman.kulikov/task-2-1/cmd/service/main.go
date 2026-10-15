package main

import (
	"fmt"
	"os"
)

const (
	minTemperature = 15
	maxTemperature = 30
)

func main() {
	var departmentsCount int
	if _, err := fmt.Scan(&departmentsCount); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	for i := 0; i < departmentsCount; i++ {
		var employeesCount int
		if _, err := fmt.Scan(&employeesCount); err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}

		low := minTemperature
		high := maxTemperature

		for j := 0; j < employeesCount; j++ {
			var sign string
			var temperatureValue int
			if _, err := fmt.Scan(&sign, &temperatureValue); err != nil {
				fmt.Fprintln(os.Stderr, err)
				return
			}

			switch sign {
			case ">=":
				if temperatureValue > low {
					low = temperatureValue
				}
			case "<=":
				if temperatureValue < high {
					high = temperatureValue
				}
			}

			if low <= high {
				fmt.Println(low)
			} else {
				fmt.Println(-1)
			}
		}
	}
}