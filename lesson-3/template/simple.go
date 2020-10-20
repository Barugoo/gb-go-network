package main

import (
	"os"
	"text/template"
)

func main() {
	tmpl := template.Must(template.New("first").Parse(`
	{{define "T"}}
	<html>
		<body><h1>{{.}}</h1><p><b>{{.}}</b></p></body>
	</html>
	{{end}}
	`))

	tmpl.ExecuteTemplate(os.Stdout, "T", "Hello, world!")
}
