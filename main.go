package main

import (
	"app/db"
	"app/handlers"
	"app/routes"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	var defaultPort int64 = 4000
	db.DBConnection()
	port := strconv.FormatInt(defaultPort, 10)
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", handlers.IndexRoute)

	routes.TaskRoutes(router)

	log.Println("Running on port ", port)

	log.Fatal(http.ListenAndServe(":"+port, router))
}
