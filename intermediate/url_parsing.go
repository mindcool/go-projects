package intermediate

import (
	"fmt"
	"net/url"
)

func main() {
	// [scheme://][userinfo@]host[:port][/path][?query][#fragment]
	rawUrl := "https://example.com:8080/path?query=param#fragment"
	parsedUrl, err := url.Parse(rawUrl)
	if err != nil {
		fmt.Println("Error parsing URL: ", err)
		return
	}
	fmt.Println("Scheme:", parsedUrl.Scheme)
	fmt.Println("Host:", parsedUrl.Host)
	fmt.Println("Port:", parsedUrl.Port())
	fmt.Println("Path:", parsedUrl.Path)
	fmt.Println("Raw Query:", parsedUrl.RawQuery)
	fmt.Println("Fragment:", parsedUrl.Fragment)

	rawUrl1 := "https://example.com/path?name=John&age=38"
	parseUrl1, err := url.Parse(rawUrl1)
	if err != nil {
		fmt.Println("Error parsing URL: ", err)
		return
	}
	queryParams := parseUrl1.Query()
	fmt.Println(queryParams)
	fmt.Println("Name: ", queryParams.Get("name"))
	fmt.Println("Age: ", queryParams.Get("age"))

	// ======= BUILDING URLs =======
	fmt.Println("\n--- Building URLs ---")

	// Method 1: Build URL from scratch using url.URL struct
	// This creates a URL object with individual components
	baseUrl := &url.URL{
		Scheme: "https",       // Protocol (http, https, ftp, etc.)
		Host:   "example.com", // Domain name (can include port like "example.com:8080")
		Path:   "/api/users",  // Path part of URL
	}
	// At this point: baseUrl = "https://example.com/api/users"

	// Adding query parameters to the URL object
	query := baseUrl.Query()  // Returns an empty url.Values map
	query.Set("name", "John") // Set adds or replaces a parameter
	query.Set("age", "42")
	baseUrl.RawQuery = query.Encode() // IMPORTANT: Must encode and assign back to RawQuery
	fmt.Println("Built URL: ", baseUrl.String())
	// Output: https://example.com/api/users?name=John&age=42

	// Method 2: Using url.Values separately (more flexible)
	// url.Values is just a map[string][]string - it can hold multiple values per key
	values := url.Values{}

	// Add() can add multiple values for the same key
	values.Add("name", "Arda")
	values.Add("age", "42")
	values.Add("city", "Istanbul")
	values.Add("country", "Turkey")
	// You could do values.Add("hobby", "coding") and values.Add("hobby", "gaming")
	// to have multiple hobby values

	// Encode() converts url.Values to a query string
	encodedQuery := values.Encode()
	fmt.Println("Encoded Query String:", encodedQuery)
	// Output: age=42&city=Istanbul&country=Turkey&name=Arda (alphabetically sorted)

	// Method 3: Combine base URL string + query string
	baseUrl1 := "https://example.com/search"
	fullUrl := baseUrl1 + "?" + encodedQuery
	fmt.Println("Full URL from string concatenation:", fullUrl)

	// Method 4: More complex URL with all components
	fmt.Println("\n--- Complete URL Example ---")
	completeUrl := &url.URL{
		Scheme:   "https",
		User:     url.UserPassword("admin", "secret"), // Optional: username:password
		Host:     "api.example.com:443",               // Host with port
		Path:     "/v1/search",
		Fragment: "results", // The #fragment part
	}
	params := url.Values{}
	params.Set("q", "golang tutorials")
	params.Set("limit", "10")
	completeUrl.RawQuery = params.Encode()

	fmt.Println("Complete URL:", completeUrl.String())
	// Output: https://admin:secret@api.example.com:443/v1/search?limit=10&q=golang+tutorials#results
}
