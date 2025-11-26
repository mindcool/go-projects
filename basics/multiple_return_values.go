package basics

import (
	"errors"
	"fmt"
)

func main() {

	// func functionName(parameter1 type1, parameter2 type2, ...)(returnType1, returnType2, ...) {
	// code block
	// return returnValue1, returnValue2
	// }
	q, r := divide(10, 3)
	fmt.Printf("Quotient: %d, Remainder %d\n", q, r)
	result, err := compare(3, 3)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println(result)
	}
}

// We are getting two return values from our function divide
func divide(a int, b int) (int, int) {
	quotient := a / b
	remainder := a % b
	return quotient, remainder
}

func compare(a, b int) (string, error) {
	if a > b {
		return "a is greater than b", nil
	} else if b > a {
		return "b is greater than a", nil
	} else {
		return "", errors.New("Unable to compare which is greater")
	}
}

// We can also use named return variables
func div(a, b int) (quotient int, remainder int) {
	quotient = a / b
	remainder = a % b
	// Since we named return variables go compiler smart enough to know what needs to be returned
	return
}
