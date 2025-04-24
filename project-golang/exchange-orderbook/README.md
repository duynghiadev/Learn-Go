# [https://github.com/anhthii/exchange-orderbook](https://github.com/anhthii/exchange-orderbook)

# Simple OrderBook Implementation

- [Simple OrderBook Implementation](#simple-orderbook-implementation)
  - [Run program](#run-program)
  - [Run test](#run-test)
  - [Result](#result)
  - [Technical decisions](#technical-decisions)
  - [Limitation](#limitation)

## Run program

```
go run *.go
```

## Run test

```
go test ./... -v
```

## Result

- Initial state

```
================ ASK ================
Price: 119.000000, Size: 100.000000
Price: 115.000000, Size: 100.000000
Price: 114.000000, Size: 100.000000
Price: 110.000000, Size: 200.000000
Price: 109.000000, Size: 100.000000
=============== BID ===============
Price: 100.000000, Size: 200.000000
Price: 99.000000, Size: 200.000000
Price: 97.000000, Size: 100.000000
Price: 90.000000, Size: 100.000000
=============== FILLS ===============
====================================

```

- Create an order at $150.0

```
================ ASK ================
Price: 150.000000, Size: 100.000000
Price: 119.000000, Size: 100.000000
Price: 115.000000, Size: 100.000000
Price: 114.000000, Size: 100.000000
Price: 110.000000, Size: 200.000000
Price: 109.000000, Size: 100.000000
=============== BID ===============
Price: 100.000000, Size: 200.000000
Price: 99.000000, Size: 200.000000
Price: 97.000000, Size: 100.000000
Price: 90.000000, Size: 100.000000
=============== FILLS ===============
====================================
```

- Create an order at $120.0 with size = 50

```
================ ASK ================
Price: 150.000000, Size: 100.000000
Price: 119.000000, Size: 100.000000
Price: 115.000000, Size: 100.000000
Price: 114.000000, Size: 100.000000
Price: 110.000000, Size: 200.000000
Price: 109.000000, Size: 50.000000
=============== BID ===============
Price: 100.000000, Size: 200.000000
Price: 99.000000, Size: 200.000000
Price: 97.000000, Size: 100.000000
Price: 90.000000, Size: 100.000000
=============== FILLS ===============
Price: 109.000000, Size: 50.000000
====================================
```

- Create an order at $99.0 with size = 100

```
================ ASK ================
Price: 150.000000, Size: 100.000000
Price: 119.000000, Size: 100.000000
Price: 115.000000, Size: 100.000000
Price: 114.000000, Size: 100.000000
Price: 110.000000, Size: 200.000000
Price: 109.000000, Size: 50.000000
=============== BID ===============
Price: 100.000000, Size: 100.000000
Price: 99.000000, Size: 200.000000
Price: 97.000000, Size: 100.000000
Price: 90.000000, Size: 100.000000
=============== FILLS ===============
Price: 109.000000, Size: 50.000000
Price: 100.000000, Size: 100.000000
====================================
```

## Technical decisions

- Data structures:
  - Linked list: for storing orders in a limit. This allows quick order insertion with 0(1) complexity.
  - Arrays: for storing limits. I know this is not really efficient but it has built in sorting function so I can leverage it to save development time.
  - Maps: for storing pointers to orders which allow fast access when retrieve info of an order.
- Matching algorithm : First In First Out (FIFO). This algorithm is simple and quite straight forward so I used it.

## Limitation

- Need to sort limits every time the engine needs to match an order.
- Scheduler to expire GTT orders hasn't been implemented
- Code coverage is quite low due to time constraint
- Some part of the code is not super efficient. Limits could be stored in a more efficient data structure like binary tree or red-black tree.
