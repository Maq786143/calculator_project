package main

import (
	"fmt"
)

type Temperature struct {
	Value float64
}

// Method to convert Celsius to Fahrenheit
func (t *Temperature) ToFahrenheit() {
	t.Value = (t.Value * 9 / 5) + 32
}

// Method to convert Fahrenheit to Celsius
func (t *Temperature) ToCelsius() {
	t.Value = (t.Value - 32) * 5 / 9
}

func main() {
	// Example usage
	temp := &Temperature{Value: 25} // 25°C
	fmt.Printf("Original Temperature: %.2f°C\n", temp.Value)

	temp.ToFahrenheit()
	fmt.Printf("Converted to Fahrenheit: %.2f°F\n", temp.Value)

	temp.ToCelsius()
	fmt.Printf("Converted back to Celsius: %.2f°C\n", temp.Value)
}
