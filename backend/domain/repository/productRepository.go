package repository

import "backend/domain/entities"

type ProductRepository interface {
	GetProducts(productEntity *entities.Product) (*[]entities.Product, error)
}
