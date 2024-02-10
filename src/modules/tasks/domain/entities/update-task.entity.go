package entities

type UpdateTask struct {
	Name    string `json:"name" validate:"omitempty"`
	Content string `json:"content" validate:"omitempty"`
}
