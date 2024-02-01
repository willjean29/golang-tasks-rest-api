package dtos

type CreateTaskDto struct {
	Name    string `json:"name" validate:"required"`
	Content string `json:"content" validate:"required"`
}

type UpdateTaskDto struct {
	Name    string `json:"name" validate:"omitempty"`
	Content string `json:"content" validate:"omitempty"`
}
