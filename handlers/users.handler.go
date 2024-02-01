package handlers

import (
	"app/dtos"
	"app/error"
	"app/models"
	"app/services"
	"app/utils"
	"app/validators"
	"encoding/json"
	"net/http"
)

var userValidator validators.UserValidator = validators.NewUserValidator()
var userService services.UserService = services.UserService{}

type UserHandler struct{}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var user models.User
	var loginDto dtos.LoginDto
	var errorApp error.Error
	err := utils.TransformBody(r.Body, &loginDto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error.New("Invalid data", http.StatusBadRequest, err))
		return
	}

	err = userValidator.ValidateLogin(loginDto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error.New("Invalid data", http.StatusBadRequest, err))
		return
	}

	user, errorApp = userService.GetUserByEmail(loginDto.Email)
	if errorApp.StatusCode != 0 {
		w.WriteHeader(errorApp.StatusCode)
		json.NewEncoder(w).Encode(errorApp)
		return
	}

	if user.Password != loginDto.Password {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(error.New("Invalid data (password)", http.StatusUnauthorized, nil))
		return
	}

	userJson, _ := user.MarshalJSON()
	w.Write(userJson)
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var newUser models.User
	var createUserDto dtos.CreateUserDto
	var errorApp error.Error
	err := utils.TransformBody(r.Body, &createUserDto, &newUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error.New("Invalid data", http.StatusBadRequest, err))
		return
	}

	err = userValidator.ValidateCreateUser(createUserDto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error.New("Invalid data", http.StatusBadRequest, err))
		return
	}

	errorApp = userService.CreateUser(&newUser)
	if errorApp.StatusCode != 0 {
		w.WriteHeader(errorApp.StatusCode)
		json.NewEncoder(w).Encode(errorApp)
		return
	}
	userJson, _ := newUser.MarshalJSON()
	w.Write(userJson)
}
