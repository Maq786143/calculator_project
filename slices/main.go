package main

import "fmt"

func main() {
	// Initialize the slice with the initial values
	numbers := []int{10, 20, 30, 40, 50}

	// Manually expand the slice by creating a new slice with one extra element
	newNumbers := make([]int, len(numbers)+1)
	copy(newNumbers, numbers)

	// Add 60 to the new slice
	newNumbers[len(newNumbers)-1] = 60

	// Loop through the slice using a traditional for loop
	for i := 0; i < len(newNumbers); i++ {
		fmt.Println("Index:", i, "Value:", newNumbers[i])
	}

	// Print the length and capacity of the new slice
	fmt.Println("Length:", len(newNumbers))
	fmt.Println("Capacity:", cap(newNumbers))
}
