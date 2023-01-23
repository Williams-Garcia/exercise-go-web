package impl

import (
	"api_rest/internal/domain"
	"api_rest/internal/product"
	msg "api_rest/internal/product"
	"time"
)

type ProductService struct {
	serviceImpl product.ProductServiceImpl
}

// func (rp *ServiceImpl) ReadFile() (err error) {
// 	data, err := ioutil.ReadFile("products.json")
// 	if err != nil {
// 		return errOpen
// 	}

// 	err = json.Unmarshal(data, &rp.ListProducts)
// 	if err != nil {
// 		return ErrParser
// 	}

// 	return err
// }

func (ps *ProductService) GetProducts() (products []domain.Product, err error) {
	return ps.serviceImpl.GetProducts()
}

func (ps *ProductService) GetProduct(id int) (product domain.Product, err error) {
	return ps.serviceImpl.GetProduct(id)
	// var foundProduct domain.Product
	// for _, product := range rp.Products {
	// 	if product.Id == id {
	// 		foundProduct = product
	// 		break
	// 	}
	// }

	// if foundProduct.Id != id {
	// 	return domain.Product{}, ErrNotFound
	// }

	// return foundProduct, nil
}

func (ps *ProductService) SearchProduct(query float64) (products []domain.Product, err error) {
	return ps.serviceImpl.SearchProduct(query)
	// var filteredProduct []domain.Product
	// for _, product := range rp.Products {
	// 	if query != 0 && product.Price > query {
	// 		filteredProduct = append(filteredProduct, product)
	// 	}
	// }
	// return filteredProduct, nil
}

func (ps *ProductService) AddProduct(product domain.Product) (newProduct domain.Product, err error) {
	if ps.serviceImpl.ExistProduct(product.Name) {
		return domain.Product{}, msg.ErrItemExist
	}

	if ps.serviceImpl.UniqueCodeValue(product.CodeValue) {
		return domain.Product{}, msg.ErrCodeValueRepeat
	}

	expDate, err := parseDate(product.Expiration)
	if err != nil {
		return domain.Product{}, msg.ErrParser
	}

	if validDate(expDate) {
		return domain.Product{}, msg.ErrDateExp
	}
	return ps.serviceImpl.AddProduct(product)
	// if existProduct(product.Name, ps.Products) {
	// 	return domain.Product{}, ErrItemExist
	// }

	// if uniqueCodeValue(product.CodeValue, ps.Products) {
	// 	return domain.Product{}, ErrCodeValueRepeat
	// }

	// expDate, err := parseDate(product.Expiration)
	// if err != nil {
	// 	return domain.Product{}, err
	// }

	// if validDate(expDate) {
	// 	return domain.Product{}, ErrDateExp
	// }

	// lastID := ps.Products[len(ps.Products)-1].Id
	// product.Id = lastID + 1

	// ps.Products = append(ps.Products, product)

	// return product, nil
}

func (ps *ProductService) UpdateProduct(id int, product domain.Product) (updatedProduct domain.Product, err error) {
	return ps.serviceImpl.UpdateProduct(id, product)
}

func (ps *ProductService) UpdatePatchProduct(id int, product domain.ProductPatch) (updatePatchProduct domain.Product, err error) {
	return ps.serviceImpl.UpdatePatchProduct(id, product)
}

func (ps *ProductService) DeleteProduct(id int) (err error) {
	return ps.serviceImpl.DeleteProduct(id)
}

func NewProductService(rp product.ProductServiceImpl) *ProductService {
	return &ProductService{
		serviceImpl: rp,
	}
}

// func existProduct(pName string, products []domain.Product) bool {
// 	for _, p := range products {
// 		if p.Name == pName {
// 			return true
// 		}
// 	}

// 	return false
// }

// func uniqueCodeValue(pCodeValue string, products []domain.Product) bool {
// 	for _, p := range products {
// 		if p.CodeValue == pCodeValue {
// 			return true
// 		}
// 	}

// 	return false
// }

func validDate(date time.Time) bool {
	t := time.Now()

	return t.After(date)
}

func parseDate(date string) (time.Time, error) {
	parseDate, err := time.Parse("01/02/2006", date)
	if err != nil {
		return parseDate, err
	}
	return parseDate, nil
}
