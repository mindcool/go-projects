package basics

import "fmt"

func main() {
	var i = 1
	for i <= 5 {
		fmt.Println("I will iterate forever: ", i)
		i++
	}

	sum := 0
	for {
		sum += 10
		fmt.Println("Sum: ", sum)
		if sum > 50 {
			break
		}
	}
}
