package main

import (
	"html/template"
	"os"
)

var page = struct {
	Title string
	List  []string
}{"Список задач", []string{"Сходить за хлебом", "Выполнить пз по курсам на GB", "Не забыть поспать"}}

func main() {
	tmpl := template.Must(template.New("first").Parse(`
	{{define "T"}}
	<html>
		<head>
			<title>{{.Title}}</title>
		</head>
		<body>
			<h1>{{.Title}}</h1>
			<ul>
			{{range .List}}
				{{template "ListItem" . }}
			{{end}}
			</ul>
		</body>
	</html>
	{{end}}

	{{define "ListItem"}}<li>{{.}}<li>{{end}}
  `))
	_ = tmpl.ExecuteTemplate(os.Stdout, "T", page)
}
