package main

import (
	"fmt"
)

type Employee struct {
	ID     int
	Name   string
	Salary float64
}

// Method to update salary
func (e *Employee) UpdateSalary(newSalary float64) {
	e.Salary = newSalary
}

// Method to display employee details
func (e Employee) Display() {
	fmt.Printf("ID: %d, Name: %s, Salary: %.2f\n", e.ID, e.Name, e.Salary)
}

// Function to add an employee to the database
func AddEmployee(emp *Employee, db map[int]*Employee) {
	db[emp.ID] = emp
}

// Function to remove an employee from the database
func RemoveEmployee(id int, db map[int]*Employee) {
	delete(db, id)
}

func main() {
	// Employee database
	employeeDB := make(map[int]*Employee)

	// Adding employees
	emp1 := &Employee{ID: 1, Name: "John Doe", Salary: 50000}
	emp2 := &Employee{ID: 2, Name: "Jane Smith", Salary: 60000}
	AddEmployee(emp1, employeeDB)
	AddEmployee(emp2, employeeDB)

	// Displaying all employees
	fmt.Println("Employee List:")
	for _, emp := range employeeDB {
		emp.Display()
	}

	// Updating salary
	emp1.UpdateSalary(55000)
	fmt.Println("\nAfter Salary Update:")
	emp1.Display()

	// Removing an employee
	RemoveEmployee(2, employeeDB)
	fmt.Println("\nAfter Removing Employee 2:")
	for _, emp := range employeeDB {
		emp.Display()
	}
}
