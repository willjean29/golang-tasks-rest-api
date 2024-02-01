package models

import (
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Name    string
	Content string
	Image   string `gorm:"default:null"`
}

type TaskJSON struct {
	ID        uint           `json:"id"`
	Name      string         `json:"name"`
	Content   string         `json:"content"`
	Image     string         `json:"image"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func (t *Task) MarshalJSON() ([]byte, error) {
	taskJSON := TaskJSON{
		ID:        t.ID,
		Name:      t.Name,
		Content:   t.Content,
		Image:     t.Image,
		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
		DeletedAt: t.DeletedAt,
	}
	return json.Marshal(taskJSON)
}

type ListTask []Task
