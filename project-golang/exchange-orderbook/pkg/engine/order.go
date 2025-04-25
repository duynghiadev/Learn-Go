package engine

import (
	"time"
	"trading-engine/pkg/uuid"
)

type Side string
type TimeInForce string
type OrderStatus string

const (
	Bid Side = "BUY"
	Ask Side = "SELL"

	GTC TimeInForce = "GTC"
	GTT TimeInForce = "GTT"

	StatusOpen            OrderStatus = "Open"
	StatusPartiallyFilled OrderStatus = "PartiallyFilled"
	StatusFilled          OrderStatus = "Filled"
	StatusCancelled       OrderStatus = "Cancelled"
	StatusExpired         OrderStatus = "Expired"
)

type Order struct {
	OrderId    uuid.UUID `json:"orderId"`
	ProductId  string    `json:"productId"`
	CustomerId int64     `json:"customerId"`

	Side   Side        `json:"side"`
	Size   float64     `json:"size"`
	Filled float64     `json:"filled"`
	Price  float64     `json:"price"`
	Status OrderStatus `json:"status"`

	TimeInForce TimeInForce `json:"timeInForce"`
	CreatedAt   int64       `json:"createdAt"`
	ExpiredAt   *int64      `json:"expiredAt"`

	Fills []*Fill `json:"fills"`
}

func newOrder(
	productId string,
	customerId int64,
	size float64,
	price float64,
	side Side,
	timeInForce TimeInForce,
	expiredAt *int64,
) *Order {
	return &Order{
		OrderId:     uuid.New(),
		ProductId:   productId,
		CustomerId:  customerId,
		Status:      StatusOpen,
		Size:        size,
		Price:       price,
		CreatedAt:   time.Now().Unix(),
		Side:        side,
		TimeInForce: timeInForce,
		ExpiredAt:   expiredAt,
	}
}

func (o *Order) validate() bool {
	if o.Size <= 0 || o.Price <= 0 {
		return false
	}

	if o.TimeInForce == GTT {
		return o.ExpiredAt != nil && o.CreatedAt < *o.ExpiredAt
	}

	return true
}

func (o *Order) fill(sizeFilled float64) {
	o.Filled += sizeFilled
	if o.Filled == o.Size {
		o.Status = StatusFilled
		return
	} else if o.Filled > 0 {
		o.Status = StatusPartiallyFilled
		return
	}

}

func (o *Order) isCompletelyFilled() bool {
	return o.Filled == o.Size
}

func (o *Order) cancel() {
	o.Status = StatusCancelled
}
