package intermediate

import "fmt"

func main() {
	// We will start with printing functions
	fmt.Print("Hello")
	fmt.Print("World!")
	fmt.Println(12, 456)

	fmt.Println("Hello")
	fmt.Println("World!")
	fmt.Println(12, 456)

	name := "John"
	age := 25
	fmt.Printf("Name: %s, Age: %d \n", name, age)
	fmt.Printf("Binary %b, Hex: %X\n", age, age)

	// Formatting functions
	// Formatting functions could also be used to concatenate strings
	s := fmt.Sprint("Hello ", "World!", 123, 456)
	fmt.Println(s)
	// Add spaces in between arguments and also add a new line at the end
	n := fmt.Sprintln("Hello ", "World!", 123, 456)
	fmt.Println(n)

	// Sprintf will return the value with special format
	sf := fmt.Sprintf("Name %s, Age: %d", name, age)
	fmt.Print(sf)
	fmt.Print(sf)

	// Scanning functions
	// Scan function scans imput from standard input and store in a valuable
	var name2 string
	var age2 int

	fmt.Println("Enter your name and age")
	// fmt.Scan(&name2, &age2)
	// Scanln will stop execution when it see a new line character
	fmt.Scanln(&name2, &age2)
	fmt.Printf("Name: %s Age %d \n", name2, age2)
	// We should enter the input the exact format we defined with scanf
	fmt.Scanf("%s %d", &name, &age)
	fmt.Printf("Name: %s, Age: %d\n", name, age)

	// Error formatting functions
	err := checkAge(15)
	fmt.Println(err)

}

func checkAge(age int) error {
	if age < 18 {
		return fmt.Errorf("Age %d is too young to drive.", age)
	}
	return nil
}
