package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mynk1997/golang-react-todo/models"
	"github.com/mynk1997/golang-react-todo/router"
)

var task []models.ToDoList

func main() {
	r := router.Router()
	fmt.Println("Starting the server on port 9000....")
	log.Fatal(http.ListenAndServe(":9000", r))

	task = append(task, models.ToDoList{Task: "ReactJS", Status: false})

}
