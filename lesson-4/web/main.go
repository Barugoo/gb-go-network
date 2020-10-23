package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var tmpl = template.Must(template.New("MyTemplate").ParseFiles("static/tmpl.html"))

var database *sql.DB

func main() {
	db, err := sql.Open("mysql", "root:my-secret-pw@/task_list_app")
	if err != nil {
		log.Fatal(err)
	}
	database = db

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	router := http.NewServeMux()

	router.HandleFunc("/", viewLists)
	router.HandleFunc("/list", viewList)

	log.Fatal(http.ListenAndServe(":8099", router))
}

func viewLists(w http.ResponseWriter, r *http.Request) {
	lists, err := GetAllLists()
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	if err := tmpl.ExecuteTemplate(w, "alllists", lists); err != nil {
		log.Println(err)
	}
}

func viewList(w http.ResponseWriter, r *http.Request) {
	list, err := GetList(r.URL.Query().Get("id"))
	if err != nil {
		log.Println(err)
		w.WriteHeader(404)
		return
	}

	if err := tmpl.ExecuteTemplate(w, "list", list); err != nil {
		log.Println(err)
	}
}
