package domain

type Product struct {
	Id          int     `json:"id"`
	Name        string  `json:"name" validate:"required"`
	Quantity    int     `json:"quantity" validate:"required"`
	CodeValue   string  `json:"code_value" validate:"required"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration" validate:"required"`
	Price       float64 `json:"price" validate:"required"`
}

type ProductPatch struct {
	Name        string  `json:"name,omitempty"`
	Quantity    int     `json:"quantity,omitempty"`
	CodeValue   string  `json:"code_value,omitempty"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration,omitempty"`
	Price       float64 `json:"price,omitempty"`
}

type ListProducts struct {
	Products []Product `json:"products"`
}
