package task

import (
	"errors"
	"log"
	"os"
	"path/filepath"

	"github.com/fatih/color"
)

var FILE string

func init() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	FILE = filepath.Join(cwd, "tasks.json")
}

func AddTask(name string) {
	tasks, _ := LoadFileContent[[]Task](FILE)

	task := Task{
		Id:     uint(len(tasks) + 1),
		Name:   name,
		Status: "PENDING",
	}

	tasks = append(tasks, task)

	err2 := WriteFile(tasks, FILE)
	color.New(color.FgGreen).Println("Task Created!")

	if err2 != nil {
		log.Fatal(err2)
		panic(err2)
	}
}

func ListTasks() {
	tasks, err := LoadFileContent[[]Task](FILE)
	red := color.New(color.FgRed)
	green := color.New(color.FgGreen)
	blue := color.New(color.FgBlue)

	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	color.New(color.BgHiWhite).Printf("%-4s %-20s %-10s %-10s\n", "ID", "Name", "Completed", "Canceled")

	for _, task := range tasks {
		completed := "[ ]"
		canceled := "[ ]"
		color := blue

		if isCompleted(task) {
			completed = "[x]"
			color = green
		}
		if isCanceled(task) {
			canceled = "[x]"
			color = red
		}

		color.Printf("%-4d %-20s %-10s %-10s\n", task.Id, task.Name, completed, canceled)
	}

}

func CompleteTask(taskId uint) error {
	return editStatus(taskId, "COMPLETE")
}

func CancelTask(taskId uint) error {
	return editStatus(taskId, "CANCELED")
}

func editStatus(taskId uint, status string) error {
	tasks, err := LoadFileContent[[]Task](FILE)
	if err != nil {
		return err
	}

	found := false

	for idx := range tasks {
		if tasks[idx].Id == taskId {
			tasks[idx].Status = status
			found = true

			color.New(color.FgHiYellow).Printf("Changing task status %d -  %s \n", taskId, tasks[idx].Status)
			break
		}

	}

	if !found {
		return errors.New("Task not found")
	}

	if err := WriteFile(tasks, FILE); err != nil {
		return err
	}

	return nil

}

func isCanceled(task Task) bool {
	return task.Status == "CANCELED"
}
func isCompleted(task Task) bool {
	return task.Status == "COMPLETE"
}
