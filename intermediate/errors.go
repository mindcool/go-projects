package intermediate

import (
	"errors"
	"fmt"
	"math"
)

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Math square root of a negative number")
	}
	// compute the square root
	return math.Sqrt(x), nil
}

func process(data []byte) error {
	if len(data) == 0 {
		return errors.New("Empty data")
	}
	// Process data if nothing is wrong we will just return nil
	return nil
}

func main() {
	// result, err := sqrt(16)

	// if err != nil {
	// 	fmt.Println("Error: ", err)
	// 	return
	// }
	// fmt.Println("Result: ", result)

	// result, err = sqrt(-16)

	// if err != nil {
	// 	fmt.Println("Error: ", err)
	// 	return
	// }
	// fmt.Println("Result: ", result)

	// data := []byte{}
	// Variable decleration and checking is on the same line
	// which prevents variable polution
	// if err := process(data); err != nil {
	// 	fmt.Println("Error: ", err)
	// 	return
	// }
	// fmt.Println("Data processed successfully")

	// if err1 := eprocess(); err1 != nil {
	// 	fmt.Println("Error: ", err1)
	// 	return
	// }
	// fmt.Println("Data processed successfully")

	err := readData()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Data read successfully")
}

type myError struct {
	message string
}

// In go we got error interface and that interface got a function named error
func (ex *myError) Error() string {
	return fmt.Sprintf("Error %s", ex.message)
}

func eprocess() error {
	return &myError{"Custom error message"}
}

func readConfig() error {
	return errors.New("Config error")
}

func readData() error {
	err := readConfig()
	if err != nil {
		// That's for formatting error
		return fmt.Errorf("readData: %w", err)
	}
	return nil
}
