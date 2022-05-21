package entities

type Product struct {
	ID          uint    `json:"id,omitempty"`
	Name        string  `json:"name,omitempty"`
	Price       float64 `json:"price,omitempty"`
	Description string  `json:"description,omitempty"`
}
