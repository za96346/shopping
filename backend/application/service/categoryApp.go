package service

import (
	"backend/domain/entities"
	"backend/domain/repository"
)

type CategoryApp struct {
	CategoryRepository repository.CategoryRepository
}

var _ CategoryAppInterface = &CategoryApp{}

type CategoryAppInterface interface {
	GetCategories(categoryEntity *entities.Category) (*[]entities.Category, error)
}

func (app *CategoryApp) GetCategories(categoryEntity *entities.Category) (*[]entities.Category, error) {
	return (*app).CategoryRepository.GetCategories(categoryEntity)
}
