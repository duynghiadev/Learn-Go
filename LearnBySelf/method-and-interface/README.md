# Methods and Interfaces in Go

Go provides powerful features for object-oriented programming through **methods** and **interfaces**. Understanding these concepts is crucial for writing robust and idiomatic Go code.

## Methods

A **method** is a function with a special receiver argument. The receiver appears in its own argument list between the `func` keyword and the method name. This receiver can be a value or a pointer.

### Syntax

```go
func (receiverType ReceiverName) MethodName(parameters) returnType {
    // Method body
}
```

### Example

```go
type Rectangle struct {
    Width, Height float64
}

// Method with value receiver
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

// Method with pointer receiver
func (r *Rectangle) Scale(factor float64) {
    r.Width *= factor
    r.Height *= factor
}

func main() {
    rect := Rectangle{Width: 10, Height: 5}
    fmt.Println("Area:", rect.Area())
    rect.Scale(2)
    fmt.Println("Scaled Area:", rect.Area())
}
```

### Key Points

- Use **value receivers** if the method doesn’t modify the receiver.
- Use **pointer receivers** to modify the receiver or avoid copying large structs.

---

## Interfaces

An **interface** is a type that specifies a set of method signatures. A value satisfies an interface if it implements those methods.

### Syntax

```go
type InterfaceName interface {
    MethodName1(parameters) returnType1
    MethodName2(parameters) returnType2
}
```

### Example

```go
type Shape interface {
    Area() float64
    Perimeter() float64
}

type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
    return 2 * math.Pi * c.Radius
}

type Rectangle struct {
    Width, Height float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
    return 2 * (r.Width + r.Height)
}

func main() {
    var s Shape

    s = Circle{Radius: 5}
    fmt.Println("Circle Area:", s.Area())
    fmt.Println("Circle Perimeter:", s.Perimeter())

    s = Rectangle{Width: 10, Height: 5}
    fmt.Println("Rectangle Area:", s.Area())
    fmt.Println("Rectangle Perimeter:", s.Perimeter())
}
```

### Key Points

- Interfaces are **implicitly satisfied** in Go.
- Interfaces can be composed using embedding:
  ```go
  type ReadWrite interface {
      Reader
      Writer
  }
  ```

---

## Empty Interface

The **empty interface** `interface{}` can hold any value since every type satisfies it.

### Example

```go
func PrintAnything(value interface{}) {
    fmt.Println(value)
}

func main() {
    PrintAnything(42)
    PrintAnything("Hello")
    PrintAnything([]int{1, 2, 3})
}
```

### Use Cases

- Generic functions.
- Handling unknown types.

---

## Type Assertions

Type assertions allow you to extract the dynamic value of an interface.

### Syntax

```go
value, ok := interfaceValue.(ConcreteType)
```

### Example

```go
func Describe(i interface{}) {
    switch v := i.(type) {
    case string:
        fmt.Println("String:", v)
    case int:
        fmt.Println("Integer:", v)
    default:
        fmt.Println("Unknown type")
    }
}

func main() {
    Describe("Hello")
    Describe(42)
    Describe(3.14)
}
```

---

## Practical Example

Combining methods and interfaces for a real-world scenario:

### Code

```go
type Logger interface {
    Log(message string)
}

type ConsoleLogger struct{}

func (c ConsoleLogger) Log(message string) {
    fmt.Println("Log:", message)
}

type FileLogger struct {
    FileName string
}

func (f FileLogger) Log(message string) {
    // Simulate writing to a file
    fmt.Printf("Writing to %s: %s\n", f.FileName, message)
}

func main() {
    var logger Logger

    logger = ConsoleLogger{}
    logger.Log("This is a console log.")

    logger = FileLogger{FileName: "app.log"}
    logger.Log("This is a file log.")
}
```

---

By mastering methods and interfaces, you can write flexible and reusable code in Go. Let me know if you need further explanations or examples!
