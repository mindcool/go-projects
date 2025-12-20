package intermediate

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// rand.Intn will generate numbers in an open interval [0 n]
	// So this one will create numbers from 1 to 100
	fmt.Println(rand.Intn(100) + 1)
	// We can also give it a fixed seed
	val := rand.New(rand.NewSource(48))
	// Since we seeded with 48 it will always give the same number
	fmt.Println(val.Intn(101))
	// fmt.Println(rand.New(rand.NewSource(48)).Intn(101))
	// We can also use unix time as seed
	valUnixTimeSeed := rand.New(rand.NewSource(time.Now().Unix()))
	fmt.Println(valUnixTimeSeed.Intn(101))
	// We can also create random float numbers
	fmt.Println(rand.Float64()) // between 0.0 and 1.0 0.0 included and 1.0 excluded

	for {
		// Show the menu
		fmt.Println("Welcome to the Dice Game!")
		fmt.Println("1. Roll the dice")
		fmt.Println("2. Exit")
		fmt.Print("Enter your choice (1 or 2)")

		var choice int
		_, err := fmt.Scan(&choice)
		if err != nil || (choice != 1 && choice != 2) {
			fmt.Println("Invalid choice please choose 1 or 2")
			continue
		}
		if choice == 2 {
			fmt.Println("Thanks for playing! Goodbye")
			break
		}
		die1 := rand.Intn(6) + 1
		die2 := rand.Intn(6) + 1

		// Show the result
		fmt.Printf("You rolled a %d and a %d\n", die1, die2)
		fmt.Println("Total: ", die1+die2)

	}
}
