package controllers

import (
	"app/db"
	token "app/providers/TokenProvider"
	usecases "app/src/modules/users/app"
	"app/src/modules/users/domain/models"
	"app/src/modules/users/infra/gorm/repositories"
	error "app/src/shared/errors"
	"app/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

var tokenProvider token.TokenProvider = token.NewJwtProvider()

type UsersController struct{}

func (u *UsersController) List(w http.ResponseWriter, r *http.Request) {
	usecase := usecases.ListUserUseCase{
		UserRepository: &repositories.UsersRepository{
			Repository: db.DB,
		},
	}
	users, errorApp := usecase.Execute()
	if errorApp.StatusCode != 0 {
		w.WriteHeader(errorApp.StatusCode)
		json.NewEncoder(w).Encode(errorApp)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func (u *UsersController) Show(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error.New("Invalid ID", http.StatusBadRequest, err))
		return
	}
	usecase := usecases.GetUserUseCase{
		UserRepository: &repositories.UsersRepository{
			Repository: db.DB,
		},
	}
	user, errorApp := usecase.Execute(userId)
	if errorApp.StatusCode != 0 {
		w.WriteHeader(errorApp.StatusCode)
		json.NewEncoder(w).Encode(errorApp)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func (u *UsersController) Login(w http.ResponseWriter, r *http.Request) {
	var createSession models.ICreateSession

	err := utils.TransformBody(r.Body, &createSession)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error.New("Insert a Valid Task Data", http.StatusBadRequest, err))
		return
	}

	usecase := usecases.CreateSessionUseCase{
		UserRepository: &repositories.UsersRepository{
			Repository: db.DB,
		},
	}

	user, errorApp := usecase.Execute(createSession)
	if errorApp.StatusCode != 0 {
		w.WriteHeader(errorApp.StatusCode)
		json.NewEncoder(w).Encode(errorApp)
		return
	}

	tokenString, err := tokenProvider.GenerateToken("userId", fmt.Sprintf("%d", user.ID))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(error.New("Internal server error", http.StatusInternalServerError, err))
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: time.Now().Add(24 * time.Hour), // Ajusta la duración de la cookie según tus necesidades
	})

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}
