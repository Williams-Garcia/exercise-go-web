package impl

import (
	"api_rest/internal/domain"
	"api_rest/internal/product"
)

type ProductService struct {
	serviceImpl product.ProductServiceImpl
}

func (ps *ProductService) GetProducts() (products []domain.Product, err error) {
	return ps.serviceImpl.GetProducts()
}

func (ps *ProductService) GetProduct(id int) (product domain.Product, err error) {
	return ps.serviceImpl.GetProduct(id)
}

func (ps *ProductService) SearchProduct(query float64) (products []domain.Product, err error) {
	return ps.serviceImpl.SearchProduct(query)
}

func (ps *ProductService) AddProduct(product domain.Product) (newProduct domain.Product, err error) {
	return ps.serviceImpl.AddProduct(product)
}

func NewProductService(rp product.ProductServiceImpl) *ProductService {
	return &ProductService{
		serviceImpl: rp,
	}
}
