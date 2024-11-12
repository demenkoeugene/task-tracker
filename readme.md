# Task Tracker CLI

The **Task Tracker CLI** is a simple command-line tool for managing tasks. It allows you to add, update, list, and mark the status of tasks easily from the terminal.
https://roadmap.sh/projects/task-tracker

## Table of Contents

- [Features](#features)
- [Requirements](#requirements)
- [Installation](#installation)
- [Usage](#usage)
    - [Adding a Task](#adding-a-task)
    - [Listing Tasks](#listing-tasks)
    - [Updating a Task](#updating-a-task)
    - [Marking Task Status](#marking-task-status)
- [Author](#author)

## Features

- Add new tasks with a description.
- List all tasks or filter by status (`todo`, `in-progress`, `done`).
- Update the description of existing tasks.
- Mark tasks with different statuses (`todo`, `in-progress`, `done`).
- Data is saved persistently in a `tasks.json` file.

## Requirements

- Go 1.18 or higher installed on your system.

## Installation

1. **Clone the repository**:

    ```bash
    git clone https://github.com/yourusername/task-tracker.git
    cd task-tracker
    ```

2. **Initialize the Go module and tidy up dependencies**:

    ```bash
    go mod tidy
    ```

3. **Build the project**:

    ```bash
    go build -o task-cli ./cmd/main.go
    ```

4. **Verify that the `task-cli` binary was created**:

    ```bash
    ls
    ```

## Usage

### Adding a Task

```bash
./task-cli add "Walk the dog"
```
### Listing Tasks
List all tasks:
```bash
./task-cli list
```
List tasks by status:
```bash
./task-cli list todo
./task-cli list done
./task-cli list in-progress
```

### Updating a Task
```bash
./task-cli update 1 "Walk the dog and feed it"
```

### Marking Task Status
Mark a task as in-progress:
```bash
./task-cli mark 1 in-progress
```

Mark a task as done:
```bash
./task-cli mark 1 done
```

Author

Developed by Yevhenii Demenko.
Feel free to reach out on [LinkedIn](https://linkedin.com/in/demenkoeugene)
