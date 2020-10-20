package main

import (
	"fmt"
	"html/template"
	"os"
)

var page = struct {
	Title string
	List  []Task
	Perm  Permission
}{"Список задач",
	[]Task{
		Task{true, "Сходить за хлебом"},
		Task{false, "Выполнить пз по курсам на GB"},
		Task{true, "Не забыть поспать"},
	},
	Permission{true},
}

type Task struct {
	Complete bool
	Text     string
}

type Permission struct {
	admin bool
}

func (p Permission) AdminNeeded(status string) bool {
	if status == "admin" {
		return p.admin
	}
	return true
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
				{{if .Perm.AdminNeeded "user"}}
					<h3>А для этой строки права не нужны.</h3>
				{{end}}
				{{if .Perm.AdminNeeded "admin"}}
					<h3>Ты админ. Поздравляю</h3>
				{{end}}
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

	fmt.Println(tmpl.ExecuteTemplate(os.Stdout, "T", page))
}
