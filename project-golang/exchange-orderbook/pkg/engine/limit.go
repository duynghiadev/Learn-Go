package engine

import (
	"math"
)

type OrderPtr struct {
	Order *Order
	Next  *OrderPtr
}

type Limit struct {
	Price         float64
	AvailableSize float64

	Head *OrderPtr
	Tail *OrderPtr
}

func NewLimit(price float64) *Limit {
	return &Limit{
		Price: price,
	}
}

func (l *Limit) AddOrder(order *Order) {
	if l.Head == nil {
		l.Head = &OrderPtr{
			Order: order,
		}
		l.Tail = l.Head
	} else {
		l.Tail.Next = &OrderPtr{
			Order: order,
		}
	}

	l.AvailableSize += order.Size

}

func (l *Limit) RemoveOrder(order *Order) {
	for curr, prev := l.Head, l.Head; curr != nil; curr = curr.Next {
		if curr.Order.OrderId == order.OrderId {
			if curr == l.Head {
				l.Head = curr.Next
			} else if curr == l.Tail {
				l.Tail = nil
			} else {
				prev.Next = curr.Next

			}

		}
		prev = curr

	}

	remain := order.Size - order.Filled
	l.AvailableSize -= remain
}

func (l *Limit) FillOrder(order *Order) (totalFilled float64, fills []*Fill) {
	remain := order.Size - order.Filled
	for ptr := l.Head; ptr != nil && remain > 0; ptr = ptr.Next {
		matchedOrder := ptr.Order

		if matchedOrder.CustomerId == order.CustomerId {
			continue
		}

		available := matchedOrder.Size - matchedOrder.Filled
		filledSize := math.Min(remain, available)
		remain -= filledSize
		totalFilled += filledSize

		matchedOrder.fill(filledSize)

		filledPrice := math.Max(order.Price, matchedOrder.Price)
		if order.Side == Bid {
			filledPrice = math.Min(order.Price, matchedOrder.Price)
		}

		if matchedOrder.isCompletelyFilled() {
			l.RemoveOrder(matchedOrder)
		}

		l.AvailableSize -= filledSize
		fills = append(fills, NewFill(
			order.OrderId,
			matchedOrder.OrderId, filledPrice, filledSize))

	}

	return
}

func (l *Limit) GetDepth() (depth int) {
	for order := l.Head; order != nil; order = order.Next {
		depth++
	}
	return
}
