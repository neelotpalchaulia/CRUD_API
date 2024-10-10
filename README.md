# Building a Simple CRUD API in GO

This project demonstrates a RESTful CRUD(Create, Read, Update and Delete) API built using the GO programmimg language, `net/http` package, and the `gorilla/mux` router. This API manages a collection of tasks which has basic fileds like `ID`, `Title`, `Description`and `Status`.

## Features

- Create a Task: Add a new task to the list.
- Read Tasks: Retrieve a list of all tasks or a single task by its ID.
- Update a Task: Modify the details of an existing task.
- Delete a Task: Remove a task from the list.

## Getting Started

1. Clone the repository

   ```bash
   git clone https://github.com/neelotpalchaulia/CLOD2003_week-5_ICLA-3.git
   cd CLOD2003_week-5_ICLA-3
   ```

2. **Install Dependencies**: This project uses the `gorilla/mux` package. Install it using:

   ```bash
   go get -u github.com/gorilla/mux
   ```

3. Run the application:

   ```bash
   go run main.go
   ```

4. **API Endpoint:**

    - `POST /tasks`: Create a new task.
    - `GET /tasks`: Retrieve all tasks.
    - `GET /tasks/{id}`: Retrieve a specific task by its ID.
    - `PUT /tasks/{id}`: Update a specific task.
    - `DELETE /tasks/{id}`: Delete a specific task.

## Code Explaination

Below is a breakdown of the `main.go` file:

### Imports

   ```bash
   import (
   "encoding/json"
   "fmt"
   "net/http"
   "github.com/gorilla/mux"
   )
   ```

    - `encoding/json`: Handles JSON encoding and decoding.
    - `fmt`: Formats strings for output.
    - `net/http`: Provides HTTP client and server implementations.
    - `github.com/gorilla/mux`: A powerful HTTP router and URL matcher for building RESTful APIs.

### Define your Task struct

   ```bash
   type Task struct {
   ID          int    `json:"id"`
   Title       string `json:"title"`
   Description string `json:"description"`
   Status      string `json:"status"`
   }
   ```

    - Defines a `Task` struct with fields: `ID`, `Title`, `Description`, and `Status`.
    - The struct tags (`json:"fieldname"`) ensure proper JSON encoding/decoding while sending or receiving data.

### Global Variable

   ```bash
   var tasks []Task
   var nextID = 1
   ```

    - `tasks`: We declare this slice variable to store the list of tasks.
    - `nextID`: A counter to assign unique IDs to new tasks.

### Create a New Task (POST /tasks)

   ```bash
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
   ```

    - Purpose: Handles the creation of a new task.
    - The `createTaskHandler` checks if the HTTP request method is `POST`. If not, it returns a 405 error.
    - The new code is then decoded from the request body into a `Task` object. If there exists an error, a 400 error is returned.
    - Assigns a unique ID and appends the task to the `tasks` slice.
    - Responds with the created task in JSON format.

### Read All Tasks (GET /tasks)

   ```bash
   func getAllTasksHandler_Oviya(w http.ResponseWriter, r *http.Request) {
       if r.Method != http.MethodGet {
           http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
           return
       }

       w.Header().Set("Content-Type", "application/json")
       json.NewEncoder(w).Encode(tasks)
   }
   ```

    - Purpose: This handler checks if the request method is GET and returns the list of tasks in JSON format.

### Get a Task by its ID (GET /tasks/{id})

   ```bash
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
   ```

    - Purpose: This handler checks if the request method is **GET** Fetches a task by ID.
    - Extracts the `id` from the URL, searches for the task within the `tasks` slice, and returns it if found else a 404 error is displayed.

### Update an Existing Task (PUT /tasks/{id})

   ```bash
   func updateTaskHandler_Nauman(w http.ResponseWriter, r *http.Request) {
       if r.Method != http.MethodPut {
           http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
           return
       }

       // id := r.URL.Path[len("/tasks/"):] // Extract ID from URL
       id := mux.Vars(r)["id"]
       var updatedTask Task
       err := json.NewDecoder(r.Body).Decode(&updatedTask)
       if err != nil {
           http.Error(w, "Invalid request body", http.StatusBadRequest)
           return
       }

       for i, task := range tasks {
           filtered_id := fmt.Sprintf("%d", task.ID)
           if filtered_id == id {
               tasks[i].Title = updatedTask.Title
               tasks[i].Description = updatedTask.Description
               tasks[i].Status = updatedTask.Status
               w.Header().Set("Content-Type", "application/json")
               json.NewEncoder(w).Encode(tasks[i])
               return
           }
       }

       http.Error(w, "Task not found", http.StatusNotFound) 
   }
   ```

    - Purpose: This is very much similar to the GET handler, however it retrieves the ID from the URL, decodes the updated task from the request body, and then updates the task if it exists.

### Delete Task by ID (DELETE /tasks/{id})

   ```bash
   func deleteTaskHandler_Anjani(w http.ResponseWriter, r *http.Request) {
       if r.Method != http.MethodDelete {
           http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
           return
       }

       // id := r.URL.Path[len("/tasks/"):] // Extract ID from URL
       id := mux.Vars(r)["id"]
       for i, task := range tasks {
           filtered_id := fmt.Sprintf("%d", task.ID)
           if filtered_id == id {
               tasks = append(tasks[:i], tasks[i+1:]...) // Remove task from slice
               w.WriteHeader(http.StatusNoContent)       // No content
               return
           }
       } 

       http.Error(w, "Task not found", http.StatusNotFound)
   }
   ```

    - Purpose: This handler checks if the method is DELETE, retrieves the ID, and removes the task from the slice if found.
### Set up Routing 

   ```bash
   func main() {
       router := mux.NewRouter()

       router.HandleFunc("/tasks", createTaskHandler_Neel).Methods("POST")          // Handle POST /tasks
       router.HandleFunc("/tasks", getAllTasksHandler_Oviya).Methods("GET")         // Handle GET /tasks
       router.HandleFunc("/tasks/{id}", getTaskByIDHandler_Srinidhi).Methods("GET") // Handle GET /tasks/{id}
       router.HandleFunc("/tasks/{id}", updateTaskHandler_Nauman).Methods("PUT")    // Handle PUT /tasks/{id}
       router.HandleFunc("/tasks/{id}", deleteTaskHandler_Anjani).Methods("DELETE") // Handle DELETE /tasks/{id}

       fmt.Println("Server is running on port :8080")
       http.ListenAndServe(":8080", router) 
   }
   ```

    - Sets up routes for the API endpoints.
    - Uses `gorilla/mux` for routing and starts an HTTP server on port `8080`.
