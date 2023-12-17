package main

import (
	"fmt"
	"html/template"
	"os"
)

//go:generate go run fibonacci_gen.go
const fibonacciTemplate = `
package main
import (
	"fmt"
)

func {{.FunctionName}}({{.Variable}} int) {{.Type}} {
	if {{.Variable}} {{.Comparison}} 1 {
		return {{.Variable}}
	} else {
		return {{.FunctionName}}({{.Variable}}-1) + {{.FunctionName}}({{.Variable}}-2)
	}
}

func main() {
	fmt.Println({{.FunctionName}}(6))
}
`

func main() {
	type TemplateData struct {
		FunctionName string
		Variable     string
		Type         string
		Comparison   template.HTML
	}

	data := TemplateData{
		FunctionName: "Fibonacci",
		Variable:     "n",
		Type:         "int",
		Comparison:   template.HTML("<="), // Prevents HTML escaping
	}

	file, err := os.Create("generate_fibonacci.go")
	if err != nil {
		panic(err)
	}
	defer func(body *os.File) {
		errorclosing := body.Close()
		if errorclosing != nil {
			panic(errorclosing)
		}
	}(file)

	tmpl, err := template.New("fibonacci").Parse(fibonacciTemplate)
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return
	}

	err = tmpl.Execute(file, data)
	if err != nil {
		fmt.Println("Error executing template:", err)
		return
	}
	fmt.Println("Generated a file that calculates fibonacci")
}

//if {{.Variable}} <=1  this while generating will give weird output like  if a &lt;= 1

//template.HTML("<=") Needed to use this
