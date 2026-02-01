package intermediate

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	user := os.Getenv("USER")
	home := os.Getenv("HOME")

	fmt.Println("User: ", user)
	fmt.Println("Home: ", home)

	err := os.Setenv("FRUIT", "APPLE")
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, env := range os.Environ() {
		kvpair := strings.SplitN(env, "=", 2)
		fmt.Println(kvpair[0])
	}
	err = os.Unsetenv("FRUIT")
	if err != nil {
		fmt.Println(err)
		return
	}

}
