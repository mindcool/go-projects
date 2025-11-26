package basics

import "fmt"

func main() {

	// func <name>(parameters list) returnType {
	// code block
	// return value
	// }

	// In go convention if we are making a public function it will start with uppercase letter
	// and private functions will start with lowercase letter
	sum := add(1, 2)
	fmt.Println(sum)

	// That's an anonymous function and called immediately
	func() {
		fmt.Println("Hello Anonymous Function")
	}()
	// We can also create a function and assign it to a variable
	greet := func() {
		fmt.Println("Variable assigned anonymous function")
	}

	greet()

	// We can also assign named functions to variables
	operation := add
	result := operation(3, 5)
	fmt.Println(result)

	// Passing a function as argument
	resultFunc := applyOperation(5, 3, add)
	fmt.Println("5 + 3", resultFunc)

	// Returning and using a function
	multiplyBy2 := createMultiplier(2)
	fmt.Println("6 * 2 = ", multiplyBy2(6))
}

func add(a int, b int) int {
	return a + b
}

// Function that takes a function as an argument
func applyOperation(x int, y int, operation func(int, int) int) int {
	return operation(x, y)
}

// Function that returns a function
func createMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}
