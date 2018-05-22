package arrays

import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
)

// Sum tests

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3}

	expected := 6
	actual := Sum(numbers)
	assert.Equal(t, expected, actual)
}

func BenchmarkSum(b *testing.B) {
	numbers := []int{1, 2, 3, 4, 5}

	for i := 0; i < b.N; i++ {
		Sum(numbers)
	}
}

func ExampleSum() {
	numbers := []int{1, 2, 3}
	sum := Sum(numbers)
	fmt.Println(sum)
	// Output: 6
}

// SumAll tests

func TestSumAll(t *testing.T) {
	slice1 := []int{1, 2, 3}
	slice2 := []int{4, 5, 6}

	expected := []int{6, 15}
	actual := SumAll(slice1, slice2)
	assert.Equal(t, expected, actual)
}

func BenchmarkSumAll(b *testing.B) {
	slice1 := []int{1, 2, 3}
	slice2 := []int{4, 5, 6}

	for i := 0; i < b.N; i++ {
		SumAll(slice1, slice2)
	}
}

func ExampleSumAll() {
	slice1 := []int{1, 2, 3}
	slice2 := []int{4, 5, 6}

	sums := SumAll(slice1, slice2)
	fmt.Println(sums)
	// Output: [6 15]
}

// SumAllTails tests

func TestSumAllTails(t *testing.T) {

	t.Run("should sum non-empty slices", func(t *testing.T) {
		slice1 := []int{1, 2, 3}
		slice2 := []int{4, 5, 6}

		expected := []int{5, 11}
		actual := SumAllTails(slice1, slice2)
		assert.Equal(t, expected, actual)
	})

	t.Run("should sum empty slices", func(t *testing.T) {
		slice1 := []int{}
		slice2 := []int{}

		expected := []int{0, 0}
		actual := SumAllTails(slice1, slice2)
		assert.Equal(t, expected, actual)
	})
}
