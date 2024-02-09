package middlewares

import (
	error "app/src/shared/errors"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

func UploadFile(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Parsea la solicitud
		r.ParseMultipartForm(10 << 20) // 10mb limite

		// Obtiene el archivo de la solicitud
		file, handler, err := r.FormFile("file")
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(error.New("Error retrieving the file", http.StatusBadRequest, err))
			return
		}
		defer file.Close()

		// Crea la carpeta destino si no existe
		err = os.MkdirAll("uploads", os.ModePerm)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(error.New("Error creating the folder", http.StatusInternalServerError, err))
			return
		}

		// Crea un nuevo archivo en el servidor
		newFile, err := os.Create("uploads/" + handler.Filename)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
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
		// se crea la url del archivo
		log.Println("filename ", handler.Filename)
		ctx := context.WithValue(r.Context(), "filename", handler.Filename)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
