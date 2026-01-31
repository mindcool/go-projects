package intermediate

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("output.txt")
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	// Keyword to filter lines
	keyword := "Lorem"
	// Lets read and filter lines
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, keyword) {
			updatedLine := strings.ReplaceAll(line, keyword, "Not So Lodash")
			fmt.Println("Filtered line: ", line)
			fmt.Println("Updated Line: ", updatedLine)
		}
	}
	err = scanner.Err()
	if err != nil {
		fmt.Println("Error scanning file:", err)
		return
	}
}
