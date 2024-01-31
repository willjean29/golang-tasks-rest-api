package main

import (
	"app/db"
	"app/error"
	"app/handlers"
	"app/models"
	"app/routes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	var defaultPort int64 = 4000
	db.DBConnection()
	db.DB.AutoMigrate(&models.Task{})
	port := strconv.FormatInt(defaultPort, 10)
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", handlers.IndexRoute)

	router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(error.New("Endpoint not found", http.StatusFound, errors.New("Not found - "+r.URL.Path)))
	})

	routes.TaskRoutes(router)

	log.Println("Running on port ", port)

	log.Fatal(http.ListenAndServe(":"+port, router))
}
