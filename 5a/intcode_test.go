package main

import (
	"testing"

	"gotest.tools/assert"
)

func Test_parseParamaterModes(t *testing.T) {
	assert.DeepEqual(t, parseParamaterModes(1001, 3), []int{0, 1, 0})
	assert.DeepEqual(t, parseParamaterModes(1201, 3), []int{2, 1, 0})
	assert.DeepEqual(t, parseParamaterModes(1, 3), []int{0, 0, 0})
}

func Test_getTheDigit(t *testing.T) {
	assert.Equal(t, getTheDigit(123, 1), 3)
	assert.Equal(t, getTheDigit(123, 2), 2)
	assert.Equal(t, getTheDigit(123, 3), 1)
	assert.Equal(t, getTheDigit(123, 4), 0)
}

func Test_peek(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6}
	assert.Equal(t, peek(input, 1, 1), 2)
	assert.Equal(t, peek(input, 1, 0), 3)
}
