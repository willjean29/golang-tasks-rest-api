package controllers

import (
	error "app/src/shared/errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FileControllerGin struct{}

func (f *FileControllerGin) UploadFile(ctx *gin.Context) {
	var isFoundCollection = false
	paramId := ctx.Param("id")
	id, err := strconv.Atoi(paramId)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, error.New("Invalid ID", http.StatusBadRequest, err))
		return
	}

	for _, collection := range collections {
		if collection == ctx.Param("collection") {
			isFoundCollection = true
			break
		}
	}

	if !isFoundCollection {
		ctx.JSON(http.StatusBadRequest, error.New("Invalid collection", http.StatusBadRequest, err))
		return
	}
	filename := ctx.MustGet("filename").(string)
	switch ctx.Param("collection") {
	case "tasks":
		log.Println("Upload file for task with id:", id)
		message, errorApp := uploadTaskFile(id, filename)
		if errorApp.StatusCode != 0 {
			ctx.JSON(errorApp.StatusCode, errorApp)
			return
		}
		ctx.JSON(http.StatusOK, message)

	case "users":
		log.Println("Upload file for user with id:", id)
		message, errorApp := uploadUserFile(id, filename)
		if errorApp.StatusCode != 0 {
			ctx.JSON(errorApp.StatusCode, errorApp)
			return
		}
		ctx.JSON(http.StatusOK, message)
	}
}
