package basics

import "fmt"

func main() {
	fruit := "apple"

	switch fruit {
	case "apple":
		fmt.Println("It's an apple.")
	case "banana":
		fmt.Println("It's a banana")
	default:
		fmt.Println("Unknown Fruit!")
	}

	// We can also use multiple case for comparison
	day := "Monday"

	switch day {
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("It's a weekday")
	case "Sunday":
		fmt.Println("It's a weekend")
	}

	number := 15

	switch {
	case number < 10:
		fmt.Println("Number is less than 10")
	case number >= 10 && number < 20:
		fmt.Println("Number is between 10 and 19")
	default:
		fmt.Println("Number is 20 or more")
	}

	// Let see how fallthrough works
	num := 2
	switch {
	case num > 1:
		fmt.Println("Greater than 1")
		// Normally it will stop here but
		// if we want it to keep going we can use fallthrough
		fallthrough
	case num == 2:
		fmt.Println("Number is 2")
	default:
		fmt.Println("Not Two")
	}
	checkType(10)
	checkType(2.12)
	checkType("Hello")
	checkType(true)
}

// We can't use fallthrough for type interference
func checkType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("It's an integer")
	case float64:
		fmt.Println("It's a float")
	case string:
		fmt.Println("It's a string")
	default:
		fmt.Println("Unknown Type")
	}
}
