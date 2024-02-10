package entities

type CreateTask struct {
	Name    string `json:"name" validate:"required"`
	Content string `json:"content" validate:"required"`
}
