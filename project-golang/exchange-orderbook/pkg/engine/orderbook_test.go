package engine

import (
	"testing"
	"trading-engine/pkg/assert"
)

type orderParams struct {
	product     string
	customerId  int64
	size        float64
	price       float64
	side        Side
	timeInForce TimeInForce
	expiredAt   *int64
}

func TestAddProduct(t *testing.T) {
	orderbook := NewOrderBook()

	product := "BTC-USD"
	orderbook.AddProduct(product)

	_, exist := orderbook.products[product]
	assert.Equal(t, exist, true)

}

func TestOrderBook(t *testing.T) {
	orderbook := NewOrderBook()

	params := []orderParams{{
		product:     "BTC-USD",
		customerId:  1,
		size:        100.0,
		price:       100.0,
		side:        Bid,
		timeInForce: GTC,
		expiredAt:   nil,
	}, {
		product:     "BTC-USD",
		customerId:  2,
		size:        200.0,
		price:       100.0,
		side:        Ask,
		timeInForce: GTC,
		expiredAt:   nil,
	}, {
		product:     "BTC-USD",
		customerId:  2,
		size:        100.0,
		price:       90.0,
		side:        Ask,
		timeInForce: GTC,
		expiredAt:   nil,
	}, {
		product:     "BTC-USD",
		customerId:  1,
		size:        300.0,
		price:       110.0,
		side:        Bid,
		timeInForce: GTC,
		expiredAt:   nil,
	},
	}

	orderbook.AddProduct(params[0].product)

	var createdOrders []*Order

	for _, param := range params {
		created, err := orderbook.CreateOrder(param.product, param.customerId, param.side, param.size, param.price, param.timeInForce, param.expiredAt)
		assert.Equal(t, err, nil)
		createdOrders = append(createdOrders, created)
	}

	orders := orderbook.GetOrders()
	assert.Equal(t, len(orders), 4)

	fills := orderbook.GetFills()
	assert.Equal(t, len(fills), 3)

	ordersOfCustomer1 := orderbook.GetCustomerOrders(1, true)
	assert.Equal(t, len(ordersOfCustomer1), 2)
	ordersOfCustomer2 := orderbook.GetCustomerOrders(2, true)
	assert.Equal(t, len(ordersOfCustomer2), 2)

	pendings1 := orderbook.GetPendingOrders(1)
	pendings2 := orderbook.GetPendingOrders(2)
	assert.Equal(t, len(pendings1), 1)
	assert.Equal(t, len(pendings2), 0)

	assert.Equal(t, createdOrders[0].Status, StatusFilled)
	assert.Equal(t, createdOrders[1].Status, StatusFilled)
	assert.Equal(t, createdOrders[2].Status, StatusFilled)
	assert.Equal(t, createdOrders[3].Status, StatusPartiallyFilled)

	assert.Equal(t, createdOrders[0].Filled, params[0].size)
	assert.Equal(t, createdOrders[1].Filled, params[1].size)
	assert.Equal(t, createdOrders[2].Filled, params[2].size)
	assert.Equal(t, createdOrders[3].Filled, 200.0)

	// Testing cancel order
	orderbook.CancelOrder(createdOrders[3].OrderId)
	updatedOrder := orderbook.GetOrder(createdOrders[3].OrderId)
	assert.Equal(t, updatedOrder.Status, StatusCancelled)

	// If user1 create a new order, it will not be matched to the previous cancelled order
	newOrderParam := orderParams{
		product:     "BTC-USD",
		customerId:  2,
		size:        100.0,
		price:       110.0,
		side:        Ask,
		timeInForce: GTC,
		expiredAt:   nil,
	}

	created, err := orderbook.CreateOrder(
		newOrderParam.product,
		newOrderParam.customerId,
		newOrderParam.side,
		newOrderParam.size,
		newOrderParam.price,
		newOrderParam.timeInForce,
		newOrderParam.expiredAt,
	)

	assert.Equal(t, err, nil)
	assert.Equal(t, created.Status, StatusOpen)
	assert.Equal(t, created.Filled, 0.0)
}

func TestSortLimits(t *testing.T) {
	// generate a list of mocks limits
	bidLimits := []*Limit{
		{Price: 100.0},
		{Price: 90.0},
		{Price: 200.0},
		{Price: 300.0},
		{Price: 50.0},
		{Price: 120.0},
	}

	askLimits := make([]*Limit, len(bidLimits))
	copy(askLimits, bidLimits)

	orderbook := NewOrderBook()
	orderbook.sortLimitsByPrice(bidLimits, Bid)
	orderbook.sortLimitsByPrice(askLimits, Ask)

	assert.Equal(t, bidLimits[0].Price, 300.0)
	assert.Equal(t, bidLimits[len(bidLimits)-1].Price, 50.0)

	assert.Equal(t, askLimits[0].Price, 50.0)
	assert.Equal(t, askLimits[len(askLimits)-1].Price, 300.0)

}
