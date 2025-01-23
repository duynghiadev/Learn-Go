package main

import "fmt"

// List represents a singly-linked list that holds values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

// NewList creates a new empty list
func NewList[T any]() *List[T] {
	return nil
}

// Insert adds a new value at the beginning of the list
func (l *List[T]) Insert(val T) *List[T] {
	return &List[T]{
		next: l,
		val:  val,
	}
}

// Length returns the number of elements in the list
func (l *List[T]) Length() int {
	count := 0
	current := l
	for current != nil {
		count++
		current = current.next
	}
	return count
}

// Print displays all elements in the list
func (l *List[T]) Print() {
	current := l
	fmt.Print("List: [")
	for current != nil {
		fmt.Printf("%v", current.val)
		if current.next != nil {
			fmt.Print(" -> ")
		}
		current = current.next
	}
	fmt.Println("]")
}

func main() {
	// Example with integers
	intList := NewList[int]()
	intList = intList.Insert(3)
	intList = intList.Insert(2)
	intList = intList.Insert(1)

	fmt.Printf("Integer ")
	intList.Print()
	fmt.Printf("List length: %d\n\n", intList.Length())

	// Example with strings
	stringList := NewList[string]()
	stringList = stringList.Insert("world")
	stringList = stringList.Insert("hello")

	fmt.Printf("String ")
	stringList.Print()
	fmt.Printf("List length: %d\n\n", stringList.Length())

	// Example with custom type
	type Person struct {
		name string
		age  int
	}

	personList := NewList[Person]()
	personList = personList.Insert(Person{"Alice", 25})
	personList = personList.Insert(Person{"Bob", 30})

	fmt.Printf("Person ")
	personList.Print()
	fmt.Printf("List length: %d\n", personList.Length())
}
