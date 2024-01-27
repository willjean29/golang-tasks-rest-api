package main

import (
	"app/handlers"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	var defaultPort int64 = 4000
	port := strconv.FormatInt(defaultPort, 10)
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", handlers.IndexRoute)
	router.HandleFunc("/tasks", handlers.GetTasks).Methods(http.MethodGet)
	router.HandleFunc("/tasks", handlers.CreateTask).Methods(http.MethodPost)
	router.HandleFunc("/tasks/{id}", handlers.GetTask).Methods(http.MethodGet)
	router.HandleFunc("/tasks/{id}", handlers.DeleteTask).Methods(http.MethodDelete)
	router.HandleFunc("/tasks/{id}", handlers.UpdateTask).Methods(http.MethodPut)

	fmt.Println("Running on port ", port)

	log.Fatal(http.ListenAndServe(":"+port, router))
}
