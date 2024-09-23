package repository

import "backend/domain/entities"

type OrderRepository interface {
	GetOrders(orderEntity *entities.Order) (*[]entities.Order, error)
	AddOrder(orderEntity *entities.Order) (*entities.Order, error)
}
