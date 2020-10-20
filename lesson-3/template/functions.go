package main

import (
	"fmt"
	"html/template"
	"os"
	"time"
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
	tmpl := template.Must(template.New("first").Funcs(template.FuncMap{
		"year": func(arr ...string) string {
			var result string
			for _, elem := range arr {
				result += elem
			}
			return fmt.Sprintf("%s %d", result, time.Now().Unix())
		},
	}).Parse(`
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
     {{year "dsds" "sdsds" "3232"}}
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
