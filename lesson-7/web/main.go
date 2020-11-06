package main

import (
	"context"
	"log"

	// gin-swagger middleware
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	_ "github.com/Barugoo/gb-go-network/lesson-6/web/docs"
)

func init() {

}

// @title Task List
// @version 1.0
// @description This is Task List

// @contact.name Dmitry Shelamov
// @contact.email barugoo@yandex.ru

// @host localhost
// @BasePath /
func main() {
	db, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := db.Disconnect(context.Background()); err != nil {
			panic(err)
		}
	}()

	err = db.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	mongodb := &dbMongo{
		client: db,
	}

	c := &controller{
		db: mongodb,
	}

	srv := setupServer(c)
	srv.Run(":8099")
}

// type UpdateTaskRequest struct {
// 	Complete *bool   `json:"complete"`
// 	Name     *string `json:"name"`
// }

// func createTask(w http.ResponseWriter, r *http.Request) {
// 	var task Task
// 	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
// 		log.Println(err)
// 		w.WriteHeader(400)
// 		return
// 	}

// 	task, err := CreateTask(task)
// 	if err != nil {
// 		log.Println(err)
// 		w.WriteHeader(500)
// 		return
// 	}

// 	if err := tmpl.ExecuteTemplate(w, "task", task); err != nil {
// 		log.Println(err)
// 	}
// }

// func createTaskList(w http.ResponseWriter, r *http.Request) {
// 	var list TaskList
// 	if err := json.NewDecoder(r.Body).Decode(&list); err != nil {
// 		log.Println(err)
// 		w.WriteHeader(400)
// 		return
// 	}

// 	list, err := CreateList(list)
// 	if err != nil {
// 		log.Println(err)
// 		w.WriteHeader(500)
// 		return
// 	}

// 	if err := tmpl.ExecuteTemplate(w, "list", list); err != nil {
// 		log.Println(err)
// 	}
// }
