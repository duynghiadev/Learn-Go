# Go: More Types - Pointers, Structs, Arrays, Slices, Ranges, and Maps

## Pointers

Pointers in Go store the memory address of a value. They allow you to directly access and modify the value at that memory location.

```go
var a int = 10
var p *int = &a  // p is a pointer to the memory address of a
fmt.Println(*p)   // Dereferencing the pointer to get the value of a
*p = 20           // Modifying the value of a through the pointer
fmt.Println(a)    // Output: 20
```

## Structs

A struct is a composite data type that groups together variables under a single name. These variables are called fields.

```go
type Person struct {
    Name string
    Age  int
}

var p Person = Person{Name: "Alice", Age: 25}
fmt.Println(p.Name)  // Accessing a field
```

## Struct Fields

Fields in a struct are accessed using the dot `.` operator.

```go
type Rectangle struct {
    Width  float64
    Height float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

r := Rectangle{Width: 5.0, Height: 3.0}
fmt.Println(r.Area())  // Output: 15
```

## Arrays

An array is a fixed-size collection of elements of the same type.

```go
var arr [3]int = [3]int{1, 2, 3}
fmt.Println(arr[0])  // Accessing an element by index
```

## Slices

A slice is a dynamically-sized, flexible view into the elements of an array.

```go
var s []int = []int{1, 2, 3}
s = append(s, 4)  // Adding elements dynamically
fmt.Println(s)    // Output: [1 2 3 4]
```

### Creating a Slice from an Array

```go
arr := [5]int{1, 2, 3, 4, 5}
slice := arr[1:4]  // Slice from index 1 to 3
fmt.Println(slice) // Output: [2 3 4]
```

## Range

The `range` keyword is used to iterate over elements in a collection like arrays, slices, maps, or strings.

```go
nums := []int{1, 2, 3, 4}
for index, value := range nums {
    fmt.Println(index, value)
}
```

## Maps

A map is an unordered collection of key-value pairs.

```go
m := make(map[string]int)
m["one"] = 1
m["two"] = 2
fmt.Println(m["one"])  // Accessing a value by key

// Iterating over a map
for key, value := range m {
    fmt.Println(key, value)
}
```

### Deleting an Element from a Map

```go
delete(m, "one")
fmt.Println(m)
```

---

This overview covers pointers, structs, arrays, slices, ranges, and maps in Go. These fundamental types and constructs are crucial for building more complex programs in Go.
