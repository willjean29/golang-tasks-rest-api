package handlers

import (
	"app/dtos"
	"app/error"
	"app/models"
	hash "app/providers/HashProvider"
	token "app/providers/TokenProvider"
	"app/services"
	"app/utils"
	"app/validators"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var userValidator validators.UserValidator = validators.NewUserValidator()
var userService services.UserService = services.UserService{}
var hashProvider hash.HashProvider = &hash.BcryptProvider{}
var tokenProvider token.TokenProvider = &token.JwtProvider{}

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

	err = hashProvider.ComparePasswords(user.Password, loginDto.Password)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(error.New("Invalid data (password)", http.StatusUnauthorized, err))
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

	hashPassword, _ := hashProvider.HashPassword(createUserDto.Password)
	newUser.Password = hashPassword

	errorApp = userService.CreateUser(&newUser)
	if errorApp.StatusCode != 0 {
		w.WriteHeader(errorApp.StatusCode)
		json.NewEncoder(w).Encode(errorApp)
		return
	}
	userJson, _ := newUser.MarshalJSON()
	w.Write(userJson)
}
