package tasks

import (
	"encoding/json"
	"net/http"
	"sync"
)

type Task struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var (
	tasks  = []Task{}
	nextID = 1
	mu     sync.Mutex
)

func ListTasks(w http.ResponseWriter, r *http.Request) {
	mu.Lock()         // Acquire the lock to ensure only one goroutine can modify the tasks slice at a time
	defer mu.Unlock() // Ensure the lock is released after this function finishes (even if there's a panic or return)

	json.NewEncoder(w).Encode(tasks)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	var task Task
	/*
		err := json.NewDecoder(r.Body).Decode(&task)  // Decode JSON into task
		if err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}
	*/
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	task.ID = nextID
	nextID++
	tasks = append(tasks, task)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

func GetTask(w http.ResponseWriter, r *http.Request, id int) {
	mu.Lock()
	defer mu.Unlock()

	for _, task := range tasks {
		if task.ID == id {
			json.NewEncoder(w).Encode(task)
			return
		}
	}

	http.Error(w, "Task not found", http.StatusNotFound)
}

func DeleteTask(w http.ResponseWriter, r *http.Request, id int) {
	mu.Lock()
	defer mu.Unlock()

	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.Error(w, "Task not found", http.StatusNotFound)
}

// Add the UpdateTask (PUT method)
func UpdateTask(w http.ResponseWriter, r *http.Request, id int) {
	mu.Lock()
	defer mu.Unlock()

	var updatedTask Task
	if err := json.NewDecoder(r.Body).Decode(&updatedTask); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	for i, task := range tasks {
		if task.ID == id {
			// Update the task's title and content with new data
			tasks[i].Title = updatedTask.Title
			tasks[i].Content = updatedTask.Content

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(tasks[i]) // Send the updated task back as a response
			return
		}
	}

	http.Error(w, "Task not found", http.StatusNotFound)
}
