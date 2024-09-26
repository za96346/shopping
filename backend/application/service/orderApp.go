package service

import (
	"backend/domain/entities"
	"backend/domain/aggregate"
	"backend/domain/repository"
)

type OrderApp struct {
	OrderRepository repository.OrderRepository
	OrderItemRepository repository.OrderItemRepository
	ProductRepository repository.ProductRepository
}

type OrderAppInterface interface {
	GetOrders(orderEntity *entities.Order) (*[]entities.Order, error)
	AddOrder(orderEntity *entities.Order) (*entities.Order, error)
}

func (app *OrderApp) GetOrders(orderEntity *entities.Order) (*[]entities.Order, error) {
	return (*app).OrderRepository.GetOrders(orderEntity)
}

func (app *OrderApp) AddOrder(orderItems *[]entities.OrderItem) (*entities.Order, error) {
	orderAggregateRoot := aggregate.OrderAggregateRoot{}

	// 查找訂單最大id
	orderAggregateRoot.Order.ID = (*app).OrderRepository.GetNewOrderId()

	for _, orderItem := range *orderItems {
		orderAggregateRoot.AddOrderItem(&orderItem)

		// 商品檢驗
		_, err := (*app).ProductRepository.GetProduct(&entities.Product{
			ID: orderItem.ProductId,
		})
		if err != nil {
			return nil, err
		}

		(*app).OrderItemRepository.AddOrderItem(&orderItem)
	}

	return (*app).OrderRepository.AddOrder(orderAggregateRoot.Order)
} 
