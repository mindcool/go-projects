package intermediate

import (
	"bufio"
	"fmt"
	"html/template"
	"os"
	"strings"
)

func main() {
	// We are going to create a template and parameter it takes is name of the template
	// tmpl := template.New("example")
	// tmpl, err := tmpl.Parse("Welcome, {{.name}}! How are you doing?")
	// Instead of creating template and parsin it we mostly do everything in a single line
	tmpl, err := template.New("example").Parse("Welcome, {{.name}}! How are you doing?\n")
	if err != nil {
		panic(err)
	}

	// Define data for the welcome message template
	data := map[string]interface{}{
		"name": "John",
	}

	err = tmpl.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}

	// We can also use template.Must which will automatically panic and handle the error
	tmpl1 := template.Must(template.New("example").Parse("Welcome, {{.name}}! How are you doing?\n"))
	err = tmpl1.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}

	// Take inputs from user we use os.Stdin since it is standard input to console
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name: ")
	// That's how we read user's input and since it is accepting byte as delimiter, delimeter will be define with single quote
	// It will return a string and an error if something goes wrong
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	// Define named templates for different types of
	templates := map[string]string{
		"welcome":      "Welcome, {{.name}}! We're glad you joined.",
		"notification": "{{.name}}, you have a new notification: {{.notification}}",
		"error":        "Ooops! An error occured {{.errorMessage}}",
	}

	// Parse and store templates
	parsedTemplates := make(map[string]*template.Template)
	for tempName, templ := range templates {
		parsedTemplates[tempName] = template.Must(template.New(tempName).Parse(templ))
	}

	for {
		// Show menu
		fmt.Println("\nMenu:")
		fmt.Println("1. Join")
		fmt.Println("2. Get Notification")
		fmt.Println("3. Get error")
		fmt.Println("4. Exit")
		fmt.Println("Choose an option:")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		var data map[string]interface{}
		var tmpl *template.Template

		switch choice {
		case "1":
			tmpl = parsedTemplates["welcome"]
			data = map[string]interface{}{"name": name}
		case "2":
			fmt.Println("Enter your notification message: ")
			notification, _ := reader.ReadString('\n')
			notification = strings.TrimSpace(notification)
			tmpl = parsedTemplates["notification"]
			data = map[string]interface{}{"name": name, "notification": notification}
		case "3":
			fmt.Println("Enter your error message:")
			errorMessage, _ := reader.ReadString('\n')
			errorMessage = strings.TrimSpace(errorMessage)
			tmpl = parsedTemplates["error"]
			data = map[string]interface{}{"name": name, "errorMessage": errorMessage}
		case "4":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please select a valid option")
			continue
		}
		// render and print the template to the console
		err := tmpl.Execute(os.Stdout, data)
		if err != nil {
			fmt.Println("Error executing template: ")
		}
	}
}
