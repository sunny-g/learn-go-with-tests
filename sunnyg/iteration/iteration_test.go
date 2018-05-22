package iteration

import (
	"testing"
	"fmt"
)

func Assert(t *testing.T, actual string, expected string) {
	if (actual != expected) {
		t.Errorf("actual: '%s', expected: '%s'", actual, expected)
	}
}

func TestRepeat(t *testing.T) {
	actual := Repeat("a", 5)
	expected := "aaaaa"
	Assert(t, actual, expected)
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat_first() {
	str := Repeat("a", 5)
	fmt.Println(str)
	// Output: aaaaa
}

func ExampleRepeat_second() {
	str := Repeat("aa", 2)
	fmt.Println(str)
	// Output: aaaa
}
