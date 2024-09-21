package service

import (
	"backend/domain/entities"
	"backend/domain/repository"
)

type OrderApp struct {
	OrderRepository repository.OrderRepository
}

type OrderAppInterface interface {
	GetOrders(orderEntity *entities.Order) (*[]entities.Order, error)
}

func (app *OrderApp) GetOrders(orderEntity *entities.Order) (*[]entities.Order, error) {
	return (*app).OrderRepository.GetOrders(orderEntity)
}