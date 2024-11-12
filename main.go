package main

import (
	"fmt"
	"os"
	"strconv"
	"task-tracker/services"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: task-cli <command> [arguments]")
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		handleAdd()
	case "list":
		handleList()
	case "update":
		handleUpdate()
	case "mark":
		handleMark()
	case "help", "--help":
		printHelp()
	default:
		fmt.Println("Unknown command")
	}
}

func handleAdd() {
	if len(os.Args) < 3 {
		fmt.Println("Please provide a task description")
		return
	}
	services.AddTask(os.Args[2])
}

func handleList() {
	status := ""
	if len(os.Args) > 2 {
		status = os.Args[2]
	}
	services.ListTasks(status)
}

func handleUpdate() {
	if len(os.Args) < 4 {
		fmt.Println("Please provide task ID and new description")
		return
	}
	id := atoi(os.Args[2])
	services.UpdateTask(id, os.Args[3])
}

func handleMark() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: task-cli mark <taskID> <status>")
		fmt.Println("Please provide task ID and status")
		return
	}
	id := atoi(os.Args[2])
	status := os.Args[3]
	services.MarkTask(id, status)
}

func printHelp() {
	fmt.Println(`
Task Tracker CLI
Usage:
  ./task-cli <command> [arguments]

Available Commands:
  add <description>        Add a new task
  list [status]            List all tasks (optional status: todo, in-progress, done)
  update <taskID> <desc>   Update task description by ID
  mark <taskID> <status>   Mark task status (valid statuses: todo, in-progress, done)
  help, --help             Show this help message

Examples:
  ./task-cli add "Buy groceries"
  ./task-cli list
  ./task-cli list done
  ./task-cli update 1 "Buy groceries and cook dinner"
  ./task-cli mark 1 done
  ./task-cli mark 1 in-progress
`)
}

func atoi(str string) int {
	result, _ := strconv.Atoi(str)
	return result
}
