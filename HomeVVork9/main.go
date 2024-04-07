package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
)

type ClassInfo struct {
	ClassName string    `json:"class_name"`
	Students  []Student `json:"students"`
}

type Student struct {
	ID       string `json:"id"`
	FullName string `json:"full_name"`
	Grades   Grades `json:"grades"`
}

type Teacher struct {
	ID       string `json:"id"`
	FullName string `json:"full_name"`
}

var teachersByClassID map[string][]Teacher

type Grades map[string]float64

type ClassStore struct {
	classInfo map[string]ClassInfo
	mu        sync.RWMutex
}

var store ClassStore

func isAuthenticated(r *http.Request) bool {
	vars := mux.Vars(r)
	classID := vars["classID"]
	teacherID := vars["teacherID"]
	teachers, ok := teachersByClassID[classID]
	if !ok {
		return false
	}

	for _, t := range teachers {
		if t.ID == teacherID {
			return true
		}
	}

	return false
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	store.mu.RLock()
	defer store.mu.RUnlock()
	classInfo, ok := store.classInfo["10-А"]
	if !ok {
		http.Error(w, "Інформація про клас не знайдена", http.StatusNotFound)
		return
	}

	jsonData, err := json.Marshal(classInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func getStudentHandler(w http.ResponseWriter, r *http.Request) {
	if !isAuthenticated(r) {
		http.Error(w, "Доступ заборонено", http.StatusForbidden)
		return
	}

	vars := mux.Vars(r)
	studentID := vars["id"]


	store.mu.RLock()
	defer store.mu.RUnlock()
	for _, classInfo := range store.classInfo {
		for _, student := range classInfo.Students {
			if student.ID == studentID {
				jsonData, err := json.Marshal(student)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				w.Write(jsonData)
				return
			}
		}
	}

	http.Error(w, "Учень не знайдений", http.StatusNotFound)
}

func main() {
	store.classInfo = make(map[string]ClassInfo)
	store.classInfo["10-А"] = ClassInfo{
		ClassName: "10-А",
		Students: []Student{
			{ID: "1", FullName: "Олександра Петренко", Grades: Grades{"Математика": 10, "Фізика": 9}},
			{ID: "2", FullName: "Максим Іванов", Grades: Grades{"Математика": 9, "Фізика": 8}},
			{ID: "3", FullName: "Лариса Коваленко", Grades: Grades{"Математика": 11, "Фізика": 9}},
			{ID: "4", FullName: "Віталій Сидоренко", Grades: Grades{"Математика": 10, "Фізика": 9}},
			{ID: "5", FullName: "Сергій Ковальчук", Grades: Grades{"Математика": 9, "Фізика": 10}},
			{ID: "6", FullName: "Артем Бондаренко", Grades: Grades{"Математика": 10, "Фізика": 9}},
		},
	}

	router := mux.NewRouter()
)
	router.HandleFunc("/", indexHandler)

	router.HandleFunc("/student/{id}", getStudentHandler).Methods("GET")
  
	fmt.Println("Сервер запущено на порті 8080...")
	if err := http.ListenAndServe(":8080", router); err != nil {
		fmt.Println("Помилка запуску сервера:", err)
	}
}
