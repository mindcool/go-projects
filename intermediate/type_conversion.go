package intermediate

import "fmt"

func main() {
	var a int = 32
	b := int32(a)
	c := float64(b)

	e := 3.14
	f := int(e)
	fmt.Println(f, c)

	// Syntax for conversion type(value)
	g := "Hello"
	var h []byte
	h = []byte(g)
	fmt.Println(h)
	i := []byte{255}
	j := string(i)
	fmt.Println("Converted i: ", j)
}
