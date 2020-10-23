package main

import (
	"fmt"
	"log"
)

// GetAllLists — получение всех списков с задачами
func GetAllLists() ([]TaskList, error) {
	res := []TaskList{}

	rows, err := database.Query("select * from task_list_app.lists")
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		list := TaskList{}

		err := rows.Scan(&list.ID, &list.Name, &list.Description)
		if err != nil {
			log.Println(err)
			continue
		}

		res = append(res, list)
	}

	return res, nil
}

// GetList — получение списка по id
func GetList(id string) (TaskList, error) {
	list := TaskList{}

	row := database.QueryRow(fmt.Sprintf("select * from task_list_app.lists where lists.id = %s", id))
	err := row.Scan(&list.ID, &list.Name, &list.Description)
	if err != nil {
		return list, err
	}

	rows, err := database.Query(fmt.Sprintf("select * from task_list_app.tasks WHERE tasks.list_id = %v", id))
	if err != nil {
		return list, err
	}
	defer rows.Close()

	for rows.Next() {
		task := Task{}

		err := rows.Scan(&task.ID, new(int), &task.Text, &task.Complete)
		if err != nil {
			log.Println(err)
			continue
		}

		list.List = append(list.List, task)
	}

	return list, nil
}
