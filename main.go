package main

import (
	"app/db"
	"app/error"
	"app/handlers"
	"app/middlewares"
	"app/models"
	"app/routes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

func uploadFileHandler(w http.ResponseWriter, r *http.Request) {
	// Parsea la solicitud
	r.ParseMultipartForm(10 << 20) // 10mb limite

	// Obtiene el archivo de la solicitud
	file, handler, err := r.FormFile("file")
	if err != nil {
		json.NewEncoder(w).Encode(error.New("Error retrieving the file", http.StatusBadRequest, err))
		return
	}
	defer file.Close()

	// Crea la carpeta destino si no existe
	err = os.MkdirAll("uploads", os.ModePerm)
	if err != nil {
		json.NewEncoder(w).Encode(error.New("Error creating the folder", http.StatusInternalServerError, err))
		return
	}

	// Crea un nuevo archivo en el servidor
	newFile, err := os.Create("uploads/" + handler.Filename)
	if err != nil {
		json.NewEncoder(w).Encode(error.New("Error creating the file", http.StatusInternalServerError, err))
		return
	}
	defer newFile.Close()

	// Copia el contenido del archivo subido al nuevo archivo
	_, err = io.Copy(newFile, file)

	if err != nil {
		json.NewEncoder(w).Encode(error.New("Error copying the file", http.StatusInternalServerError, err))
		return
	}

	// Envía una respuesta de éxito
	json.NewEncoder(w).Encode(map[string]string{"file": handler.Filename, "message": "File uploaded successfully"})
}

func main() {
	var defaultPort int64 = 4000
	db.DBConnection()
	db.DB.AutoMigrate(&models.Task{})
	port := strconv.FormatInt(defaultPort, 10)
	router := mux.NewRouter().StrictSlash(true)

	// statics files to filesystem
	fs := http.FileServer(http.Dir("./uploads/"))
	router.PathPrefix("/uploads/").Handler(http.StripPrefix("/uploads/", fs))

	router.Use(middlewares.ContentType)

	router.HandleFunc("/", handlers.IndexRoute)
	router.HandleFunc("/upload", uploadFileHandler).Methods(http.MethodPost)
	routes.TaskRoutes(router)

	router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(error.New("Endpoint not found", http.StatusNotFound, errors.New("Not found - "+r.URL.Path)))
	})

	log.Println("Running on port ", port)

	log.Fatal(http.ListenAndServe(":"+port, router))
}
