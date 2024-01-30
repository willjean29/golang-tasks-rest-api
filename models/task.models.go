package models

import (
	"encoding/json"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Name    string `validate:"required"`
	Content string `validate:"required"`
}

type TaskJSON struct {
	ID        uint           `json:"id"`
	Name      string         `json:"name"`
	Content   string         `json:"content"`
	CreatedAt string         `json:"created_at"`
	UpdatedAt string         `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func (t *Task) MarshalJSON() ([]byte, error) {
	taskJSON := TaskJSON{
		ID:        t.ID,
		Name:      t.Name,
		Content:   t.Content,
		CreatedAt: t.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: t.UpdatedAt.Format("2006-01-02 15:04:05"),
		DeletedAt: t.DeletedAt,
	}
	return json.Marshal(taskJSON)
}

type ListTask []Task
