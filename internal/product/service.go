package product

import "api_rest/internal/domain"

type ProductServiceImpl interface {
	GetProducts() (products []domain.Product, err error)
	GetProduct(id int) (product domain.Product, err error)
	SearchProduct(price float64) (products []domain.Product, err error)
	AddProduct(product domain.Product) (newProduct domain.Product, err error)
	UpdateProduct(id int, product domain.Product) (updatedProduct domain.Product, err error)
	UpdatePatchProduct(id int, product domain.ProductPatch) (updatedProduct domain.Product, err error)
	DeleteProduct(id int) (err error)
	ExistProduct(productName string) bool
	UniqueCodeValue(productCode string) bool
}
