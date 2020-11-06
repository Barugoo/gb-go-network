package main

import (
	"html/template"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func setupServer(c *controller) *gin.Engine {
	router := gin.Default()
	router.SetHTMLTemplate(template.Must(template.New("MyTemplate").ParseFiles("static/tmpl.html")))
	
	router.GET("/lists", c.listTaskLists)
	router.GET("/list", c.getTaskList)
	router.POST("/lists/add", c.createTaskList)
	router.GET("/lists/add", c.createTaskListForm)
	router.PUT("/lists/edit", c.updateTaskList)
	router.GET("/lists/edit", c.updateTaskListForm)
	
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}