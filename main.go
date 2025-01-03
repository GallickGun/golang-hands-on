package main

import (
	"fmt"
	"net/http"
)

var (
	taskOne   = "Learn Go"
	taskTwo   = "Eat Dinner"
	taskThree = "Read Email"
)

var taskItems = []string{taskOne, taskTwo, taskThree}

func main() {

	http.HandleFunc("/hello", helloUser)
	http.HandleFunc("/show-tasks", showTasks)
	http.ListenAndServe(":8080", nil)
}

func helloUser(writer http.ResponseWriter, request *http.Request) {
	var greeting = "Hello, User! Welcome to the To Do List App!"
	fmt.Fprintln(writer, greeting)
}

func showTasks(writer http.ResponseWriter, request *http.Request) {
	for _, task := range taskItems {
		fmt.Fprintln(writer, task)
	}
}
