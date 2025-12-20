package intermediate

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	// Unixtime is in seconds
	unixTime := now.Unix()
	fmt.Println("Go Time: ", now)
	fmt.Println("Unix Time: ", unixTime)
	t := time.Unix(unixTime, 0)
	fmt.Println("Converted Unix Time:", t)
	fmt.Println("Formatted Time: ", t.Format("2006-01-02"))
}
