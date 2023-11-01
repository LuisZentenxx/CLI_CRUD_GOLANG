package task

import (
	"encoding/json"
	"fmt"
	"os"
)

type Task struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Complete bool   `json:"complete"`
}

func ListTasks(tasks []Task) {
	if len(tasks) == 0 {
		fmt.Println("No tasks to list")
		return
	}

	for _, task := range tasks {

		status := "✓"
		if task.Complete {
			status = ""
		}

		fmt.Printf("[%s] %d %s\n", status, task.ID, task.Name)
	}

}

func AddTask(tasks []Task, name string) []Task {
	newTask := Task{
		ID:       10,
		Name:     name,
		Complete: false,
	}
	return append(tasks, newTask)
}

func SaveTask(file *os.File, tasks []Task) {
	bytes, err := json.Marshal(tasks)
	if err != nil {
		panic(err)
	}

	_, err := file.Seek(0, 0)
	if err != nil {
		panic(err)
	}

	err = file.Truncate(0)
	if err != nil {
		panic(err)
	}

}
