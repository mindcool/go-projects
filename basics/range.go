package basics

import "fmt"

func main() {
	message := "Hello World!"
	// This will show index and unicode character number for the letter
	for i, v := range message {
		// fmt.Println("Index: ", i)
		// fmt.Println("Variable: ", v)
		// We call characters rune in go and %c is for character
		fmt.Printf("Index %d, Rune %c \n", i, v)
	}

	// Range iterate over the underlying object by creating a copy of it
}
