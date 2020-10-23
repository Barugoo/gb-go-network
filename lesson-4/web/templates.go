package main

// TaskList - список задач
type TaskList struct {
	ID          int
	Name        string
	Description string
	List        []Task
}

// Task - задача и ее статус
type Task struct {
	ID       int
	ListID   int
	Text     string
	Complete bool
}
