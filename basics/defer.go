package basics

import "fmt"

func main() {
	// Defer is a mechanism that allows you to postpone the execution of a function
	//  until surronding function returns
	// process()
	processDeferOrder(10)
}

func process() {
	// Any statement or any function with defer keyword
	// will be defered till the end of the execution
	defer fmt.Println("Deferred statement executed")
	fmt.Println("Normal execution statement")
}

// The defer statement execution is LIFO (Last one executed first)
// Also if a defered statement or function getting a parameter
// Parameter will be passed at the time of execution
func processDeferOrder(i int) {
	defer fmt.Println("Deferred i value: ", i)
	defer fmt.Println("First Defer Stament")
	defer fmt.Println("Second Defer Stament")
	defer fmt.Println("Third Defer Stament")
	fmt.Println("Function Normal Execution Lifecycle")
	i++
	fmt.Println("I value: ", i)
	defer fmt.Println("Defered i value after increase: ", i)
}
