package main

import (
	"testing"

	"gotest.tools/assert"
)

func Test_OnlyAscending(t *testing.T) {

	assert.Equal(t, OnlyAscending("1111111"), true)
	assert.Equal(t, OnlyAscending("223450"), false)

}

func Test_HasDoubleDigits(t *testing.T) {

	assert.Equal(t, HasDoubleDigits("1111111"), true)
	assert.Equal(t, HasDoubleDigits("223450"), true)
	assert.Equal(t, HasDoubleDigits("123789"), false)

}
