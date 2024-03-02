package models

type CreateInput struct {
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}
type Product struct {
	Id    uint32  `json:"id"`
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}
