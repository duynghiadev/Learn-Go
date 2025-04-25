package engine

import (
	"errors"
	"fmt"
	"sort"
	"trading-engine/pkg/uuid"
)

var (
	ErrProductNotFound = errors.New("Product not found")
	ErrInvalidOrder    = errors.New("Invalid order")
	ErrOrderNotFound   = errors.New("Order not found")
)

type OrderBook struct {
	products     map[string]bool
	productCount int
	// cache order for fast access
	orders         map[uuid.UUID]*Order
	customerOrders map[int64][]*Order

	askLimits map[float64]*Limit
	bidLimits map[float64]*Limit

	asks  []*Limit
	bids  []*Limit
	fills []*Fill
}

func NewOrderBook() *OrderBook {
	return &OrderBook{
		products:       make(map[string]bool),
		asks:           make([]*Limit, 0, 100),
		bids:           make([]*Limit, 0, 100),
		askLimits:      make(map[float64]*Limit),
		bidLimits:      make(map[float64]*Limit),
		orders:         make(map[uuid.UUID]*Order),
		customerOrders: make(map[int64][]*Order),
		fills:          make([]*Fill, 0, 100),
	}
}

func (b *OrderBook) AddProduct(productId string) {
	if b.productCount >= 1 {
		// :(( we only support 1 product for now due to time constraint and for the sake of simplicity
		panic("Only support 1 product")
	}

	b.products[productId] = true
	b.productCount += 1
}

func (b *OrderBook) CreateOrder(
	productId string,
	customerId int64,
	side Side,
	size float64,
	price float64,
	timeInForce TimeInForce,
	expiredAt *int64,
) (*Order, error) {
	if _, ok := b.products[productId]; !ok {
		return nil, ErrProductNotFound
	}
	order := newOrder(productId, customerId, size, price, side, timeInForce, expiredAt)
	b.orders[order.OrderId] = order
	b.customerOrders[customerId] = append(b.customerOrders[customerId], order)

	valid := order.validate()
	if !valid {
		return nil, ErrInvalidOrder
	}

	fills := b.fillOrder(order, side)
	b.fills = append(b.fills, fills...)
	return order, nil
}

// CreateLimitOrder is a helper function to create limit order with default timeInForce = GTC
func (b *OrderBook) CreateLimitOrder(
	productId string,
	customerId int64,
	side Side,
	size float64,
	price float64,
) (*Order, error) {
	return b.CreateOrder(productId, customerId, side, size, price, GTC, nil)
}

func (b *OrderBook) CancelOrder(orderId uuid.UUID) error {
	order, exist := b.orders[orderId]
	if !exist {
		return ErrInvalidOrder
	}

	limit := b.askLimits[order.Price]
	if order.Side == Bid {
		limit = b.bidLimits[order.Price]
	}

	if limit != nil {
		limit.RemoveOrder(order)
	}

	order.cancel()
	return nil
}

func (b *OrderBook) getMatchLimits(price float64, size float64, side Side) []*Limit {
	// TODO: Could potentially use a binary tree to improve the performance if the number of limits is large
	var limits []*Limit
	var availableSize float64

	if side == Bid {
		b.sortLimitsByPrice(b.asks, Ask)
		for _, limit := range b.asks {
			if limit.Price <= price {
				limits = append(limits, limit)
				availableSize += limit.AvailableSize

			}

			if availableSize >= size {
				break
			}
		}
	} else {
		b.sortLimitsByPrice(b.bids, Bid)
		for _, limit := range b.bids {
			if limit.Price >= price {
				limits = append(limits, limit)
				availableSize += limit.AvailableSize
			}
			if availableSize >= size {
				break
			}
		}
	}

	return limits
}

func (b *OrderBook) addOrderToLimits(order *Order) {
	limits := b.bidLimits
	if order.Side == Ask {
		limits = b.askLimits
	}

	limit, exist := limits[order.Price]
	if !exist {
		l := NewLimit(order.Price)
		limits[order.Price] = l
		l.AddOrder(order)

		if order.Side == Bid {
			b.bids = append(b.bids, l)
		} else {
			b.asks = append(b.asks, l)
		}
	} else {
		limit.AddOrder(order)
	}

}

func (b *OrderBook) fillOrder(order *Order, side Side) (orderFills []*Fill) {
	matchLimits := b.getMatchLimits(order.Price, order.Size, order.Side)

	for _, matchLimit := range matchLimits {
		sizeFilled, fills := matchLimit.FillOrder(order)
		order.fill(sizeFilled)
		orderFills = append(orderFills, fills...)
		if order.isCompletelyFilled() {
			break
		}

	}

	if !order.isCompletelyFilled() {
		b.addOrderToLimits(order)
	}

	return
}

func (b *OrderBook) sortLimitsByPrice(limits []*Limit, side Side) {
	sort.Slice(limits, func(i, j int) bool {
		if side == Bid {
			return limits[i].Price > limits[j].Price
		} else {

			return limits[i].Price < limits[j].Price
		}
	})
}

func (b *OrderBook) PrintOrderBook() {
	b.PrintAskWall()
	b.PrintBidWall()
	b.PrintFills()
	fmt.Printf("\n")
}

func (b *OrderBook) PrintAskWall() {
	fmt.Println("================ ASK ================")
	// need to inverse the side to display the Limits from highest to lowest
	inversedSize := Bid
	b.sortLimitsByPrice(b.asks, inversedSize)
	for _, limit := range b.asks {
		fmt.Printf("Price: %f, Size: %f\n", limit.Price, limit.AvailableSize)
	}
}

func (b *OrderBook) PrintBidWall() {
	fmt.Println("=============== BID ===============")
	b.sortLimitsByPrice(b.bids, Bid)
	for _, limit := range b.bids {
		fmt.Printf("Price: %f, Size: %f\n", limit.Price, limit.AvailableSize)
	}
}

func (b *OrderBook) PrintFills() {
	fmt.Println("=============== FILLS ===============")
	for _, fill := range b.fills {
		fmt.Printf("Price: %f, Size: %f\n", fill.Price, fill.Size)
	}
	fmt.Println("====================================")
}
