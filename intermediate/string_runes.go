package intermediate

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// String can be created by using double quotes or backticks
	message := "Hello, \nGo!"
	message1 := "Hello \tGo!"
	message2 := "Hello \rGo!"
	// When we use backticks it is raw string
	// Backtick one will discard any escape sequence and threat everything like characters
	rawMessage := `Hello\nGo`

	fmt.Println(message)
	fmt.Println(message1)
	fmt.Println(message2)
	fmt.Println(rawMessage)

	fmt.Println("Length of message variable is", len(message))

	fmt.Println("The first character in message var is", message[0]) // ASCII value of first character

	greeting := "Hello "
	name := "Alice"
	msg := greeting + name
	fmt.Println(msg)

	str1 := "Apple"
	str2 := "Banana"
	// Go do string comparison according to byte values of characters
	fmt.Println(str1 < str2)

	// String iteration
	for i, char := range message {
		fmt.Printf("Character at index %d is %c\n", i, char)
	}

	for _, char := range message {
		fmt.Printf("%x\n", char)
	}

	// Lets get the rune count in a string
	fmt.Println("Rune count: ", utf8.RuneCountInString(greeting))

	greetingWithName := greeting + name
	fmt.Println(greetingWithName)

	var ch rune = 'a'
	var trchar rune = 'Ç'

	// ASCII value of characters
	fmt.Println(ch)
	fmt.Println(trchar)

	fmt.Printf("%c\n", ch)
	fmt.Printf("%c\n", trchar)

	// Converted string
	cstr := string(trchar)
	fmt.Println(cstr)

	// How to understand format of cstr
	fmt.Printf("Type of cstr %T\n", cstr)

	jhello := "こんにちは"
	for _, runeValue := range jhello {
		fmt.Printf("%c\n", runeValue)
	}
	r := '😊'
	fmt.Printf("%v\n", r)
	fmt.Printf("%c\n", r)
}
