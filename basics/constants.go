package basics

import "fmt"

// We can define with constant keyword and let compiler interpret
const pi = 3.14
const GRAVITY = 9.81

func main() {
	// We can also give data type
	const days int = 7

	// That's a constant block
	const (
		monday        = 1
		tuesday       = 2
		wednesday     = 3
		thursday      = 4
		friday    int = 5
	)
	fmt.Println("Pi: ", pi)
	fmt.Println("Gravity: ", GRAVITY)
	fmt.Println("Days: ", days)
	fmt.Println("Monday: ", monday)
}
