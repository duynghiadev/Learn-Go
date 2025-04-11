package fibonacci

// NewFibonacci generates a standalone Fibonacci closure.
func NewFibonacci() func() int {
	n1, n2 := 0, 1
	return func() int {
		result := n1
		n1, n2 = n2, n1+n2
		return result
	}
}
