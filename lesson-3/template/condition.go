package main

import (
	"html/template"
	"os"
)

var page = struct {
	Title string
	List  []Task
}{"Список задач",
	[]Task{
		Task{true, "Сходить за хлебом"},
		Task{false, "Выполнить пз по курсам на GB"},
		Task{true, "Не забыть поспать"},
	},
}

type Task struct {
	Complete bool
	Text     string
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
			<ul>
			{{range .List}}
				{{template "ListItem" . }}
			{{end}}
			</ul>
		</body>
	</html>
	{{end}}
	 
	{{define "ListItem"}}
		{{if .Complete}}
			<li>(Выполнено) {{.Text}}</li>
		{{else}}
			<li>{{.Text}}</li>
		{{end}}
	{{end}}
	`))

	_ = tmpl.ExecuteTemplate(os.Stdout, "T", page)
}
