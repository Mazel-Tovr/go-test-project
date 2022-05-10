package test

import (
	"testing"
)

// couldn't run (((

func TestMultiplicationOk(t *testing.T) {
	z := multiply(2, 2)
	if z != 2*2 {
		t.Error("test failed")
	}
}

func multiply(first, second int) int {
	return first * second
}
