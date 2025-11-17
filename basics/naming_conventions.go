package basics

import "fmt"

type Employee struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	// Pascal case
	// PascalCase looks like this
	// Ex. CalculateArea, UserInfo, NewHttpRequest
	// Structs, interfaces, enums ...

	// Snake case
	// snake_case looks like this
	// Ex. user_id, first_name, http_request
	// Snake case commonly used for variables, file names

	// UPPERCASE
	// Upper case used explicity named for constants
	// PI, AREAMULTIPLIER
	// Constants are immutable

	// Mix Case
	// Ex javaScript, htmlDocument, isValid
	// Dealing with external libraries or following other language(s) consistancy

	// Package names always should be without underscores and should always be lowercase
	const MAXENTRIES = 5
	var employeeID = 1001
	fmt.Println("EmployeeID: ", employeeID)
}
