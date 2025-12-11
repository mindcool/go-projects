package intermediate

import "fmt"

type Rectangle struct {
	length float64
	width  float64
}

// Method associated with embedded struct
//
//	will be promoted to outer struct directly
type Shape struct {
	Rectangle
}

// Method with value receiver
func (rec Rectangle) Area() float64 {
	return rec.length * rec.width
}

// We can use pointer receiver if we are going to modify values of it
// We can also use pointer receiver if we want to modify received instance
func (rec *Rectangle) Scale(factor float64) {
	rec.length *= factor
	rec.width *= factor
}

type MyInt int

// Method on a user-defined type
func (m MyInt) IsPositive() bool {
	return m > 0
}

// You don't need to create an instance
func (MyInt) welcomeMessage() string {
	return "Welcome to MyInt Type"
}

func main() {
	rect := Rectangle{
		length: 10,
		width:  9,
	}
	area := rect.Area()
	fmt.Println("Area of rectangle with width 9 and length 10 is: ", area)
	rect.Scale(10)
	area2 := rect.Area()
	fmt.Println("Area of rectangle with width 9 and length 10 is: ", area2)

	num := MyInt(-5)
	fmt.Println("Number is Positive", num.IsPositive())
	num = MyInt(9)
	fmt.Println("Number is Positive", num.IsPositive())
	fmt.Println(num.welcomeMessage())

	s := Shape{Rectangle: Rectangle{width: 10, length: 8}}
	fmt.Println(s.Area())
}
