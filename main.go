package main

import "fmt"

func main() {
	//taskOne := "Learn Go"
	//taskTwo := "Eat Dinner"
	//taskThree := "Read Email"

	taskItems := []string{"Learn Go", "Eat Dinner", "Read Email"}

	//categoryOne := "Private"
	// categoryTwo := "Work"

	fmt.Println("To Do List")

	fmt.Println()

	for _, task := range taskItems {
		fmt.Println(task)
	}

}
