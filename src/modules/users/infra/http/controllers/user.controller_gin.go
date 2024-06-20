package controllers

import (
	usecases "app/src/modules/users/app"
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
