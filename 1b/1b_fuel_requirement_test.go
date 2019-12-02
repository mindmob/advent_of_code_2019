package main

import (
	"testing"

	"gotest.tools/assert"
)

func TestCalculateFuel(t *testing.T) {

	assert.Equal(t, calculateModuleFuel(14.0), 2.0)
	assert.Equal(t, calculateModuleFuel(1969.0), 966.0)
	assert.Equal(t, calculateModuleFuel(100756.0), 50346.0)

}
