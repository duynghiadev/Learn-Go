package operations

// Operation struct to encapsulate operation name and logic
type Operation struct {
	Name   string
	Action func(float64, float64) float64
}
