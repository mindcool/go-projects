package basics

import "fmt"

func main() {
	// ... Ellipsis indicates that function could accept zero or more arguments
	// func functionName(param1 type1, param2 type2, param3 ...type3) returnType{}
	// Thats a variadic function and variadic function contains variadic parameters
	fmt.Println("Sum of 1,2,3 and 4", sum(1, 2, 3, 4))

	numbers := []int{1, 2, 3, 4, 5}
	// We can pass numbers slice as parameter to sum function
	fmt.Println(sum(numbers...))
}

func sum(nums ...int) int {
	total := 0
	for _, v := range nums {
		total += v
	}
	return total
}
