package operations

import "fmt"

// Add returns the sum of two numbers.
func Add(a, b int) int {
	return a + b
}

// Subtract returns the difference between two numbers.
func Subtract(a, b int) int {
	return a - b
}

// Multiply returns the product of two numbers.
func Multiply(a, b int) int {
	return a * b
}

// Divide returns the quotient of two numbers, handling division by zero.
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero is not allowed")
	}
	return a / b, nil
}
