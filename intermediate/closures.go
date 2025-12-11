package intermediate

import "fmt"

func main() {
	// sequence := adder()
	// fmt.Println(sequence())
	// fmt.Println(sequence())
	// fmt.Println(sequence())
	// adder()
	// fmt.Println(sequence())
	subtractor := func() func(int) int {
		countdown := 99
		return func(x int) int {
			countdown -= x
			return countdown
		}
	}

	fmt.Println(subtractor()(9))
	fmt.Println(subtractor()(9))
}

func adder() func() int {
	i := 0
	fmt.Println("previous value of i: ", i)
	return func() int {
		i++
		fmt.Println("added 1 to i")
		return i
	}
}
