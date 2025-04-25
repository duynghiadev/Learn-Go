package main

import "trading-engine/pkg/engine"

func main() {

	orderbook := engine.NewOrderBook()

	productId := "Theory of Cryptography"
	orderbook.AddProduct(productId)

	size := 100.0
	// Create ASK wall
	orderbook.CreateLimitOrder(productId, randUser(), engine.Ask, size, 115.0)
	orderbook.CreateLimitOrder(productId, randUser(), engine.Ask, size, 110.0)
	orderbook.CreateLimitOrder(productId, randUser(), engine.Ask, size, 109.0)
	orderbook.CreateLimitOrder(productId, randUser(), engine.Ask, size, 110.0)
	orderbook.CreateLimitOrder(productId, randUser(), engine.Ask, size, 114.0)
	orderbook.CreateLimitOrder(productId, randUser(), engine.Ask, size, 119.0)

	// Create BID wall
	orderbook.CreateLimitOrder(productId, randUser(), engine.Bid, size, 100.0)
	orderbook.CreateLimitOrder(productId, randUser(), engine.Bid, size, 97.0)
	orderbook.CreateLimitOrder(productId, randUser(), engine.Bid, size, 90.0)
	orderbook.CreateLimitOrder(productId, randUser(), engine.Bid, size, 99.0)
	orderbook.CreateLimitOrder(productId, randUser(), engine.Bid, size, 99.0)
	orderbook.CreateLimitOrder(productId, randUser(), engine.Bid, size, 100.0)

	orderbook.PrintOrderBook()

	// create an order at $150.0
	orderbook.CreateLimitOrder(productId, randUser(), engine.Ask, size, 150.0)
	orderbook.PrintOrderBook()

	// create an order at $120.0 with size = 50
	orderbook.CreateLimitOrder(productId, randUser(), engine.Bid, 50.0, 120.0)
	orderbook.PrintOrderBook()

	// create an order at $99.0 with size = 100
	orderbook.CreateLimitOrder(productId, randUser(), engine.Ask, 100.0, 99.0)
	orderbook.PrintOrderBook()

}
