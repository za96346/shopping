package persistence

import (
	"backend/domain/entities"
	"backend/domain/repository"

	"gorm.io/gorm"
)

type ProductRepo struct {
	db        *gorm.DB
	tableName string
}

func NewProductRepository(db *gorm.DB) *ProductRepo {
	return &ProductRepo{db, "product"}
}

var _ repository.ProductRepository = &ProductRepo{}

func (r *ProductRepo) GetProducts(productEntity *entities.Product) (*[]entities.Product, error) {
	var products []entities.Product

	searchQuery := r.db.
		Debug().
		Table(r.tableName).
		Where("delete_flag = 0").
		Order("sort asc")

	// product name
	if productEntity.Name != "" {
		searchQuery = searchQuery.Where("name like ?", "%"+productEntity.Name+"%")
	}

	err := searchQuery.Find(&products).Error
	return &products, err
}

func (r *ProductRepo) GetProduct(productEntity *entities.Product) (*entities.Product, error) {
	var product entities.Product

	searchQuery := r.db.
		Debug().
		Table(r.tableName)
	
	if productEntity.ID != 0 {
		searchQuery = searchQuery.Where("id = ?", productEntity.ID)
	}

	err := searchQuery.First(&product).Error
	return &product, err
}