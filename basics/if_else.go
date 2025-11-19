package basics

import "fmt"

func main() {
	age := 25

	if age >= 18 {
		fmt.Println("You are an adult.")
	}

	temperature := 25
	if temperature >= 30 {
		fmt.Println("It is hot outside")
	} else {
		fmt.Println("It's cold outside")
	}
}
