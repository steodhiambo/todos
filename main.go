package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type Todo struct {
	ID       int    `json:"id"`
	Task     string `json:"task"`
	Complete bool   `json:"complete"`
}

var todos []Todo

func main() {
	loadTodos()

	if len(os.Args) < 2 {
		fmt.Println("Usage: add <task>, list, done <id>, delete <id>")
		return
	}

	switch os.Args[1] {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Task description is required.")
			return
		}
		addTask(os.Args[2])
	case "list":
		listTasks()
	case "done":
		if len(os.Args) < 3 {
			fmt.Println("Task ID is required.")
			return
		}
		markDone(os.Args[2])
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Task ID is required.")
			return
		}
		deleteTask(os.Args[2])
	default:
		fmt.Println("Unknown command.")
	}
	saveTodos()
}

func loadTodos() {
	file, err := os.ReadFile("todos.json")
	if err == nil {
		_ = json.Unmarshal(file, &todos)
	}
}

func saveTodos() {
	data, _ := json.MarshalIndent(todos, "", "  ")
	_ = os.WriteFile("todos.json", data, 0644)
}

func addTask(task string) {
	id := 1
	if len(todos) > 0 {
		id = todos[len(todos)-1].ID + 1
	}
	todos = append(todos, Todo{ID: id, Task: task, Complete: false})
	fmt.Println("Task added:", task)
}

func listTasks() {
	if len(todos) == 0 {
		fmt.Println("No tasks available.")
		return
	}
	for _, todo := range todos {
		status := " "
		if todo.Complete {
			status = "x"
		}
		fmt.Printf("[%s] %d: %s\n", status, todo.ID, todo.Task)
	}
}

func markDone(idStr string) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Invalid task ID.")
		return
	}
	for i, todo := range todos {
		if todo.ID == id {
			todos[i].Complete = true
			fmt.Println("Marked task as done:", todo.Task)
			return
		}
	}
	fmt.Println("Task not found.")
}

func deleteTask(idStr string) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Invalid task ID.")
		return
	}
	for i, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			fmt.Println("Deleted task:", todo.Task)
			return
		}
	}
	fmt.Println("Task not found.")
}
