package intermediate

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Name    string  `json:"name"`
	Age     int     `json:"age,omitempty"`
	Email   string  `json:"email"`
	Address Address `json:"address"`
}

type Address struct {
	City  string `json:"city"`
	State string `json:"state"`
}

func main() {
	persona := Person{Name: "John", Age: 30, Email: "sample@example.com"}
	personb := Person{Name: "Daria"}
	jsonData, err := json.Marshal(persona)
	if err != nil {
		fmt.Println("Error marshaling to JSON", err)
		return
	}
	fmt.Println(string(jsonData))
	jsonData, _ = json.Marshal(personb)
	fmt.Println(string(jsonData))

	newPerson := Person{Name: "Mary", Age: 33, Address: Address{City: "Istanbul", State: "Turkey"}}
	jsonData, _ = json.Marshal(newPerson)
	fmt.Println(string(jsonData))

	jData := `{"full_name":"Mike Murphy", "emp_id": "0009", "age": 30, "address":{"city":"San Jose", "state": "UT"}}`
	var EmployeeFromJson Employee
	err = json.Unmarshal([]byte(jData), &EmployeeFromJson)
	if err != nil {
		return
	}
	fmt.Println("Employee From Json", EmployeeFromJson)

	listOfCityState := []Address{
		{City: "New York", State: "NY"},
		{City: "San Jose", State: "CA"},
		{City: "Las Vegas", State: "NV"},
		{City: "Modesto", State: "CA"},
		{City: "Lehi", State: "UT"},
	}

	fmt.Println(listOfCityState)
	jsonList, err := json.Marshal(listOfCityState)
	if err != nil {
		return
	}
	fmt.Println(string(jsonList))
	// Handling unknown JSON structures
	jsonData2 := `{"name":"Memo", "age":30, "address":{"city":"American Fork", "state": "UT"}}`
	var data map[string]interface{}
	err = json.Unmarshal([]byte(jsonData2), &data)
	if err != nil {
		log.Fatalln("Error Unmarshalling Json: ", err)
	}
	fmt.Println("Decoded Unmarshalled JSON: ", data)
}

type Employee struct {
	FullName string  `json:"full_name"`
	EmpId    string  `json:"emp_id"`
	Age      int     `json:"age"`
	Address  Address `json:"address"`
}
