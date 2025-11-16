package basics

import "fmt"

// Global variables can't use gofor notation, go for notation can only be used for local variables
// If you try to use gofor notation as global declaration it will say expected declaration
var middleName = "Cane"

func main() {
	// Uninitialize variable
	// var age int
	// Type of variable is optional if we are initializing the variable
	var name string = "John"
	var name1 = "Jane"

	// We can also use gofor syntax to make assignment easier
	// count := 10
	lastName := "Smith"

	fmt.Println(middleName + name + name1 + lastName)
	// Default types
	// Numeric Types: 0
	// Boolean Types: false
	// String type: ""
	// Pointers, slices, maps, functions and structs: nil

	// Scope variables in go has block scope

}
