package intermediate

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(strings.NewReader("Hello, Bufio packageeee!\n"))

	// Reading byte slice
	data := make([]byte, 20)
	n, err := reader.Read(data)
	if err != nil {
		fmt.Println("Error reading: ", err)
		return
	}
	fmt.Printf("Read %d bytes: %s\n", n, data[:n])

	// Dont forget thats a second call of readstring
	// so it will read what first one left of
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading string: ", err)
		return
	}
	fmt.Println(line)

	writer := bufio.NewWriter(os.Stdout)
	dataSlice := []byte("Hello bufio packagedddddcccc!\n")
	num, err := writer.Write(dataSlice)
	if err != nil {
		fmt.Println("Error writing: ", err)
		return
	}
	fmt.Printf("Wrote %d bytes \n", num)
	// Flush the data to ensure all the data is written to os.Stdout
	err = writer.Flush()
	if err != nil {
		fmt.Println("Error flushing writer:", err)
		return
	}

	// Writing string
	var str string = "This is a string.\n"
	n, err = writer.WriteString(str)
	if err != nil {
		fmt.Println("Error writing string", err)
		return
	}
	fmt.Printf("Wrote %d bytes\n", n)
	// Flush the buffer
	err = writer.Flush()
	if err != nil {
		fmt.Println("Error flushing writer: ", err)
		return
	}
}
