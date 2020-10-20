package main

import (
	"html/template"
	"os"
)

var page = struct {
	Title   string
	Content string
}{
	Title:   "Шаблонизированная страница",
	Content: "Ее содержимое",
}

func main() {
	tmpl := template.Must(template.New("first").Parse(`
		{{define "T"}}
		<html>
			<head>
				<title>{{.Title}}</title>
			</head>
			<body>
				<h1>{{.Title}}</h1>
				<p>{{.Content}}</p>
			</body>
		</html>
		{{end}}
	`))

	_ = tmpl.ExecuteTemplate(os.Stdout, "T", page)
}
