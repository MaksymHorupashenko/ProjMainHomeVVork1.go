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
	parts := strings.Split(r.URL.Path, "/")
	taskID, err := strconv.Atoi(parts[len(parts)-1])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	for i, task := range tasks {
		if task.ID == taskID {
			tasks = append(tasks[:i], tasks[i+1:]...)
			break
		}
	}
	w.WriteHeader(http.StatusNoContent)
}

func markTaskDone(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	taskID, err := strconv.Atoi(parts[len(parts)-1])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	for i, task := range tasks {
		if task.ID == taskID {
			tasks[i].Completed = true
			break
		}
	}
	w.WriteHeader(http.StatusNoContent)
}

func markTaskIsnotdone(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	taskID, err := strconv.Atoi(parts[len(parts)-1])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	for i, task := range tasks {
		if task.ID == taskID {
			tasks[i].Completed = false
			break
		}
	}
	w.WriteHeader(http.StatusNoContent)
}

func main() {
	tasks = []Task{}
	currentID = 1

	http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getTasks(w, r)
		case http.MethodPost:
			addTask(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/tasks/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getTaskByID(w, r)
		case http.MethodPut:
			updateTask(w, r)
		case http.MethodDelete:
			deleteTask(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/tasks/completed/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPut {
			markTaskDone(w, r)
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/tasks/incomplete/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPut {
			markTaskIsnotdone(w, r)
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("Server starts on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

