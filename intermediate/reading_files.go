package intermediate

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("output.txt")
	if err != nil {
		fmt.Println("Error opening the file:", err)
		return
	}
	// If we open a file we shouldn't forget to close it
	defer func() {
		file.Close()
		fmt.Println("Closing open file")
	}()

	fmt.Println("File opened successfully")

	// Read the contents of the opened file
	// data := make([]byte, 1024)
	// _, err = file.Read(data)
	// if err != nil {
	// 	fmt.Println("Error reading data from file", err)
	// 	return
	// }
	// fmt.Println("File Content:", string(data))

	// We can also use bufio package to read line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println("Line: ", line)
	}
	err = scanner.Err()
	if err != nil {
		fmt.Println("Error reading the file", err)
		return
	}
}
