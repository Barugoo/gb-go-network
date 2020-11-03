package main

import (
	"context"

	bson "go.mongodb.org/mongo-driver/bson"
	primitive "go.mongodb.org/mongo-driver/bson/primitive"
)

const dbName = "db"

func GetLists() ([]TaskList, error) {
	var res []TaskList

	collection := client.Database(dbName).Collection("lists")

	cur, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}

	err = cur.All(context.TODO(), &res)
	return res, err
}

// GetList — получение списка по id
func GetList(id string) (TaskList, error) {
	collection := client.Database(dbName).Collection("lists")

	var obj TaskList
	hex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return TaskList{}, err
	}
	res := collection.FindOne(context.Background(), bson.M{"_id": hex})
	err = res.Decode(&obj)
	return obj, err
}

// CreateList — создание листа задач
func CreateList(list TaskList) (TaskList, error) {
	list.ID = primitive.NewObjectID()

	collection := client.Database(dbName).Collection("lists")
	_, err := collection.InsertOne(context.Background(), &list)
	return list, err
}

func UpdateList(list TaskList) (TaskList, error) {
	c := client.Database(dbName).Collection("lists")
	_, err := c.UpdateOne(context.Background(),
		bson.M{"_id": list.ID},
		bson.D{
			bson.E{"$set", bson.D{
				bson.E{"name", list.Name},
				bson.E{"description", list.Description},
			},
			},
		},
	)
	return list, err
}

// // DeleteTask — получение списка по id
// func DeleteTask(taskID string) error {
// 	_, err := database.Exec("DELETE FROM `task_list_app`.`tasks` WHERE (`id` = ?)", taskID)
// 	return err
// }

// // UpdateTask - обновление задачи
// func UpdateTask(id int, name *string, completed *bool) (Task, error) {
// 	if name != nil {
// 		_, err := database.Exec("UPDATE `task_list_app`.`tasks` SET `name` = ? WHERE (`id` = ?)", *name, id)
// 		if err != nil {
// 			return Task{}, nil
// 		}
// 	}

// 	if completed != nil {
// 		_, err := database.Exec("UPDATE `task_list_app`.`tasks` SET `completed` = ? WHERE (`id` = ?)", *completed, id)
// 		if err != nil {
// 			return Task{}, nil
// 		}
// 	}
// 	task := Task{}

// 	row := database.QueryRow(fmt.Sprintf("select * from task_list_app.tasks where task.id = `%s`", id))
// 	err := row.Scan(&task.ID, &task.ListID, &task.Name, &task.Completed)
// 	return task, err
// }

// // CreateTask — создание задачи
// func CreateTask(task Task) (Task, error) {
// 	res, err := database.Exec("insert into task_list_app.tasks (list_id, name, completed) values (?, ?, ?)", task.ListID, task.Name, task.Completed)
// 	if err != nil {
// 		return Task{}, err
// 	}
// 	taskID, err := res.LastInsertId()
// 	if err != nil {
// 		return Task{}, err
// 	}
// 	task.ID = int(taskID)
// 	return task, nil
// }

// GetAllLists — получение всех списков с задачами
// func GetAllLists() ([]TaskList, error) {
// 	res := []TaskList{}

// 	rows, err := database.Query("select * from task_list_app.lists")
// 	if err != nil {
// 		return res, err
// 	}
// 	defer rows.Close()

// 	for rows.Next() {
// 		list := TaskList{}

// 		err := rows.Scan(&list.ID, &list.Name, &list.Description)
// 		if err != nil {
// 			log.Println(err)
// 			continue
// 		}

// 		res = append(res, list)
// 	}

// 	return res, nil
// }

// // GetTask — получение задачи по id
// func GetTask(taskID int) (Task, error) {
// 	task := Task{}

// 	row := database.QueryRow(fmt.Sprintf("select * from task_list_app.tasks where task.id = `%d`", taskID))
// 	err := row.Scan(&task.ID, &task.ListID, &task.Name, &task.Completed)
// 	return task, err
// }
