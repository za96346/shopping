package repository

import "backend/domain/entities"

type CategoryRepository interface {
	GetCategories(categoryEntity *entities.Category) (*[]entities.Category, error)
}
