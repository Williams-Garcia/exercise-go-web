package product

import "api_rest/internal/domain"

type ProductServiceImpl interface {
	GetProducts() (products []domain.Product, err error)
	GetProduct(id int) (product domain.Product, err error)
	SearchProduct(price float64) (products []domain.Product, err error)
	AddProduct(product domain.Product) (newProduct domain.Product, err error)
}
