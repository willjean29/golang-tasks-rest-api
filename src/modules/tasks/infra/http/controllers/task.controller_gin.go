package controllers

import (
	usecases "app/src/modules/tasks/app"

	"app/src/modules/tasks/domain/entities"
	error "app/src/shared/errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TaskControllerGin struct{}

func (c TaskControllerGin) List(ctx *gin.Context) {
	userId := ctx.MustGet("userId").(int)
	listTasksUseCase := usecases.ListTasksUseCase{
		TaskRepository: taskRepository,
	}

	tasks, errorApp := listTasksUseCase.Execute(uint(userId))

	if errorApp.StatusCode != 0 {
		ctx.JSON(errorApp.StatusCode, errorApp)
		return
	}

	ctx.JSON(http.StatusOK, tasks)
}

func (c TaskControllerGin) Show(ctx *gin.Context) {
	id := ctx.Param("id")
	taskId, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, error.New("Invalid ID", http.StatusBadRequest, err))
		return
	}
	getTaskUseCase := usecases.GetTaskUseCase{
		TaskRepository: taskRepository,
	}

	tasks, errorApp := getTaskUseCase.Execute(taskId)

	if errorApp.StatusCode != 0 {
		ctx.JSON(errorApp.StatusCode, errorApp)
		return
	}

	ctx.JSON(http.StatusOK, tasks)
}

func (c TaskControllerGin) Create(ctx *gin.Context) {
	var createTask entities.CreateTask
	userId := ctx.MustGet("userId").(int)

	if err := ctx.BindJSON(&createTask); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Insert a Valid Task Data", "details": err.Error()})
		return
	}

	err := taskValidator.ValidateCreateTask(createTask)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, error.New("Insert a Valid Task Data", http.StatusBadRequest, err))
		return
	}

	createTaskUseCase := usecases.CreateTaskUseCase{
		TaskRepository: taskRepository,
	}

	task, errorApp := createTaskUseCase.Execute(createTask, uint(userId))
	if errorApp.StatusCode != 0 {
		ctx.JSON(errorApp.StatusCode, errorApp)
		return
	}

	ctx.JSON(http.StatusCreated, task)
}

func (c TaskControllerGin) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	taskId, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, error.New("Invalid ID", http.StatusBadRequest, err))
		return
	}

	var updateTask entities.UpdateTask
	if err := ctx.BindJSON(&updateTask); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Insert a Valid Task Data", "details": err.Error()})
		return
	}

	err = taskValidator.ValidateUpdateTask(updateTask)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, error.New("Insert a Valid Task Data", http.StatusBadRequest, err))
		return
	}

	updateTaskUseCase := usecases.UpdateTaskUseCase{
		TaskRepository: taskRepository,
	}

	task, errorApp := updateTaskUseCase.Execute(updateTask, taskId)
	if errorApp.StatusCode != 0 {
		ctx.JSON(errorApp.StatusCode, errorApp)
		return
	}

	ctx.JSON(http.StatusOK, task)
}

func (c TaskControllerGin) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	taskId, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, error.New("Invalid ID", http.StatusBadRequest, err))
		return
	}

	deleteTaskUseCase := usecases.DeleteTaskUseCase{
		TaskRepository: taskRepository,
	}

	errorApp := deleteTaskUseCase.Execute(taskId)

	if errorApp.StatusCode != 0 {
		ctx.JSON(errorApp.StatusCode, errorApp)
		return
	}

	ctx.JSON(http.StatusOK, map[string]string{
		"message": "The task with ID " + strconv.Itoa(taskId) + " has been remove successfully",
	})
}
