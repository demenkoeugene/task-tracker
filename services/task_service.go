package services

import (
	"fmt"
	"task-tracker/models"
	"task-tracker/utils"
	"time"
)

func AddTask(description string) {
	tasks, err := utils.LoadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	newTask := models.Task{
		Id:          len(tasks) + 1,
		Description: description,
		Status:      models.StatusTodo,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	tasks = append(tasks, newTask)

	if err := utils.SaveTasks(tasks); err != nil {
		fmt.Println("Error saving tasks:", err)
		return
	}

	fmt.Println("Task added successfully:")
	newTask.Print()
}

func ListTasks(status string) {
	tasks, err := utils.LoadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	for _, task := range tasks {
		if status == "" || string(task.Status) == status {
			task.Print()
		}
	}
}

func UpdateTask(id int, newDescription string) {
	tasks, err := utils.LoadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	for i, task := range tasks {
		if task.Id == id {
			tasks[i].UpdateDescription(newDescription)
			utils.SaveTasks(tasks)
			fmt.Println("Task updated successfully")
			return
		}
	}
	fmt.Println("Task not found")
}

func MarkTask(id int, status string) {
	if !isValidStatus(models.TaskStatus(status)) {
		fmt.Println("Invalid status!")
		return
	}

	tasks, err := utils.LoadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	for i, task := range tasks {
		if task.Id == id {
			tasks[i].UpdateStatus(models.TaskStatus(status))
			utils.SaveTasks(tasks)
			fmt.Printf("Task marked as %s\n", status)
			return
		}
	}
	fmt.Println("Task not found")
}

func isValidStatus(status models.TaskStatus) bool {
	return status == models.StatusTodo || status == models.StatusInProgress || status == models.StatusDone
}
