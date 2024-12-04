package main

import (
	"log"
	"net/http"
	"strconv"
	tasks "to-do-list-app/todo"
)

func main() {
	http.HandleFunc("/tasks", handleTasks)
	http.HandleFunc("/tasks/", handleTask)
	log.Println("Server is running on port 9090...")
	log.Fatal(http.ListenAndServe(":9090", nil))
}

func handleTasks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		tasks.ListTasks(w, r)
	case http.MethodPost:
		tasks.CreateTask(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handleTask(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/tasks/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		tasks.GetTask(w, r, id)
	case http.MethodDelete:
		tasks.DeleteTask(w, r, id)
	case http.MethodPut:
		tasks.UpdateTask(w, r, id) // Added the PUT method to update a task
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
