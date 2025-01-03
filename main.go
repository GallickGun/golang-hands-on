package main

import "fmt"

var (
	taskOne   = "Learn Go"
	taskTwo   = "Eat Dinner"
	taskThree = "Read Email"
)

var taskItems = []string{taskOne, taskTwo, taskThree}

func main() {

	//categoryOne := "Private"
	// categoryTwo := "Work"

	fmt.Println("To Do List")

	fmt.Println()
	printTasks()
}

func printTasks() {
	for _, task := range taskItems {
		fmt.Println(task)
	}
}
