package controllers

import (
	usecases "app/src/modules/users/app"
	"app/src/modules/users/domain/entities"
	error "app/src/shared/errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserControllerGin struct{}

func (u *UserControllerGin) List(c *gin.Context) {

	listUserUseCase := usecases.ListUserUseCase{
		UserRepository: userRepository,
	}

	users, errorApp := listUserUseCase.Execute()
	if errorApp.StatusCode != 0 {
		c.JSON(errorApp.StatusCode, errorApp)
	}

	c.JSON(http.StatusOK, users)
}

func (u *UserControllerGin) Show(c *gin.Context) {
	id := c.Param("id")
	userId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, error.New("Invalid ID", http.StatusBadRequest, err))
		return
	}
	getUserUseCase := usecases.GetUserUseCase{
		UserRepository: userRepository,
	}
	user, errorApp := getUserUseCase.Execute(userId)
	if errorApp.StatusCode != 0 {
		c.JSON(errorApp.StatusCode, errorApp)
		return
	}

	c.JSON(http.StatusOK, user)
}

func (u *UserControllerGin) Login(c *gin.Context) {
	var createSession entities.CreateSession

	if err := c.BindJSON(&createSession); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Insert a Valid User Data", "details": err.Error()})
		return
	}
	err := userValidator.ValidateLogin(createSession)
	if err != nil {
		c.JSON(http.StatusBadRequest, error.New("Insert a Valid User Data", http.StatusBadRequest, err))
		return
	}
	createSessionUseCase := usecases.CreateSessionUseCase{
		UserRepository: userRepository,
	}

	user, token, errorApp := createSessionUseCase.Execute(createSession)
	if errorApp.StatusCode != 0 {
		c.JSON(errorApp.StatusCode, errorApp)
		return
	}

	c.SetCookie("token", token, 60*60*24, "/", "", true, true)
	c.JSON(http.StatusCreated, user)
}

func (u *UserControllerGin) Register(c *gin.Context) {
	var createUser entities.CreateUser

	if err := c.BindJSON(&createUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Insert a Valid User Data", "details": err.Error()})
		return
	}

	err := userValidator.ValidateCreateUser(createUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, error.New("Insert a Valid User Data", http.StatusBadRequest, err))
		return
	}

	createUserUseCase := usecases.CreateUserUseCase{
		UserRepository: userRepository,
	}

	user, token, errorApp := createUserUseCase.Execute(createUser)
	if errorApp.StatusCode != 0 {
		c.JSON(errorApp.StatusCode, errorApp)
		return
	}

	c.SetCookie("token", token, 60*60*24, "/", "", true, true)
	c.JSON(http.StatusCreated, user)
}
