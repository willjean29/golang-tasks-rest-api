package middlewares

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

// Middleware para manejar la subida de imágenes
func UploadFileGin() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Limitar el tamaño del archivo (ej. 10MB)
		c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, 10<<20)

		// Parsear el multipart form
		if err := c.Request.ParseMultipartForm(10 << 20); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "El archivo es demasiado grande"})
			c.Abort()
			return
		}

		// Obtener el archivo
		file, header, err := c.Request.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Error al obtener el archivo"})
			c.Abort()
			return
		}
		defer file.Close()

		// Validar el tipo de archivo (opcional)
		if header.Header.Get("Content-Type") != "image/jpeg" && header.Header.Get("Content-Type") != "image/png" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Formato de archivo no soportado"})
			c.Abort()
			return
		}

		// Guardar el archivo
		filename := filepath.Base(header.Filename)
		println(filename)
		if err := c.SaveUploadedFile(header, "./uploads/"+filename); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo guardar el archivo"})
			c.Abort()
			return
		}

		// Pasar al siguiente middleware o manejador
		c.Set("filename", filename)
		c.Next()
	}
}
