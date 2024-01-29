package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Content string `json:"content"`
}
type ListTask []Task
