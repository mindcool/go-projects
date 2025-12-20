package intermediate

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

func main() {
	str := "Hello Go!"

	fmt.Println("Length of String", len(str))

	str1 := "Hello"
	str2 := "World!"
	result := str1 + " " + str2
	fmt.Println(result)

	// We can access the letters by using indexes
	// We got the ASCII value of alphabet which is at index 0
	fmt.Println(str[0])
	// We can get substring of a string last index excluded
	fmt.Println(str[1:5])

	// String conversion
	num := 18
	str3 := strconv.Itoa(num)
	fmt.Println("Converted to String: ", str3)
	fmt.Println(len(str3))

	// Strings splitting
	fruits := "Apple, Orange, Banana"
	parts := strings.Split(fruits, ",")
	fmt.Println(parts)

	// As we have split we also got join
	countries := []string{"America", "Sweden", "Bulgaria"}
	joined := strings.Join(countries, ",")
	fmt.Println(joined)

	// Does a string contains certain word
	fmt.Println(strings.Contains(str, "Hell"))

	// Replace an occurence of a substring with another substring
	fmt.Println(strings.Replace(str, "Go", "Noyan", 1))

	// Lets trim leading and trailing white space from a string
	stringwithspace := " Hello Everyone"
	fmt.Println(stringwithspace)
	fmt.Println(strings.TrimSpace(stringwithspace))

	// Change string to lower and upper case
	fmt.Println(strings.ToLower(stringwithspace))
	fmt.Println(strings.ToUpper(stringwithspace))

	// We can also repeat a string
	fmt.Println(strings.Repeat("Boo Yeah ", 3))

	// We can also count the occurence of a substring
	fmt.Println(strings.Count(stringwithspace, "H"))
	fmt.Println(strings.Count(stringwithspace, "l"))

	// We can check prefix or suffix in a string
	fmt.Println(strings.HasPrefix(stringwithspace, " He"))
	fmt.Println(strings.HasSuffix("Hello", "lo"))

	// Go offer us a regular expression package which let us manipulate strings
	// Based on complex rules
	str4 := "Hello, 123 Go! 11"
	// Regular expression needs to be in backtick to be considered raw litterals
	re := regexp.MustCompile(`\d+`)
	// We are going to pass -1 which means we want all the matches
	fmt.Println("All regular expression matches", re.FindAllString(str4, -1))

	str5 := "Tanılama, tanımlama ve izleme amaçlarıyla kullanılabilen dil örneği analizi"
	// To count characters in a unicode string
	fmt.Println(utf8.RuneCountInString(str5))

	// Lets learn string builder
	var builder strings.Builder

	// Write some strings
	builder.WriteString("Hello")
	builder.WriteString(", ")
	builder.WriteString("world!")
	// So string builder is building a string up until we are ready it won't
	// convert it into a immutable string
	buildedStr := builder.String()
	fmt.Println(buildedStr)

	// Using write rune to add a character
	builder.WriteRune(' ')
	builder.WriteString("How are you")
	buildedStr = builder.String()
	fmt.Println(buildedStr)

	// Reset the builder to build a new string
	builder.Reset()
	builder.WriteString("Starting fresh!")
	buildedStr = builder.String()
	fmt.Println(buildedStr)
}
