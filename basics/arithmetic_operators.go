package basics

import (
	"fmt"
	"math"
)

func main() {
	var a, b = 10, 3
	var result int
	result = a + b
	fmt.Println("Addition: ", result)

	result = a - b
	fmt.Println("Subtraction: ", result)

	result = a * b
	fmt.Println("Multiplication: ", result)

	// In go when you perform division between two integers the result also an integer
	// To get float result one of the operands should be float
	result = a / b
	fmt.Println("Division: ", result)

	result = a % b
	fmt.Println("Remainder: ", result)

	const operand float64 = 7

	const p float64 = 22 / operand
	fmt.Println(p)

	// Overflow with signed integer
	var maxInt int64 = 9223372036854775807 // max value of int64
	fmt.Println(maxInt)

	maxInt = maxInt + 1
	fmt.Println(maxInt)

	// Overflow with unsigned integers
	var uMaxInt uint64 = 18446744073709551615 // max value of uint64
	fmt.Println(uMaxInt)

	uMaxInt = uMaxInt + 1
	fmt.Println(uMaxInt)

	// Underflow with floating point numbers
	var smallFloat float64 = 1.0e-323
	fmt.Println(smallFloat)

	smallFloat = smallFloat / math.MaxFloat64
	fmt.Println(smallFloat)
}
