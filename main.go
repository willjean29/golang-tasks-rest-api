package main

import (
	"app/handlers"
	"app/routes"
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

	routes.TaskRoutes(router)

	fmt.Println("Running on port ", port)

	log.Fatal(http.ListenAndServe(":"+port, router))
}
