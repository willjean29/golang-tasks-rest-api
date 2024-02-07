package controllers

import (
	"app/db"
	usecases "app/src/modules/users/app"
	"app/src/modules/users/infra/gorm/repositories"
	"encoding/json"
	"net/http"
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
