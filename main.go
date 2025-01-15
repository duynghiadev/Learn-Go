package main

import "fmt"

// Add two numbers and return the result
func add(a int, b int) int {
	return a + b
}

// Subtract two numbers and return the result
func subtract(a int, b int) int {
	return a - b
}

// Example struct with fields
type Person struct {
	Name string
	Age  int
}

// Task struct for task-related operations
type Task struct {
	ID    int
	Title string
}

// Function to update the value of a pointer
func updateValue(pointer *int) {
	fmt.Println("Address of pointer before:", pointer)
	fmt.Println("Value of pointer before:", *pointer)
	*pointer = 30
	fmt.Println("Address of pointer after:", pointer)
	fmt.Println("Value of pointer after:", *pointer)
}

// Function to print Task details using a pointer
func printTaskPointer(t *Task) {
	fmt.Println("Task ID:", t.ID)
	fmt.Println("Task Title:", t.Title)
}

// Function to update Task title using a pointer
func updateTaskTitle(t *Task, newTitle string) {
	t.Title = newTitle
}

// Function to update Task title without a pointer
func updateTaskNoPointer(t Task, newTitle string) {
	t.Title = newTitle // This only updates the local copy
}

// Function to create a new Task and return a pointer to it
func createTask(id int, title string) *Task {
	return &Task{ID: id, Title: title}
}

func main() {
	// Arithmetic operations
	addResult := add(10, 20)
	subtractResult := subtract(30, 10)

	// Working with Person struct
	p := Person{
		Name: "Alice",
		Age:  30,
	}
	fmt.Println("Person:", p)

	// Constant value
	const pi = 3.14
	fmt.Println("Pi:", pi)

	// Range over a slice
	numbers := []int{1, 2, 3, 4, 5}
	for i, num := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", i, num)
	}

	// Working with pointers
	var x int = 10
	var ptr *int = &x
	fmt.Println("Value of x:", x)
	fmt.Println("Address of x:", &x)
	fmt.Println("Value of ptr:", ptr)
	fmt.Println("Value at ptr:", *ptr)

	*ptr = 20
	fmt.Println("Updated value of x:", x)

	var pointer1 int = 40
	updateValue(&pointer1)
	fmt.Println("Pointer1 after updateValue:", pointer1)

	// Working with Task struct
	task := Task{ID: 1, Title: "Learn Go"}
	taskPtr := createTask(2, "Learn Go backend")

	fmt.Println("TaskPtr ID:", taskPtr.ID)
	fmt.Println("TaskPtr Title:", taskPtr.Title)

	printTaskPointer(&task)

	fmt.Println("Before update:", task.Title)

	updateTaskTitle(&task, "Learn Advanced Go")
	fmt.Println("After update (using pointer):", task.Title)

	updateTaskNoPointer(task, "Learn Basic Go")
	fmt.Println("After update (without pointer):", task.Title)

	// Print arithmetic results
	fmt.Println("Addition Result:", addResult)
	fmt.Println("Subtraction Result:", subtractResult)
}
