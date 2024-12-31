# **Function Closures in Go**

## **What is a Closure?**

In Go, functions can be **closures**. A closure is a function value that references variables from outside its body.

- The function can **access** and **modify** the referenced variables.
- The function is said to be "bound" to these variables.

Closures are useful for encapsulating functionality along with the state.

---

## **Example: Closure with `adder` Function**

The following example demonstrates how closures work in Go:

```go
package main

import "fmt"

func adder() func(int) int {
    sum := 0
    return func(x int) int {
        sum += x
        return sum
    }
}

func main() {
    pos := adder()
    neg := adder()

    fmt.Println(pos(1))  // Output: 1
    fmt.Println(pos(2))  // Output: 3
    fmt.Println(pos(3))  // Output: 6

    fmt.Println(neg(-1)) // Output: -1
    fmt.Println(neg(-2)) // Output: -3
    fmt.Println(neg(-3)) // Output: -6
}
```

### **Explanation**

1. **`adder` Function**:

   - `adder` returns an anonymous function that references the variable `sum` from its outer scope.
   - Each call to `adder` creates a new instance of `sum`.

2. **Closures Bound to `sum`**:
   - The `pos` and `neg` closures each have their own separate `sum` variables.
   - These variables are updated whenever the closure is called.

---

## **Key Characteristics of Closures in Go**

- **Encapsulation**: Closures encapsulate the state (`sum` in the example above) along with the function logic.
- **Independent State**: Each closure instance maintains its own independent copy of the referenced variables.
- **Reusability**: Closures can be reused with different states, making them versatile for tasks like counters or accumulators.

---

## **Use Cases for Closures**

1. **Counters**: Creating functions that maintain and update counts.
2. **Event Handlers**: Passing stateful functions as callbacks.
3. **Custom Iterators**: Generating functions that produce values in a sequence.

---

## **Summary**

- Closures in Go allow functions to retain access to variables from their enclosing scope.
- Each closure instance is bound to its own copy of the referenced variables.
- They are powerful tools for maintaining state and creating reusable, stateful functions.
