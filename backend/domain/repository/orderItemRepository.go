package repository

import "backend/domain/entities"

type OrderItemRepository interface {
	AddOrderItem(orderItem *entities.OrderItem) (*entities.OrderItem, error)
}
