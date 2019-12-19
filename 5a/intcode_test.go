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

func Test_jumpIfTrue(t *testing.T) {
	input := []int{1105, 2, 16}
	ip := 0
	jumpIfTrue(input, &ip)
	assert.Equal(t, ip, 16)

	input = []int{1105, 0, 16}
	ip = 0
	jumpIfTrue(input, &ip)
	assert.Equal(t, ip, 3)

	input = []int{5, 0, 0}
	ip = 0
	jumpIfTrue(input, &ip)
	assert.Equal(t, ip, 5)

}

func Test_jumpIfFalse(t *testing.T) {
	input := []int{1106, 2, 16}
	ip := 0
	jumpIfFalse(input, &ip)
	assert.Equal(t, ip, 3)

	input = []int{1106, 0, 16}
	ip = 0
	jumpIfFalse(input, &ip)
	assert.Equal(t, ip, 16)

	input = []int{6, 0, 0}
	ip = 0
	jumpIfFalse(input, &ip)
	assert.Equal(t, ip, 3)

}

func Test_lessThan(t *testing.T) {
	input := []int{1107, 2, 16, 0}
	ip := 0
	lessThan(input, &ip)
	assert.Equal(t, input[0], 1)

	input = []int{1107, 3, 1, 0}
	ip = 0
	lessThan(input, &ip)
	assert.Equal(t, input[0], 0)
}

func Test_equal(t *testing.T) {
	input := []int{1108, 2, 16, 0}
	ip := 0
	equals(input, &ip)
	assert.Equal(t, input[0], 0)

	input = []int{1108, 2, 2, 0}
	ip = 0
	equals(input, &ip)
	assert.Equal(t, input[0], 1)
}
