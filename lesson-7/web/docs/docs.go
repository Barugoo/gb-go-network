// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Dmitry Shelamov",
            "email": "barugoo@yandex.ru"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/list": {
            "get": {
                "description": "returns single task list page",
                "produces": [
                    "text/html"
                ],
                "summary": "PageSingleTaskList",
                "parameters": [
                    {
                        "type": "string",
                        "description": "task list id",
                        "name": "id",
                        "in": "query"
                    }
                ]
            }
        },
        "/lists": {
            "get": {
                "produces": [
                    "text/html"
                ],
                "summary": "PageTaskLists"
            }
        },
        "/lists/add": {
            "get": {
                "description": "returns create task list form",
                "produces": [
                    "text/html"
                ],
                "summary": "FormCreateTaskList"
            },
            "post": {
                "description": "creates new task list and redirects to /list?id={created tasklist ID}",
                "produces": [
                    "text/html"
                ],
                "summary": "CreateTaskList"
            }
        },
        "/lists/edit": {
            "get": {
                "description": "returns update task list form",
                "produces": [
                    "text/html"
                ],
                "summary": "FormUpdateTaskList"
            },
            "put": {
                "description": "updates task list and redirects to /list?id={created tasklist ID}",
                "produces": [
                    "text/html"
                ],
                "summary": "UpdateTaskList"
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "localhost",
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "Task List",
	Description: "This is Task List",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
