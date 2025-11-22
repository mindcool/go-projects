package basics

import (
	"fmt"
	"maps"
)

func main() {
	// var mapVariable map[keyType]valueType
	// mapVariable = make(map[keyType]valueType)

	// using a Map Literal
	// mapVariable = map[keyType]valueType{
	// 	key1: value1,
	// 	key2: value2
	// }

	myMap := make(map[string]int)
	fmt.Println(myMap)

	// Lets assign key/value pairs
	myMap["key1"] = 9
	myMap["code"] = 18
	fmt.Println(myMap)
	fmt.Println(myMap["key1"])
	// If key doesnt exist 0 value return because that's an integer type value
	fmt.Println(myMap["key"])
	// That's how we assign new values to map
	myMap["code"] = 21
	fmt.Println(myMap)

	//That's how we delete a key
	delete(myMap, "key1")
	fmt.Println(myMap)
	myMap["key1"] = 9
	myMap["key2"] = 10
	myMap["key3"] = 11
	fmt.Println(myMap)
	// We can clear the map for a blank map
	clear(myMap)
	fmt.Println(myMap)
	// So since it is returning 0 if value does not find how am I know is this a real value or just unknown key
	value, isExist := myMap["key1"]
	fmt.Println("Does key exist:", isExist)
	fmt.Println("Returned Value: ", value)

	// Lets define map in a different way
	myMap2 := map[string]int{"a": 1, "b": 2}
	fmt.Println(myMap2)

	// Thats how we check equality between two different maps
	if maps.Equal(myMap, myMap2) {
		fmt.Println("myMap and myMap2 are equal")
	}

	// Lets iterate through the map
	for k, v := range myMap2 {
		fmt.Println(k, v)
	}

	// If a map is just defined but not initialized it is nil value map
	var myMap4 map[string]string

	if myMap4 == nil {
		fmt.Println("myMap4 defined but not initialized")
	}

	// So in go when we initialize a map without any value it initialize as nil
	// So we can't just assign a key and value instead we should use make function to
	// initialize it
	myMap4 = make(map[string]string)
	myMap4["key"] = "Value"
	fmt.Println(myMap4)

	// We can have nested maps
	myMap5 := make(map[string]map[string]string)
	myMap5["map1"] = myMap4
	fmt.Println(myMap5)
}
