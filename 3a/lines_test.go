package main

import (
	"testing"
	"gotest.tools/assert"
)

func Test_NewLine(t *testing.T) {

	assert.Equal(t, NewLine(1, 'U').direction, coordinate{x: 0, y: 1})
	assert.Equal(t, NewLine(5, 'D').direction, coordinate{x: 0, y: -1})
	assert.Equal(t, NewLine(2, 'L').direction, coordinate{x: -1, y: 0})
	assert.Equal(t, NewLine(2, 'R').direction, coordinate{x: 1, y: 0})

}

func Test_NewLineFromString(t *testing.T) {

	assert.Equal(t, *NewLineFromString("R75"), simpleline{length: 75, direction: coordinate{x: 1, y: 0}})
	assert.Equal(t, *NewLineFromString("D30"), simpleline{length: 30, direction: coordinate{x: 0, y: -1}})
	assert.Equal(t, *NewLineFromString("U83"), simpleline{length: 83, direction: coordinate{x: 0, y: 1}})

}

func Test_lineEnd(t *testing.T) {

	assert.Equal(t, NewLineFromString("R75").end(), coordinate{75, 0})
	l := NewLineFromString("U75")
	l.coord = coordinate{10, 5}
	assert.Equal(t, l.end(), coordinate{10, 80})
	l2 := NewLineFromString("D75")
	l2.coord = coordinate{10, 5}
	assert.Equal(t, l2.end(), coordinate{10, -70})

}

func Test_ParseWire(t *testing.T) {

	wireString1 := "R8,U5,L5,D3"
	wire1 := ParseWire(wireString1)
	assert.Equal(t, wire1[len(wire1) - 1].end(), coordinate{3, 2})

}

func Test_get_line_intersection(t *testing.T) {

	intersects, px, py := get_line_intersection(0, 1, 0, 2, 1, 0, 2, 1)

}