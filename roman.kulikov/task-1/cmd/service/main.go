package main

import "fmt"

func main() {
	var argument1, argument2 int
	var operat string

	if _, exception := fmt.Scan(&argument1); exception != nil {
		fmt.Println("Invalid first operand")
		return
	}
	if _, exception := fmt.Scan(&argument2); exception != nil {
		fmt.Println("Invalid second operand")
		return
	}
	if _, exception := fmt.Scan(&operat); exception != nil {
		fmt.Println("Invalid operation")
		return
	}

	switch operat {
	case "+":
		fmt.Println(argument1 + argument2)
	case "-":
		fmt.Println(argument1 - argument2)
	case "*":
		fmt.Println(argument1 * argument2)
	case "/":
		if argument2 != 0 {
			fmt.Println(argument1 / argument2)
		} else {
			fmt.Println("Division by zero")
		}
	default:
		fmt.Println("Invalid operation")
	}
}
