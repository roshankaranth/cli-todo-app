package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/mergestat/timediff"
)

var task_count int = 0

func AddTask(task string) {
	now := time.Now()
	formattedTime := now.Format("2006-01-02T15:04:05")
	task_count += 1

	file, err := os.OpenFile("./tasks/tasks.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file : ", err)
		return
	}

	defer file.Close()

	_, err = file.WriteString("\n" + strconv.Itoa(task_count) + "," + task + "," + formattedTime + "," + "fasle")

	if err != nil {
		fmt.Println("Error writing to file : ", err)
		return
	}

}

func ListTasks() {
	data, err := os.ReadFile("./tasks/tasks.csv")
	check(err)

	r := csv.NewReader(strings.NewReader(string(data)))
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 4, ' ', 0)
	var field bool
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		if field {

			fmt.Fprintf(w, "%s\t%s\t%s\t%s\t\n", record[0], record[1], calculateTime(record[2]), record[3])
		} else {
			fmt.Fprintf(w, "ID\tTask\tCreated\tDone\t\n")
		}

		field = true

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

func calculateTime(taskCreated string) string {
	//2024-07-27T16:45:19-05:00
	tdate, ttime, _ := strings.Cut(taskCreated, "T")
	parts := strings.Split(ttime, "-")
	tunits := strings.Split(parts[0], ":")
	hour, _ := strconv.Atoi(tunits[0])
	min, _ := strconv.Atoi(tunits[1])
	sec, _ := strconv.Atoi(tunits[2])

	dunits := strings.Split(tdate, "-")
	dunits_now := strings.Split(time.Now().Format("2006-01-02"), "-")

	tyear, _ := strconv.Atoi(dunits[0])
	tmonth, _ := strconv.Atoi(dunits[1])
	tday, _ := strconv.Atoi(dunits[2])
	nyear, _ := strconv.Atoi(dunits_now[0])
	nmonth, _ := strconv.Atoi(dunits_now[1])
	nday, _ := strconv.Atoi(dunits_now[2])

	hour += (nyear-tyear)*8760 + (nmonth-tmonth)*730 + (nday-tday)*24

	ftime := timediff.TimeDiff(time.Now().Add(time.Duration(-hour) * time.Hour).Add(time.Duration(-min) * time.Minute).Add(time.Duration(-sec) * time.Second))

	return ftime

}
