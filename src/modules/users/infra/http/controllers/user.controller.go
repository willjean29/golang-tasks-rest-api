package controllers

import (
	usecases "app/src/modules/users/app"
	"app/src/modules/users/domain/entities"
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
var userRepository = &repositories.UserRepository{
	Datasource: &datasource.GormUserDatasource{},
	// Datasource: &datasource.FileSistemUserDatasource{},
}

type UserController struct{}

func (u *UserController) List(w http.ResponseWriter, r *http.Request) {
	listUserUseCase := usecases.ListUserUseCase{
		UserRepository: userRepository,
	}
	users, errorApp := listUserUseCase.Execute()
	if errorApp.StatusCode != 0 {
		w.WriteHeader(errorApp.StatusCode)
		json.NewEncoder(w).Encode(errorApp)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func (u *UserController) Show(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error.New("Invalid ID", http.StatusBadRequest, err))
		return
	}
	getUserUseCase := usecases.GetUserUseCase{
		UserRepository: userRepository,
	}
	user, errorApp := getUserUseCase.Execute(userId)
	if errorApp.StatusCode != 0 {
		w.WriteHeader(errorApp.StatusCode)
		json.NewEncoder(w).Encode(errorApp)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func (u *UserController) Login(w http.ResponseWriter, r *http.Request) {
	var createSession entities.CreateSession

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
	createSessionUseCase := usecases.CreateSessionUseCase{
		UserRepository: userRepository,
	}

	user, token, errorApp := createSessionUseCase.Execute(createSession)
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

func (u *UserController) Register(w http.ResponseWriter, r *http.Request) {
	var createUser entities.CreateUser

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
	createUserUseCase := usecases.CreateUserUseCase{
		UserRepository: userRepository,
	}

	user, token, errorApp := createUserUseCase.Execute(createUser)
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
