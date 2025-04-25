package engine

import "trading-engine/pkg/uuid"

func (b *OrderBook) GetOrder(orderId uuid.UUID) *Order {
	return b.orders[orderId]
}

func (b *OrderBook) GetPendingOrders(customerId int64) []*Order {
	return b.GetCustomerOrders(customerId, false)
}

func (b *OrderBook) GetCustomerOrders(customerId int64, all bool) []*Order {
	var orders []*Order

	for _, order := range b.customerOrders[customerId] {
		if all {
			orders = append(orders, order)
		} else if order.Status == StatusOpen || order.Status == StatusPartiallyFilled {
			orders = append(orders, order)
		}
	}

	return orders
}

func (b *OrderBook) GetFills() []*Fill {
	return b.fills
}

func (b *OrderBook) GetOrders() []*Order {
	var orders []*Order
	for _, order := range b.orders {
		orders = append(orders, order)
	}

	return orders
}
