package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type task struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Content string `json:"content"`
}
type ListTask []task

var tasks = ListTask{
	{
		ID:      1,
		Name:    "Task One",
		Content: "Some Content",
	},
	{
		ID:      2,
		Name:    "Task Two",
		Content: "Some Content",
	},
}

func getTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(tasks)
}

func createTask(w http.ResponseWriter, r *http.Request) {
	var newTask task
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Insert a Valid Task Data")
	}
	json.Unmarshal(reqBody, &newTask)
	newTask.ID = len(tasks) + 1
	tasks = append(tasks, newTask)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTask)
}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my API!")
}

func main() {
	var defaultPort int64 = 4000
	port := strconv.FormatInt(defaultPort, 10)
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/tasks", getTasks).Methods("GET")
	router.HandleFunc("/tasks", createTask).Methods("POST")

	fmt.Println("Running on port ", port)

	log.Fatal(http.ListenAndServe(":"+port, router))
}
