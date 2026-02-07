package intermediate

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	FirstName string `json:"first_name" db:"firstn" xml:"first"`
	LastName  string `json:"last_name,omitempty"`
	// Single - means that field will be omited from json
	Age int `json:"-"`
}

func main() {
	person := Person{FirstName: "Mike", Age: 30}
	jsonData, err := json.Marshal(person)
	if err != nil {
		log.Fatalln("Error marshalling struct: ", err)
	}
	fmt.Println(string(jsonData))
}
