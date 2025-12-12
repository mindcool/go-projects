package main

import (
	"fmt"
	"math"
)

// Since geometry is small case interface
// it won't be exported from package main
type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type rect1 struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.height * r.width
}

func (c circle) area() float64 {
	return math.Pi * c.radius * math.Pow(c.radius, 2)
}

func (r rect) perim() float64 {
	return 2 * (r.height + r.width)
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func (c circle) diameter() float64 {
	return 2 * c.radius
}

func (r rect1) area() float64 {
	return r.height * r.width
}

func (r rect1) perim() float64 {
	return 2 * (r.height + r.width)
}

// This one accepts objects matching with geometry interface
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

// That's a variadic parameter for interface
func myPrinter(i ...interface{}) {
	for _, v := range i {
		fmt.Println(v)
	}
}

func printType(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("This is an integer", v)
	case string:
		fmt.Println("This is a string", v)
	default:
		fmt.Println("Type Unknown", v)
	}

}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}
	r1 := rect1{width: 3, height: 4}
	// Since both radius and circle implementing geometry interface
	measure(r)
	measure(c)
	// Since rect1 missing a method this is not valid use for it
	// For rect1 to be used for geometry it should implement perim method
	// If we make it implement perim it is ok to pass for geometry interface
	measure(r1)
	myPrinter(1, "John", "Hello")

	printType(9)
	printType("Hello")
	printType(false)
}
