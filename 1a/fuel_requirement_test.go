package main

import (
	"testing"

	"gotest.tools/assert"
)

func TestCalculateFuel(t *testing.T) {

	fuel := calculateFuel(12.0)
	assert.Equal(t, fuel, 2.0)
	assert.Equal(t, calculateFuel(14.0), 2.0)
	assert.Equal(t, calculateFuel(1969.0), 654.0)
	assert.Equal(t, calculateFuel(100756.0), 33583.0)

}
