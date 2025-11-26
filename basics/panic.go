package basics

import "fmt"

func main() {
	// In go panic is a built in function
	// that stops normal execution of a function immediately
	// Panic let us pass any argument type, interface{} is type any
	// panic(interface{})
	// Example of a valid input
	// process(10)
	// process(-10)
	deferProcess(-10)
}

func process(input int) {
	if input < 0 {
		panic("input must be a positive number")
	}
	fmt.Println("Processing input: ", input)
}

// Defer execute when function return a value
// but it executes even after function is panic
func deferProcess(input int) {
	defer fmt.Println("Deferred 1")
	defer fmt.Println("Deferred 2")
	if input < 0 {
		fmt.Println("Before panic")
		panic("input must be a positive number")
		// Anything after panic won't be executed even if it is defer
		// defer fmt.Println("Deferred 3")
	}
	fmt.Println("Processing input: ", input)
}
