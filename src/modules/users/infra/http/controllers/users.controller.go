package controllers

import (
	usecases "app/src/modules/users/app"
	"app/src/modules/users/domain/models"
	"app/src/modules/users/domain/validators"
	"app/src/modules/users/infra/datasource"
	"app/src/modules/users/infra/repositories"
	error "app/src/shared/errors"
	"app/src/shared/utils"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

var userValidator validators.UserValidator = validators.NewUserValidator()
var userRepository = &repositories.UsersRepository{
	Datasource: &datasource.UserDatasource{},
}

type UsersController struct{}

func (u *UsersController) List(w http.ResponseWriter, r *http.Request) {
	usecase := usecases.ListUserUseCase{
		UserRepository: userRepository,
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
		UserRepository: userRepository,
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
		json.NewEncoder(w).Encode(error.New("Insert a Valid User Data", http.StatusBadRequest, err))
		return
	}
	err = userValidator.ValidateLogin(createSession)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error.New("Insert a Valid User Data", http.StatusBadRequest, err))
		return
	}
	usecase := usecases.CreateSessionUseCase{
		UserRepository: userRepository,
	}

	user, token, errorApp := usecase.Execute(createSession)
	if errorApp.StatusCode != 0 {
		w.WriteHeader(errorApp.StatusCode)
		json.NewEncoder(w).Encode(errorApp)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   token,
		Expires: time.Now().Add(24 * time.Hour), // Ajusta la duración de la cookie según tus necesidades
	})

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func (u *UsersController) Register(w http.ResponseWriter, r *http.Request) {
	var createUser models.ICreateUser

	err := utils.TransformBody(r.Body, &createUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error.New("Insert a Valid User Data", http.StatusBadRequest, err))
		return
	}

	err = userValidator.ValidateCreateUser(createUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error.New("Insert a Valid User Data", http.StatusBadRequest, err))
		return
	}
	usecase := usecases.CreateUserUseCase{
		UserRepository: userRepository,
	}

	user, token, errorApp := usecase.Execute(createUser)
	if errorApp.StatusCode != 0 {
		w.WriteHeader(errorApp.StatusCode)
		json.NewEncoder(w).Encode(errorApp)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   token,
		Expires: time.Now().Add(24 * time.Hour), // Ajusta la duración de la cookie según tus necesidades
	})

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
