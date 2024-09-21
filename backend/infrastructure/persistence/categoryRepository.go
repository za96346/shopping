package persistence

import (
	"backend/domain/entities"
	"backend/domain/repository"

	"gorm.io/gorm"
)

type CategoryRepo struct {
	db        *gorm.DB
	tableName string
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepo {
	return &CategoryRepo{db, "category"}
}

var _ repository.CategoryRepository = &CategoryRepo{}

func (r *CategoryRepo) GetCategories(categoryEntity *entities.Category) (*[]entities.Category, error) {
	var categories []entities.Category

	err := r.db.
		Debug().
		Table(r.tableName).
		Where("delete_flag = 0").
		Order("sort asc").
		Find(&categories).
		Error

	return &categories, err
}
