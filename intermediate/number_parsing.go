package intermediate

import (
	"fmt"
	"strconv"
)

func main() {
	numStr := "12345"
	num, err := strconv.Atoi(numStr)
	if err != nil {
		fmt.Println("Error parsing the value:", err)
	}
	fmt.Println("Parsed Integer: ", num)
	fmt.Println("Parsed Integer: ", num+1)

	// Second argument is base and third one is bytes of the number
	numiStr, err := strconv.ParseInt(numStr, 10, 32)
	if err != nil {
		fmt.Println("Error parsing the value:", err)
	}
	fmt.Println("Parsed Integer: ", numiStr)

	floatstr := "3.14"
	floatval, err := strconv.ParseFloat(floatstr, 64)
	if err != nil {
		fmt.Println("Error parsin value: ", err)
	}
	fmt.Printf("Parsed Float: %.2f\n", floatval)

	binaryStr := "1010"
	decimal, err := strconv.ParseInt(binaryStr, 2, 64)
	if err != nil {
		fmt.Println("Error parsing binary value: ", err)
		return
	}
	fmt.Println("Parsed binary to decimal: ", decimal)

	hexStr := "FF"
	hexadecimal, err := strconv.ParseInt(hexStr, 16, 32)
	if err != nil {
		fmt.Println("Error parsing hexadecimal value: ", err)
		return
	}
	fmt.Println("Parsed hexadecimal to decimal: ", hexadecimal)

	invalidNum := "12abc"
	invalidParse, err := strconv.Atoi(invalidNum)
	if err != nil {
		fmt.Println("Error parsin value: ", err)
	}
	fmt.Println("Parsed string to integer: ", invalidParse)
}
