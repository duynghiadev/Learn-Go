Based on the provided context, I can explain your project which is a Simple OrderBook Implementation. Here's a detailed breakdown:

### Project Overview

This is an order book implementation for a trading system, which manages and matches buy (bid) and sell (ask) orders for financial instruments.

### Key Features

1. Order Management

   - Maintains two sides of orders:
     - ASK (sell orders)
     - BID (buy orders)
   - Displays order book state with prices and sizes
   - Tracks filled orders

2. Order Matching

   - Uses First In First Out (FIFO) matching algorithm
   - Matches orders based on price and size
   - Records fills when orders are matched

### Technical Implementation

1. Data Structures Used

   - Linked List : Used for storing orders within a limit for O(1) insertion complexity
   - Arrays : Used for storing limits, leveraging built-in sorting functionality
   - Maps : Used for storing order pointers for fast access to order information

2. Example Operations The README shows three main operations:

- Creating an order at $150.0
- Creating an order at $120.0 with size = 50
- Creating an order at $99.0 with size = 100

### Current Limitations

1. Performance Issues

   - Needs to sort limits on every order match
   - Some inefficient code sections where better data structures could be used

2. Missing Features

   - GTT (Good Till Time) orders expiration scheduler not implemented
   - Low code coverage due to time constraints

3. Potential Improvements

   - Could use more efficient data structures like binary trees or red-black trees for storing limits
   - Could implement order expiration functionality
   - Could improve test coverage

### How to Run

1. Running the Program

```bash
go run *.go
```

2. Running Tests

```bash
go test ./... -v
```

This implementation provides a foundation for a trading system's order book, with room for optimization and feature additions while maintaining core functionality for order matching and management.
