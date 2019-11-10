package test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func init() {

}

// TestBeego is a sample to run an endpoint test
func TestBeego(t *testing.T) {
	assert.True(t, true, "True is true!")
}

// 计算并返回 x + 2.
func Calculate(x int) (result int) {
	result = x + 2
	return result
}
func TestCalculate(t *testing.T) {
	assert.Equal(t, Calculate(2), 4)
}

func TestStatusNotDown(t *testing.T) {
	assert.NotEqual(t, "down", "down")
}
func TestCalculate2(t *testing.T) {
	assert := assert.New(t)
	var tests = []struct {
		input    int
		expected int
	}{
		{2, 4},
		{-1, 1},
		{0, 2},
		{-5, -3},
		{99999, 100001},
	}
	
	for _, test := range tests {
		assert.Equal(Calculate(test.input), test.expected)
	}
}
