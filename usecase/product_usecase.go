package usecase

import (
	"products-api/model"
	"products-api/repository"
)

type ProductUseCase struct {
	 repository repository.ProductRepository
}

func NewProductUseCase (repo repository.ProductRepository) ProductUseCase {
	return ProductUseCase{
		repository: repo,
	}
}

func (pu *ProductUseCase) GetProducts() ([]model.Product, error) {
	return pu.repository.GetProducts()
}

func (pu *ProductUseCase) CreateProduct(product model.Product) (model.Product, error) {
	productID, err := pu.repository.CreateProduct(product)
	if err != nil {
		return model.Product{}, err
	}

	product.ID = productID
	return product, nil
}

func (pu *ProductUseCase) GetProductById(id int) (*model.Product, error) {
	product, err := pu.repository.GetProductById(id)
	if err != nil {
		return nil, err
	}

	return product, nil
}