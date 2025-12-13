package intermediate

import "fmt"

// Generics function definition
func swap[T any](a, b T) (T, T) {
	return b, a
}

// Generics struct definition
type Stack[T any] struct {
	// That's a slice holding of elements type any
	elements []T
}

func (s *Stack[T]) push(element T) {
	s.elements = append(s.elements, element)
}

func (s *Stack[T]) pop() (T, bool) {
	if len(s.elements) == 0 {
		var zero T
		return zero, false
	}

	element := s.elements[len(s.elements)-1]
	// Last element will be excluded
	s.elements = s.elements[:len(s.elements)-1]
	return element, true
}

func (s *Stack[T]) isEmpty() bool {
	return len(s.elements) == 0
}

func (s Stack[T]) printAll() {
	if len(s.elements) == 0 {
		fmt.Println("There is no element")
	} else {
		fmt.Println(s.elements)
	}
}

func main() {
	x, y := 1, 2
	x, y = swap(x, y)
	fmt.Println("X Value: ", x)
	fmt.Println("Y Value: ", y)

	x1, y1 := "Ahmet", "Memed"
	x1, y1 = swap(x1, y1)
	fmt.Println("X Value: ", x1)
	fmt.Println("Y Value: ", y1)

	intStack := Stack[int]{}
	intStack.push(1)
	intStack.push(5)
	intStack.push(10)
	intStack.printAll()
	fmt.Println(intStack.pop())
	intStack.printAll()
	fmt.Println(intStack.pop())
	fmt.Println("Is Empty: ", intStack.isEmpty())
	fmt.Println(intStack.pop())
	fmt.Println("Is Empty: ", intStack.isEmpty())

	stringStack := Stack[string]{}
	stringStack.push("Hello")
	stringStack.push("Go Lang")
	stringStack.push("Generics")
	stringStack.printAll()
	fmt.Println(stringStack.pop())
	stringStack.printAll()
	fmt.Println(stringStack.pop())
	fmt.Println("Is Empty: ", stringStack.isEmpty())
	fmt.Println(stringStack.pop())
	fmt.Println("Is Empty: ", stringStack.isEmpty())
}
