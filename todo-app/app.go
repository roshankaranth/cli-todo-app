package main

import "os"

func main() {

	if len(os.Args) == 1 {
		HelpScreen()
		os.Exit(0)
	}
	command := os.Args[1]

	switch command {
	case "add":
		AddTask(os.Args[2])
	case "list":
		ListTasks()
	case "help", "":
		HelpScreen()
	case "delete":
		DeleteTask()
	case "completed":
		CompleteTask()
	default:
		HelpScreen()
	}
}
