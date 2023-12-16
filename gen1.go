package main

import (
	"fmt"
	"html/template"
	"os"
)

//go:generate go run gen1.go
const newMainfile = `
package main

import "fmt"

func {{.FuncName}}() {
    fmt.Println("Hello, World")
}
`

type TemplateData1 struct {
	FuncName string
}

func main() {
	data := TemplateData1{
		FuncName: "MainFunction",
	}

	file, err := os.Create("hello_world.go")
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	tmpl := template.Must(template.New("Main").Parse(newMainfile))
	err = tmpl.Execute(file, data)
	if err != nil {
		panic(err)
	}

	fmt.Println("Created a new hello world file")
}
