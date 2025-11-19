package basics

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// Iterate over a collection
	numbers := []int{1, 2, 3, 4, 5, 6}
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d \n", index, value)
	}

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("Odd Number: ", i)
		if i == 5 {
			break
		}
	}

	for i := range 10 {
		i++
		fmt.Println(i)
	}
}
