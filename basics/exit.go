package basics

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("Deferred statement")
	// So calling os.exit will not invoke any defer functions including those registered using defer
	// Exit will exit with a status code anything other than 0 means abrupt termination
	fmt.Println("Starting the main function")
	// Exit with status code of 1
	os.Exit(1)
	// This will never be exacuted
	fmt.Println("End of main function")
}
