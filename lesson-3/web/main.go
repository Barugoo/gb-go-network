package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type listPostPage struct {
	Title string
	Posts []Post
}

type Post struct {
	ID        int32
	Title     string
	Text      string
	Author    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

var posts = map[int]*Post{
	1: {
		ID:        1,
		Title:     "some title 1",
		Text:      "some text 1 ",
		Author:    "some author 1",
		CreatedAt: time.Now().Add(-time.Hour),
	},
	2: {
		ID:        2,
		Title:     "some title 2",
		Text:      "some text 2",
		Author:    "some author 2",
		CreatedAt: time.Now().Add(-time.Hour * 2),
	},
	3: {
		ID:        3,
		Title:     "some title 3",
		Text:      "some text 3",
		Author:    "some author 3",
		CreatedAt: time.Now().Add(-time.Hour * 3),
	},
}

func main() {
	router := mux.NewRouter()

	// шаблон со списком всех постов (коротким без отображения поля text)
	// router.HandleFunc("/", listPostsHandler).Methods("GET")

	// шаблон с текстовыми полями для задания Title, Text и Author
	// router.HandleFunc("/", createPostHandler).Methods("POST")

	// шаблон со страницей одного поста (полного, тобишь с отображением)
	router.HandleFunc("/{id}", getPostHandler).Methods("GET")

	// шаблон с текстовыми полями для обновления Title, Text и Author
	// router.HandleFunc("/{id}", updatePostHandler).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8099", router))
}

func getPostHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postIDRaw, ok := vars["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	postID, err := strconv.Atoi(postIDRaw)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	post, ok := posts[postID]
	if !ok {
		w.Write([]byte("Post not found"))
		w.WriteHeader(http.StatusNotFound)
		return
	}

	tmpl := template.Must(template.New("first").Parse(`
	{{define "T"}}
	<html>
		<head>
			<title>{{.Title}}</title>
		</head>
		<body>
			<h1>{{.Title}}</h1>
			<p>{{.Text}}</p>
		</body>
	</html>
	{{end}}
	`))

	if err := tmpl.ExecuteTemplate(w, "T", post); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
	}
}
