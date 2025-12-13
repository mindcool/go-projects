package intermediate

import "fmt"

type person struct {
	name string
	age  int
}

// Employee struct embed the fields and methods of person
type employee struct {
	person // Embedded struct
	empId  string
	salary float64
}

func (p person) introduce() {
	fmt.Println("Name: ", p.name, "Age: ", p.age)
}

func (e employee) introduce() {
	fmt.Println("Name: ", e.name, "Age: ", e.age, "Emp ID", e.empId, "salary", e.salary)
}

func main() {
	emp := employee{
		person: person{name: "Noy", age: 43},
		empId:  "123",
		salary: 10000,
	}

	fmt.Println("Name: ", emp.name) // Accessing the embedded struct field
	fmt.Println("Age: ", emp.age)
	fmt.Println("Employee ID: ", emp.empId)
	fmt.Println("Employee Salary", emp.salary)
	emp.introduce()
}
