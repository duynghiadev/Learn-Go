package tracker

import "my-app/operations"

// Tracker struct to maintain state for operations
type Tracker struct {
	History map[string][]float64
}

// NewTracker initializes a new Tracker
func NewTracker() *Tracker {
	return &Tracker{History: make(map[string][]float64)}
}

// PerformOperation performs an operation and tracks the result
func (t *Tracker) PerformOperation(op operations.Operation, x, y float64) float64 {
	result := op.Action(x, y)
	t.History[op.Name] = append(t.History[op.Name], result)
	return result
}
