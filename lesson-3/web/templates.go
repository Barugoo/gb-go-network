package main

// TaskList - список задач
type TaskList struct {
	Name        string
	Description string
	List        []Task
}

// Task - задача и ее статус
type Task struct {
	ID       string
	Text     string
	Complete bool
}
