package main

import (
	"fmt"
	"myproject/calculator" // Importing the correct package
)

func main() {
	num1, num2 := 10.0, 5.0

	fmt.Println("Addition:", calculator.Add(num1, num2))
	fmt.Println("Subtraction:", calculator.Subtract(num1, num2))
	fmt.Println("Multiplication:", calculator.Multiply(num1, num2))

	result, err := calculator.Divide(num1, num2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Division:", result)
	}
}
