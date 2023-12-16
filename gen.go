package main

import (
	"fmt"
	"html/template"
	"os"
)

const structTemplate = `

package main

type {{ .StructName}} struct{

{{range .Fields}} {{.Name}} {{.Type}}
{{end}}

}
`

type Field struct {
	Name string
	Type string
}

type TemplateData struct {
	StructName string
	Fields     []Field
}

func main() {

	data := TemplateData{
		StructName: "Person",
		Fields: []Field{
			{Name: "Name", Type: "string"},
			{Name: "Age", Type: "int32"},
			{Name: "Email", Type: "string"},
		},
	}
	file, err := os.Create("generate_struct.go")
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
	tmpl := template.Must(template.New("structTemplate").Parse(structTemplate))
	err1 := tmpl.Execute(file, data)
	if err1 != nil {
		fmt.Printf("Returned due to the error %q", err)
	}

	fmt.Print("Created a sample struct called structTemplate")
}
