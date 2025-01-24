package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	// Helper function to perform recursive walk
	var walker func(t *tree.Tree)
	walker = func(t *tree.Tree) {
		if t == nil {
			return
		}
		// In order traversal: Left -> Root -> Right
		walker(t.Left)
		ch <- t.Value
		walker(t.Right)
	}

	walker(t)
	close(ch) // Close channel when done
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	// Compare values from both channels
	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2

		if ok1 != ok2 {
			// One channel closed before the other
			return false
		}
		if !ok1 {
			// Both channels are closed
			return true
		}
		if v1 != v2 {
			// Values don't match
			return false
		}
	}
}

func main() {
	// Test Walk function
	fmt.Println("Testing Walk function:")
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	for v := range ch {
		fmt.Printf("%d ", v)
	}
	fmt.Println()

	// Test Same function
	fmt.Println("\nTesting Same function:")
	fmt.Printf("Same(tree.New(1), tree.New(1)) = %v\n", Same(tree.New(1), tree.New(1)))
	fmt.Printf("Same(tree.New(1), tree.New(2)) = %v\n", Same(tree.New(1), tree.New(2)))
}
