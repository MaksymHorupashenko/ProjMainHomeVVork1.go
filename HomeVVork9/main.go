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

var users = map[string]string{
	"Galyna1986": "1986",
	"Andrii1991": "1991",
	"Olena1989":  "1989",
}

func BasicAuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, pass, ok := r.BasicAuth()
		if !ok || !checkCredentials(user, pass) {
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			http.Error(w, "Unauthorized.", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	}
}

func checkCredentials(username, password string) bool {
	storedPassword, ok := users[username]
	if !ok {
		return false
	}
	return storedPassword == password
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
	store.mu.RLock()
	defer store.mu.RUnlock()

	vars := mux.Vars(r)
	studentID := vars["id"]

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
			{ID: "4", FullName: "Саша Пугаченко", Grades: Grades{"Математика": 12, "Фізика": 12}},
			{ID: "5", FullName: "Сергій Ковальчук", Grades: Grades{"Математика": 9, "Фізика": 10}},
			{ID: "6", FullName: "Артем Бондаренко", Grades: Grades{"Математика": 10, "Фізика": 9}},
		},
	}

	router := mux.NewRouter()

	router.HandleFunc("/", BasicAuthMiddleware(indexHandler))

	router.HandleFunc("/student/{id}", BasicAuthMiddleware(getStudentHandler)).Methods("GET")

	fmt.Println("Сервер запущено на порті 8080...")
	if err := http.ListenAndServe(":8080", router); err != nil {
		fmt.Println("Помилка запуску сервера:", err)
	}
}
