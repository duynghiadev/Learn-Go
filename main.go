package main

import "fmt"

type Task struct {
	ID    int
	Title string
}

func updateValue(pointer *int) {
	fmt.Println("address of pointer before:", pointer)
	fmt.Println("value of pointer before:", *pointer)
	*pointer = 30
	fmt.Println("address of pointer after:", pointer)
	fmt.Println("value of pointer after:", *pointer)
}

func printTaskPointer(t *Task) {
	fmt.Println("Task ID:", t.ID)
	fmt.Println("Task Title:", t.Title)
}

func updateTaskTitle(t *Task, newTitle string) {
	t.Title = newTitle
}

func updateTaskNoPointer(t Task, newTitle string) {
	t.Title = newTitle // copy a struct
}

func createTask(id int, title string) *Task {
	return &Task{ID: id, Title: title}
}

func main() {
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
	fmt.Println("pointer1:", pointer1)

	task := Task{ID: 1, Title: "Learn Go"}
	taskPtr := createTask(2, "Learn Go backend")

	fmt.Println("Task ID:", taskPtr.ID)
	fmt.Println("Task title:", taskPtr.Title)

	printTaskPointer(&task)

	fmt.Println("Before update:", task.Title)

	updateTaskTitle(&task, "Learn Advanced Go")
	fmt.Println("After update:", task.Title)

	updateTaskNoPointer(task, "Learn Basic Go")
	fmt.Println(task.Title) // no change
}
