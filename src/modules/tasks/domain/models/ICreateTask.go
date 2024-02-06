package models

type ICreateTask struct {
	Name    string `json:"name" validate:"required"`
	Content string `json:"content" validate:"required"`
}
