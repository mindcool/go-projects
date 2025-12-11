package intermediate

import "fmt"

func main() {
	// Pointers is a variable that store memory address of another variable
	// ptr is a pointer which stores a memory address of a variable which is type int
	// Type of pointer must match the type of variable it examines
	var ptr *int
	var a int = 10
	// Now ptr is holding the address of integer a
	ptr = &a
	fmt.Println(a)
	fmt.Println(ptr)
	fmt.Println("Dereference pointer value:", *ptr)
	// &a is directly memory address of the variable
	fmt.Println(&a)
	// Pointer arithmetics does not support in go
	// So pointer operations supported in go are limited with referencing and dereferencing
	modifyValue(ptr)
	fmt.Println(a)
}

func modifyValue(ptr *int) {
	*ptr++

}
