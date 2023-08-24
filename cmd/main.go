package main

import (
	"fmt"
	"log"
	database "tasklist/database"
	task "tasklist/task"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	db, err := database.Connect()
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	timeExample := time.Date(
		2023, 8, 24, 16, 50, 01, 651387237, time.UTC)
	fmt.Println(timeExample)

	myList := task.NewToDoList(db)
	myList.AddTask("finish task list", time.Now().Add(time.Hour*24*7))
	list, err := myList.ListTasks()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(list)

}
