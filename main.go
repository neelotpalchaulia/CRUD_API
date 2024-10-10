package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"` // "pending" or "completed"
}

var tasks []Task
var nextID = 1

// Create function (to create a new task) by Neel
func createTaskHandler_Neel(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var newTask Task
	err := json.NewDecoder(r.Body).Decode(&newTask)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	newTask.ID = nextID
	nextID++
	tasks = append(tasks, newTask)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newTask)
}

// Read function - to retrieve all created/existings tasks by Oviya

// Read function - to retireve a specific task wrt its ID by Srinidhi

func getTaskByIDHandler_Srinidhi(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodGet {
                http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
                return
        }

        // id := r.URL.Path[len("/tasks/"):] // Extract ID from URL using the native method
        id := mux.Vars(r)["id"] // Extract ID from URL using Gorilla Mux
        for _, task := range tasks {
                filtered_id := fmt.Sprintf("%d", task.ID)
                if filtered_id == id {
                        w.Header().Set("Content-Type", "application/json")
                        json.NewEncoder(w).Encode(task)
                        return
                }
        }

        http.Error(w, "Task not found", http.StatusNotFound)
}

// Update function (Update a task based on ID) created by Nauman

// Delete function (Delete a task based on ID) created by Anjani

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/tasks", createTaskHandler_Neel).Methods("POST") // Handle POST /tasks
	router.HandleFunc("/tasks/{id}", getTaskByIDHandler_Srinidhi).Methods("GET") // Handle GET /tasks/{id}
	fmt.Println("Server is running on port :8080")
	http.ListenAndServe(":8080", router)
}
