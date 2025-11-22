package basics

import (
	"fmt"
	"slices"
)

func main() {
	// Slices declerared just like arrays but do not have a size
	// var sliceName[]ElementType

	// numbers variable is a slices it is kind of an array but doesnt have a fixed length
	// var numbers []int
	// var numbers1 = []int{1, 2, 3}

	// numbers2 := []int{9, 8, 7}

	// We are initializing a slice with a capacity of five
	// slice := make([]int, 5)

	a := [5]int{1, 2, 3, 4, 5}
	// 2,3 and 4 {2,3,4} will be included in this slice
	slice1 := a[1:4]
	fmt.Println(slice1)

	// We can append more elements to slice
	slice1 = append(slice1, 6, 7, 8)
	fmt.Println("Slice 1:", slice1)

	// Lets make a copy of the slice

	var sliceCopy = make([]int, len(slice1))
	copy(sliceCopy, slice1)
	fmt.Println("Slice Copy: ", sliceCopy)

	// Slices also have a concept of nil slices,
	// nil slices has zero value and not reference any underlying array it is actually blank

	// var nilSlice []int

	for i, v := range slice1 {
		fmt.Println(i, v)
	}

	// We can access values of a slice using the index
	fmt.Println("Element of index 3 of slice1", slice1[3])

	// We can also manipulate values of slices
	slice1[3] = 50

	// That's how we check equality of slices
	if slices.Equal(slice1, sliceCopy) {
		fmt.Println("slice 1 is equal to sliceCopy")
	}

	// Multi dimensional slices
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println(twoD)

	// Slices support slice operator
	// slice[low:high]
	slice2 := slice1[2:4]
	fmt.Println(slice2)

	fmt.Println("The capacity of slice2 is", cap(slice2))
	fmt.Println("The len of slice2 is", len(slice2))

}
