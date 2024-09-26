package persistence

import (
	"backend/domain/entities"
	"backend/domain/repository"

	"gorm.io/gorm"
)

type OrderRepo struct {
	db        *gorm.DB
	tableName string
}

func NewOrderRepository(db *gorm.DB) *OrderRepo {
	return &OrderRepo{db, "order"}
}

var _ repository.OrderRepository = &OrderRepo{}

func (r *OrderRepo) GetOrders(orderEntity *entities.Order) (*[]entities.Order, error) {
	var orders []entities.Order

	searchQuery := r.db.
		Debug().
		Table(r.tableName).
		Where("delete_flag = 0").
		Where("order_member = ?", orderEntity.OrderMember).
		Order("sort asc")

	if orderEntity.CreateTime != nil {
		searchQuery = searchQuery.Where("create_time = ?", orderEntity.CreateTime)
	}

	err := searchQuery.Find(&orders).Error
	return &orders, err
}

func (r *OrderRepo) AddOrder(orderEntity *entities.Order) (*entities.Order, error) {
	err := r.db.
		Debug().
		Table(r.tableName).
		Create(&orderEntity).
		Error

	return orderEntity, err
}

func (r *OrderRepo) GetNewOrderId() int {
	var MaxCount int64

	r.db.
		Debug().
		Table(r.tableName).
		Select("max(id)").
		Row().
		Scan(&MaxCount)

	return int(MaxCount) + 1
}