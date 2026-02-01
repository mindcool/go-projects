package intermediate

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// fmt.Println("Command: ", os.Args[0])
	// fmt.Println("Argument 1: ", os.Args[1])
	// fmt.Println("Argument 1: ", os.Args[2])
	// fmt.Println("Argument 1: ", os.Args[3])
	// fmt.Println("Argument 1: ", os.Args[4])
	// fmt.Println("Argument 1: ", os.Args[5])
	// Lets loop through arguments
	for idx, arg := range os.Args {
		fmt.Println("Argument ", idx, ":", arg)
	}

	// Define flags
	var name string
	var age int
	var male bool

	flag.StringVar(&name, "name", "John", "Name of the User")
	flag.IntVar(&age, "age", 40, "Age of the User")
	flag.BoolVar(&male, "sex", true, "is user male?")

	// Lets parse command line arguments
	flag.Parse()

	fmt.Println("Name: ", name)
	fmt.Println("Age: ", age)
	fmt.Println("Male: ", male)
}
