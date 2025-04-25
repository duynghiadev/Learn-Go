package assert

import (
	"reflect"
	"testing"
)

// AssertEqual check if two values are equal in a test
func Equal(t *testing.T, expected, actual interface{}) {
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected: %v, but got: %v", expected, actual)
	}
}
