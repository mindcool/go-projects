package intermediate

import (
	"errors"
	"fmt"
)

func main() {
	// err := doSomething()
	// if err != nil {
	// 	fmt.Print(err)
	// 	return
	// }
	// fmt.Println("Operation completed successfully!")
	err1 := doSomethingWrapped()
	if err1 != nil {
		fmt.Print(err1)
		return
	}
	fmt.Println("Operation completed successfully!")
}

type customError struct {
	code    int
	message string
}

type customErrorWrapped struct {
	code    int
	message string
	err     error
}

// Error returns the error message. Implementing Error() method of error interface
func (e *customError) Error() string {
	return fmt.Sprintf("Error %d: %s \n", e.code, e.message)
}

// This new interface will also implement error method
func (e *customErrorWrapped) Error() string {
	return fmt.Sprintf("Error %d: %s, %v \n", e.code, e.message, e.err)
}

// Function that return a custom error
func doSomething() error {
	return &customError{
		code:    500,
		message: "Something went wrong!",
	}
}

func doSomethingWrapped() error {
	err := doSomethingElse()
	if err != nil {
		return &customErrorWrapped{
			code:    500,
			message: "Wrapped Error Message",
			err:     err,
		}
	}
	return nil
}

func doSomethingElse() error {
	return errors.New("Internal Error!")
}
