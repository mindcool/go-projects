package basics

import "fmt"

func init() {
	fmt.Println("Initializing package (1)....")
}

func init() {
	fmt.Println("Initializing package (2)....")
}

func main() {
	fmt.Println("Inside the main function")
}

func init() {
	fmt.Println("Initializing package (3)....")
}
