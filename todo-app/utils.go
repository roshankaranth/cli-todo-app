package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
	"text/tabwriter"
)

func AddTask() {
	fmt.Printf("Adding task")
}

func ListTasks() {
	data, err := os.ReadFile("./tasks/tasks.csv")
	check(err)

	r := csv.NewReader(strings.NewReader(string(data)))
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 4, ' ', 0)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t\n", record[0], record[1], record[2], record[3])

	}
	w.Flush()
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
