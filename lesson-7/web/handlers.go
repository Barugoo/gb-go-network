package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	primitive "go.mongodb.org/mongo-driver/bson/primitive"
)

type controller struct {
	db DB
}

// @Summary PageTaskLists
// @Produce  text/html
// @Router /lists [get]
func (cnt *controller) listTaskLists(c *gin.Context) {
	lists, err := cnt.db.GetLists()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.HTML(http.StatusOK, "alllists", lists)
}

// @Summary PageSingleTaskList
// @Description returns single task list page
// @Produce  text/html
// @Param id query string false "task list id"
// @Router /list [get]
func (cnt *controller) getTaskList(c *gin.Context) {
	list, err := cnt.db.GetList(c.Request.URL.Query().Get("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.HTML(http.StatusOK, "list", list)
}

// @Summary FormCreateTaskList
// @Description returns create task list form
// @Produce  text/html
// @Router /lists/add [get]
func (cnt *controller) createTaskListForm(c *gin.Context) {
	c.HTML(http.StatusOK, "list_create_form", nil)
}

// @Summary CreateTaskList
// @Description creates new task list and redirects to /list?id={created tasklist ID}
// @Produce  text/html
// @Router /lists/add [post]
func (cnt *controller) createTaskList(c *gin.Context) {
	list := TaskList{
		Name:        c.Request.FormValue("name"),
		Description: c.Request.FormValue("description"),
	}

	list, err := cnt.db.CreateList(list)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Redirect(http.StatusFound, fmt.Sprintf("/list?id=%s", list.ID.Hex()))
}

// @Summary FormUpdateTaskList
// @Description returns update task list form
// @Produce  text/html
// @Router /lists/edit [get]
func (cnt *controller) updateTaskListForm(c *gin.Context) {
	list, err := cnt.db.GetList(c.Request.URL.Query().Get("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.HTML(http.StatusOK, "list_update_form", list)
}

// @Summary UpdateTaskList
// @Description updates task list and redirects to /list?id={created tasklist ID}
// @Produce  text/html
// @Router /lists/edit [put]
func (cnt *controller) updateTaskList(c *gin.Context) {
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

	list, err = cnt.db.UpdateList(list)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Redirect(http.StatusFound, fmt.Sprintf("/list?id=%s", list.ID.Hex()))
}
