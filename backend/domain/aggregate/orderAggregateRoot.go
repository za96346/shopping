package aggregate

import "backend/domain/entities"

type OrderAggregateRoot struct {
	Order      *entities.Order
	OrderItems *[]entities.OrderItem
	Member     *entities.Member
}

// 新增商品
func (o *OrderAggregateRoot) AddOrderItem(orderItem *entities.OrderItem) {
	*(*o).OrderItems = append(*(*o).OrderItems, *orderItem)
	(*o).Order.TotalAmount += orderItem.SinglePrice * orderItem.Quantity
}

// 根據產品 ID 移除商品
func (o *OrderAggregateRoot) RemoveOrderItem(productId int) {
	newOrderItems := []entities.OrderItem{}

	for _, item := range *o.OrderItems {
		if item.ProductId != productId {
			newOrderItems = append(newOrderItems, item)
		} else {
			o.Order.TotalAmount -= item.SinglePrice * item.Quantity
		}
	}
	*o.OrderItems = newOrderItems
}
