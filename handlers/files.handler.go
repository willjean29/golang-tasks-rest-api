package handlers

import (
	"app/error"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var collections = []string{"tasks", "users"}

type FilesHandler struct{}

func (f *FilesHandler) UploadFile(w http.ResponseWriter, r *http.Request) {
	var isFoundCollection = false
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error.New("Invalid ID", http.StatusBadRequest, err))
		return
	}

	for _, collection := range collections {
		if collection == vars["collection"] {
			isFoundCollection = true
			break
		}
	}

	if !isFoundCollection {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error.New("Invalid collection", http.StatusBadRequest, err))
		return
	}
	fileUrl := r.Context().Value("fileUrl").(string)
	switch vars["collection"] {
	case "tasks":
		log.Println("Upload file for task with id:", id)
		task, errorApp := TaskService.GetTask(id)
		if errorApp.StatusCode != 0 {
			w.WriteHeader(errorApp.StatusCode)
			json.NewEncoder(w).Encode(errorApp)
			return
		}
		task.Image = fileUrl
		errorApp = TaskService.SaveTask(&task)
		if errorApp.StatusCode != 0 {
			w.WriteHeader(errorApp.StatusCode)
			json.NewEncoder(w).Encode(errorApp)
			return
		}
		json.NewEncoder(w).Encode(map[string]string{"message": "File uploaded successfully"})
	case "users":
		log.Println("Upload file for user with id:", id)
	}

}
