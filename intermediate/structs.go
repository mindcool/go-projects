package intermediate

import "fmt"

func main() {

	p := Person{
		firstName: "John",
		lastName:  "Doe",
		age:       30,
		address: Address{
			city:    "London",
			country: "U.K.",
		},
		PhoneHomeCell: PhoneHomeCell{
			cell: "801",
			home: "808",
		},
	}

	p1 := Person{
		firstName: "Jane",
		age:       25,
	}
	p1.address.city = "New York"
	p1.address.country = "USA"
	p1.PhoneHomeCell.cell = "Eben"

	fmt.Println(p.firstName)
	fmt.Println(p1.firstName)
	fmt.Println(p.fullName())
	fmt.Println(p.address.country)
	fmt.Println(p1.address.country)
	fmt.Println(p.cell)
	fmt.Println(p1.cell)
	fmt.Println("Struct comparison: ", p1 == p)

	// Anonymous Structs
	user := struct {
		username string
		email    string
	}{
		username: "user123",
		email:    "pseudoemail@example.org",
	}

	fmt.Println(user.username)
	fmt.Println("Before Increment", p.age)
	p.incrementAgeBy1()
	fmt.Println("After Increment", p.age)

}

type Person struct {
	firstName string
	lastName  string
	age       int
	address   Address
	PhoneHomeCell
}

type PhoneHomeCell struct {
	home string
	cell string
}

func (p Person) fullName() string {
	return p.firstName + " " + p.lastName
}

// p is a pointer to a person
// Pointer receivers allow the struct to actually modify the instance
func (p *Person) incrementAgeBy1() {
	p.age++
}

type Address struct {
	city    string
	country string
}
