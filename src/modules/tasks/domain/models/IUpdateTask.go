package models

type IUpdateTask struct {
	Name    string `json:"name" validate:"omitempty"`
	Content string `json:"content" validate:"omitempty"`
}
