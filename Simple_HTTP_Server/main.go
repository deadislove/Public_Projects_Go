package main

import (
	"Simple_HTTP_Server/database"
	"Simple_HTTP_Server/models"
	"Simple_HTTP_Server/services"
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	database.InitDB()
	defer database.DB.Close()

	http.HandleFunc("/items", itemHandler)

	fmt.Println("Server running on port 8080...")
	http.ListenAndServe(":8080", nil)
}

func itemHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		items, err := services.GetAllItems()
		if err != nil {
			http.Error(w, "Error fetching items", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(items)

	case "POST":
		var newItem models.ItemDTO
		if err := json.NewDecoder(r.Body).Decode(&newItem); err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}
		if err := services.CreateItem(newItem); err != nil {
			http.Error(w, "Error creating item", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newItem)

	case "PUT":
		var updatedItem models.ItemDTO
		if err := json.NewDecoder(r.Body).Decode(&updatedItem); err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}
		if err := services.UpdateItem(updatedItem); err != nil {
			http.Error(w, "Error updating item", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(updatedItem)

	case "DELETE":
		id := r.URL.Query().Get("id")
		if id == "" {
			http.Error(w, "Missing ID parameter", http.StatusBadRequest)
			return
		}
		if err := services.DeleteItem(id); err != nil {
			http.Error(w, "Error deleting item", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
