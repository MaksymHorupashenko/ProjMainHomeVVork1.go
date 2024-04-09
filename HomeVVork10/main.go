package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Task struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}

var tasks []Task
var currentID int

func getTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func getTaskByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	parts := strings.Split(r.URL.Path, "/")
	taskID, err := strconv.Atoi(parts[len(parts)-1])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	for _, task := range tasks {
		if task.ID == taskID {
			json.NewEncoder(w).Encode(task)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

func addTask(w http.ResponseWriter, r *http.Request) {
	var newTask Task
	json.NewDecoder(r.Body).Decode(&newTask)
	newTask.ID = currentID
	currentID++
	tasks = append(tasks, newTask)
	w.WriteHeader(http.StatusCreated)
}

func updateTask(w http.ResponseWriter, r *http.Request) {
	var updatedTask Task
	json.NewDecoder(r.Body).Decode(&updatedTask)
	for i, task := range tasks {
		if task.ID == updatedTask.ID {
			tasks[i] = updatedTask
			break
		}
	}
	w.WriteHeader(http.StatusNoContent)
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	var taskToDelete Task
	json.NewDecoder(r.Body).Decode(&taskToDelete)
	for i, task := range tasks {
		if task.ID == taskToDelete.ID {
			tasks = append(tasks[:i], tasks[i+1:]...)
			break
		}
	}
	w.WriteHeader(http.StatusNoContent)
}

func markTaskDone(w http.ResponseWriter, r *http.Request) {
	var taskToUpdate Task
	json.NewDecoder(r.Body).Decode(&taskToUpdate)
	for i, task := range tasks {
		if task.ID == taskToUpdate.ID {
			tasks[i].Completed = true
			break
		}
	}
	w.WriteHeader(http.StatusNoContent)
}

func markTaskIsnotdone(w http.ResponseWriter, r *http.Request) {
	var taskToUpdate Task
	json.NewDecoder(r.Body).Decode(&taskToUpdate)
	for i, task := range tasks {
		if task.ID == taskToUpdate.ID {
			tasks[i].Completed = false
			break
		}
	}
	w.WriteHeader(http.StatusNoContent)
}

func main() {
	tasks = []Task{}
	currentID = 1
	
	http.HandleFunc("/tasks", getTasks)
	http.HandleFunc("/tasks/add", addTask)
	http.HandleFunc("/tasks/update", updateTask)
	http.HandleFunc("/tasks/delete", deleteTask)
	http.HandleFunc("/tasks/completed", markTaskDone)
	http.HandleFunc("/tasks/incomplete", markTaskIsnotdone)
	http.HandleFunc("/tasks/", getTaskByID)

	// Start server
	fmt.Println("Server starts 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
