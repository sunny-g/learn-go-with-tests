package integers

import (
	"testing"
)

func Assert(t *testing.T, actual int, expected int) {
	if actual != expected {
		t.Errorf("actual: '%d', expected: '%d'", actual, expected)
	}
}

func TestAdder(t *testing.T) {
	actual := Add(1, 2)
	expected := 3
	Assert(t, actual, expected)
}
