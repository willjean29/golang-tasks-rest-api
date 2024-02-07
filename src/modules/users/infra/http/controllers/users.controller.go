package controllers

import (
	"app/db"
	usecases "app/src/modules/users/app"
	"app/src/modules/users/infra/gorm/repositories"
	error "app/src/shared/errors"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

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
