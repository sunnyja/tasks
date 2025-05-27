package main

import (
	"context"
	"log"
	"os"

	"tasks/pkg/storage"
	"tasks/pkg/repo"
)

func main() {
	pwd := os.Getenv("postgres")
    conn := "postgres://postgres:"+pwd+"@postgres/tasks?sslmode=disable"
	database, err := storage.New(conn)

	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	defer database.Close()
	context := context.Background()
	taskStorage := task.NewTaskStorage(database.Db)

	// Пример работы
	var newTask = task.Task{
		Title: "изучить Go", 
		Description: "нужно больше практиковаться",
		Status: 1, //задача создана
		AuthorId: 1,
		AssignedId: 2,
	}
	taskId, err := taskStorage.NewTask(context, newTask)
	if err != nil {
		log.Printf("Failed to create task: %v", err)
	} else {
		log.Printf("Create task with ID: %d", taskId)
	}
}