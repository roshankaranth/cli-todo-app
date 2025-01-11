package main

import (
	"fmt"
	"os"
)

func AddTask() {
	fmt.Printf("Adding task")
}

func ListTasks() {
	data, err := os.ReadFile("./tasks/tasks.csv")
	check(err)
	fmt.Printf(string(data))
}

func DeleteTask() {
	fmt.Printf("Deleting tasks")
}

func CompleteTask() {
	fmt.Printf("Complete tasks")
}

func HelpScreen() {
	fmt.Printf("Display Help screen")
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func printTask() {

}
