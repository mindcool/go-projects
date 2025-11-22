package basics

import "fmt"

func main() {

	// var arrayName [size]elementType
	var numbers [5]int
	fmt.Println(numbers)

	numbers[4] = 20
	fmt.Println(numbers)

	numbers[0] = 9
	fmt.Println(numbers)

	// Lets create an array of string values
	fruits := [4]string{"Apple", "Banana", "Orange", "Grapes"}
	fmt.Println("Fruit array: ", fruits)

	fmt.Println("Third element: ", fruits[2])

	// In go arrays are value types so when we pass an array it creates a copy of that array
	originalArray := [3]int{1, 2, 3}
	copiedArray := originalArray

	copiedArray[0] = 100

	fmt.Println("Original Array: ", originalArray)
	fmt.Println("Copied Array: ", copiedArray)

	for i := 0; i < len(numbers); i++ {
		fmt.Println("Element at index, ", i, ":", numbers[i])
	}

	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	// We can discard unnedded value by using _ (underscore) which is blank identifier
	for _, value := range numbers {
		fmt.Printf("Value: %d\n", value)
	}

	// Usage for blank identifier
	a, _ := someFunction()
	fmt.Println(a)
	// fmt.Println(b)

	// Underscore is blank identifier used to store unused values
	b := 2
	_ = b

	// Length of an array
	fmt.Println("The length of numbers array is", len(numbers))

	// Compare arrays
	array1 := [3]int{1, 2, 3}
	array2 := [3]int{1, 2, 3}
	fmt.Println("Array 1 is equal to Arrray 2:", array1 == array2)

	//Array of arrays
	var matrix [3][3]int = [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(matrix)

	// Lets understand pointers and addresses
	originalArray1 := [3]int{1, 2, 3}
	// Copied Array 1 is a pointer to an array with 3 integers
	var copiedArray1 *[3]int
	// Getting the address of originalArray1 and assign it to the pointer
	copiedArray1 = &originalArray1

	copiedArray1[0] = 100

	fmt.Println("Original Array: ", originalArray1)
	fmt.Println("Copied Array: ", copiedArray1)
}

func someFunction() (int, int) {
	return 1, 2
}
