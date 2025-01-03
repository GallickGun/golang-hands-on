package main

import "fmt"

func main() {

	var (
		taskOne   = "Learn Go"
		taskTwo   = "Eat Dinner"
		taskThree = "Read Email"
	)

	var taskItems = []string{taskOne, taskTwo, taskThree}

	fmt.Println("To Do List")

	fmt.Println()
	printTasks(taskItems)
	taskItems = addTasks(taskItems, "Sleep")
}

func printTasks(taskItems []string) {
	for _, task := range taskItems {
		fmt.Println(task)
	}
}

func addTasks(taskItems []string, newTask string) []string {
	var updatedTasks = append(taskItems, newTask)
	return updatedTasks
}
