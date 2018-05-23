package structs

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

type StructTest struct {
	shape	  Shape
	perimeter float64
	area	  float64
}

const errorMsg = "Perimeter for %#v: expected %.2f, actual %.2f"
var tests = []StructTest{
	{Rectangle{10.0, 10.0}, 40.0, 100.0},
	{Circle{10.0}, 62.832, 314.159},
	{Triangle{6.0, 4.0}, 16.0, 12.0},
}

func TestPerimeter(t *testing.T) {

	assert := assert.New(t)
	checkPerimeter := func(shape Shape, expected float64) {
		actual := shape.Perimeter()
		assert.InDelta(expected, actual, 0.01, errorMsg, shape, expected,
			actual)
	}

	for _, test := range tests {
		checkPerimeter(test.shape, test.perimeter)
	}

}

func TestArea(t *testing.T) {

	assert := assert.New(t)
	checkArea := func(shape Shape, expected float64) {
		actual := shape.Area()
		assert.InDelta(expected, actual, 0.01, errorMsg, shape, expected,
			actual)
	}

	for _, test := range tests {
		checkArea(test.shape, test.area)
	}

}
