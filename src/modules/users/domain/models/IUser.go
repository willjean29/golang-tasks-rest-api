package models

import (
	"app/src/modules/tasks/domain/models"
	"time"
)

type IUser struct {
	ID        uint           `json:"id"`
	Name      string         `json:"name"`
	Email     string         `json:"email"`
	Password  string         `json:"password"`
	Image     string         `json:"image"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	Tasks     []models.ITask `json:"tasks"`
}

type IListUsers []IUser
