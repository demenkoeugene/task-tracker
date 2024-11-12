package utils

import (
	"encoding/json"
	"os"
	"task-tracker/models"
)

const tasksFile = "results/tasks.json"

func LoadTasks() ([]models.Task, error) {
	file, err := os.Open(tasksFile)
	if err != nil {
		if os.IsNotExist(err) {
			return []models.Task{}, nil
		}
		return nil, err
	}
	defer file.Close()

	var tasks []models.Task
	err = json.NewDecoder(file).Decode(&tasks)
	return tasks, err
}

func SaveTasks(tasks []models.Task) error {
	file, err := os.Create(tasksFile)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(tasks)
}
