package models

import (
	"app/src/modules/tasks/infra/data/gorm/models"
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Email    string `gorm:"unique"`
	Password string
	Image    string `gorm:"default:null"`
	Tasks    []models.Task
}

type UserJSON struct {
	ID        uint           `json:"id"`
	Name      string         `json:"name"`
	Email     string         `json:"email"`
	Password  string         `json:"password"`
	Image     string         `json:"image"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
	Tasks     []models.Task  `json:"tasks"`
}

func (u *User) MarshalJSON() ([]byte, error) {
	userJSON := UserJSON{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		Password:  u.Password,
		Image:     u.Image,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
		DeletedAt: u.DeletedAt,
		Tasks:     u.Tasks,
	}
	return json.Marshal(userJSON)
}

type ListUsers []User
