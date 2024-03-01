package models

type CreateInput struct {
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}
