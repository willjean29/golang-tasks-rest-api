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

	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Println(r.Method, r.URL.Path)
			w.Header().Set("Content-Type", "application/json")
			next.ServeHTTP(w, r)
		})
	})

	router.HandleFunc("/", handlers.IndexRoute)
	routes.TaskRoutes(router)

	router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(error.New("Endpoint not found", http.StatusNotFound, errors.New("Not found - "+r.URL.Path)))
	})

	log.Println("Running on port ", port)

	log.Fatal(http.ListenAndServe(":"+port, router))
}
