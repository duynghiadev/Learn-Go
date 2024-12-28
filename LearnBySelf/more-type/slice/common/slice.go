package common

// Function to modify a slice
func ModifySlice(s []float64) {
	for i := range s {
		s[i] *= 1.1 // Increase each price by 10%
	}
}

// Function to update a map
func UpdateMap(m map[string]int, key string, newValue int) {
	if _, exists := m[key]; exists {
		m[key] = newValue
	}
}

// Function to update an array (will not modify the original array due to value semantics)
func UpdateArray(a [3]int) [3]int {
	for i := range a {
		a[i] += 10
	}
	return a
}
