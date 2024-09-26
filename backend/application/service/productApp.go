package service

import (
	"backend/domain/entities"
	"backend/domain/repository"
)

type ProductApp struct {
	ProductRepository repository.ProductRepository
}

var _ ProductAppInterface = &ProductApp{}

type ProductAppInterface interface {
	GetProducts(productEntity *entities.Product) (*[]entities.Product, error)
}

func (app *ProductApp) GetProducts(productEntity *entities.Product) (*[]entities.Product, error) {
	return (*app).ProductRepository.GetProducts(productEntity)
}
