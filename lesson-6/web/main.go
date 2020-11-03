package main

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	primitive "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func init() {
	db, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	client = db
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			panic(err)
		}
	}()

	router := gin.Default()
	router.SetHTMLTemplate(template.Must(template.New("MyTemplate").ParseFiles("static/tmpl.html")))

	router.GET("/lists", listTaskLists)
	router.GET("/list", getTaskList)
	router.POST("/lists/add", createTaskList)
	router.GET("/lists/add", createTaskListForm)
	router.POST("/lists/edit", updateTaskList)
	router.GET("/lists/edit", updateTaskListForm)

	router.Run(":8099")
}

func listTaskLists(c *gin.Context) {
	lists, err := GetLists()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.HTML(http.StatusOK, "alllists", lists)
}

func getTaskList(c *gin.Context) {
	list, err := GetList(c.Request.URL.Query().Get("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.HTML(http.StatusOK, "list", list)
}

func createTaskListForm(c *gin.Context) {
	c.HTML(http.StatusOK, "list_create_form", nil)
}

func createTaskList(c *gin.Context) {
	list := TaskList{
		Name:        c.Request.FormValue("name"),
		Description: c.Request.FormValue("description"),
	}

	list, err := CreateList(list)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Redirect(http.StatusFound, fmt.Sprintf("/list?id=%s", list.ID.Hex()))
}

func updateTaskListForm(c *gin.Context) {
	list, err := GetList(c.Request.URL.Query().Get("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.HTML(http.StatusOK, "list_update_form", list)
}

func updateTaskList(c *gin.Context) {
	hex, err := primitive.ObjectIDFromHex(c.Request.URL.Query().Get("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	list := TaskList{
		ID:          hex,
		Name:        c.Request.FormValue("name"),
		Description: c.Request.FormValue("description"),
	}

	list, err = UpdateList(list)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Redirect(http.StatusFound, fmt.Sprintf("/list?id=%s", list.ID.Hex()))
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
