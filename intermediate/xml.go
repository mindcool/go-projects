package intermediate

import (
	"encoding/xml"
	"fmt"
	"log"
)

type Person struct {
	// xml name field is a special field used by encoding xml package
	XMLName xml.Name `xml:"person"`
	Name    string   `xml:"name"`
	Age     int      `xml:"age,omitempty"`
	City    string   `xml:"city,omitempty"`
	Email   string   `xml:"email"`
	Address Address  `xml:"address"`
}

type Address struct {
	State   string `xml:"state"`
	Address string `xml:"address"`
}

// <book isbn="fhjfh123442" color="blue">
type Book struct {
	XMLName    xml.Name `xml:"book"`
	ISBN       string   `xml:"isbn,attr"`
	Title      string   `xml:"title,attr"`
	Author     string   `xml:"author,attr"`
	Pseudo     string   `xml:"pseudo"`
	PseudoAttr string   `xml:"pseudoattr,attr"`
}

func main() {
	person := Person{Name: "Mark", Age: 30, City: "Mesa", Email: "email@example.com", Address: Address{State: "AZ", Address: "345 N 500 W"}}
	xmlData, err := xml.Marshal(person)
	if err != nil {
		log.Fatalln("Error Marshalling data into XML: ", err)
	}
	fmt.Println("XML Data:", string(xmlData))

	xmlData1, err := xml.MarshalIndent(person, "", " ")
	if err != nil {
		log.Fatalln("Error Marshalling data into XML: ", err)
	}
	fmt.Println("XML Data with Indent", string(xmlData1))

	// Unmarshall XML
	xmlRaw := `<person><name>Mark</name><age>30</age><city>Mesa</city><email>email@example.com</email><address><state>AZ</state><address>345 N 500 W</address></address></person>`
	var personxml Person
	err = xml.Unmarshal([]byte(xmlRaw), &personxml)
	if err != nil {
		log.Fatalln("Error Unmarshalling XML: ", err)
	}
	fmt.Println("Unmarshalled Object:", personxml)

	book := Book{
		ISBN:       "978-1098139292",
		Title:      "Learning Go: An Idiomatic Approach to Real-World Go Programming",
		Author:     "Jon Bodner",
		Pseudo:     "Pseudo",
		PseudoAttr: "Pseudo Attribute",
	}
	xmlDataWithAttr, err := xml.MarshalIndent(book, "", "  ")
	if err != nil {
		log.Fatalln("Error marshalling data")
	}
	fmt.Println(string(xmlDataWithAttr))
}
