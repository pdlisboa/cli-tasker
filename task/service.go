package task

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
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

	if err2 != nil {
		log.Fatal(err2)
		panic(err2)
	}
}

func ListTasks() []Task {
	tasks, err := LoadFileContent[[]Task](FILE)

	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	return tasks
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

			fmt.Printf("Changing task status %d -  %+v \n", taskId, tasks[idx])
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
