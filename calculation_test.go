package main

import (
	"testing"
)

func TestEqual(t *testing.T) {

	a := 5
	b := 6

	actualResult := 11
	result := Equal(a, b)

	t.Logf("%d + %d", a, b)

	if result != actualResult {
		t.Errorf("Wrong, it should be %d", actualResult)
	}

}

func Equal(a, b int) int {

	result := a + b
	return result
}
