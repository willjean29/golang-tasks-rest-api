package main

import (
	"app/db"
	"app/error"
	"app/middlewares"
	"app/models"
	"app/routes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var defaultPort int64 = 4000

	// connection database
	db.DBConnection()
	db.DB.AutoMigrate(&models.Task{}, &models.User{})

	port := strconv.FormatInt(defaultPort, 10)
	router := mux.NewRouter().StrictSlash(true)

	// statics files to filesystem
	fs := http.FileServer(http.Dir("./uploads/"))
	router.PathPrefix("/uploads/").Handler(http.StripPrefix("/uploads/", fs))

	router.Use(middlewares.ContentType)

	// routes for application
	routes.TaskRoutes(router)
	routes.FilesRoutes(router)
	routes.UserRoutes(router)

	router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(error.New("Endpoint not found", http.StatusNotFound, errors.New("Not found - "+r.URL.Path)))
	})

	log.Println("Running on port ", port)

	log.Fatal(http.ListenAndServe(":"+port, router))
}
