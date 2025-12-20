package intermediate

import "fmt"

func main() {
	num := 42
	// 5 is minimum width that number will have
	fmt.Printf("%05d\n", num)

	message := "Hello"
	// Right aligned 10 chars minimum string
	fmt.Printf("|%10s|\n", message)
	// Left aligned 10 chars minimum string
	fmt.Printf("|%-10s|\n", message)

	message1 := "Hello \nWorld!"
	// That's a raw string
	message2 := `Hello \nWorld!`
	fmt.Println(message1)
	fmt.Println(message2)
}
