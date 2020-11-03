package main

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TaskList - список задач
type TaskList struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	List        []Task             `json:"list"`
}

// Task - задача и ее статус
type Task struct {
	ID        int    `json:"id"`
	ListID    int    `json:"list_id"`
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}
