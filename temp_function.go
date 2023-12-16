package main

import (
	"os"
	"strings"
	"text/template"
)

// A custom function to convert a string to uppercase
func toUpper(s string) string {
	return strings.ToUpper(s)
}

// Another custom function to repeat a string n times
func repeat(s string, n int) string {
	return strings.Repeat(s, n)
}

func main() {
	// Create a new template
	tmpl := template.New("example")

	// Define the functions for use in the template
	tmpl.Funcs(template.FuncMap{
		"ToUpper": toUpper,
		"Repeat":  repeat,
	})

	// Define the template content
	const tpl = `Uppercase: {{ ToUpper "hello" }}
Repeat: {{ Repeat "Go! " 3 }}`

	// Parse the template
	tmpl, err := tmpl.Parse(tpl)
	if err != nil {
		panic(err)
	}

	// Execute the template
	err = tmpl.Execute(os.Stdout, nil)
	if err != nil {
		panic(err)
	}
}
