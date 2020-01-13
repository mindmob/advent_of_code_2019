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
	assert.Equal(t, wire1[len(wire1)-1].end(), coordinate{3, 2})

}

func Test_NewLineWithCoordinate(t *testing.T) {

	line1 := NewLineWithCoordinate(2, 'U', 1, 0)
	assert.Equal(t, line1.start().x, 1)
	assert.Equal(t, line1.start().y, 0)
	assert.Equal(t, line1.end().x, 1)
	assert.Equal(t, line1.end().y, 2)

	line2 := NewLineWithCoordinate(2, 'D', 1, 0)
	assert.Equal(t, line2.start().x, 1)
	assert.Equal(t, line2.start().y, 0)
	assert.Equal(t, line2.end().x, 1)
	assert.Equal(t, line2.end().y, -2)

}

func Test_get_line_intersection(t *testing.T) {

	line1 := NewLineWithCoordinate(2, 'U', 1, 0)
	line2 := NewLineWithCoordinate(2, 'U', 0, 1)
	intersects, _, _ := get_line_intersection(line1, line2)
	assert.Equal(t, intersects, false)

	line1 = NewLineWithCoordinate(2, 'U', 1, 0)
	line2 = NewLineWithCoordinate(2, 'R', 0, 1)
	intersects, px, py := get_line_intersection(line1, line2)
	assert.Equal(t, intersects, true)
	assert.Equal(t, px, 1)
	assert.Equal(t, py, 1)

	line1 = NewLineWithCoordinate(2, 'R', -1, 0)
	line2 = NewLineWithCoordinate(2, 'U', 0, -1)
	intersects, px, py = get_line_intersection(line1, line2)
	assert.Equal(t, intersects, true)
	assert.Equal(t, px, 0)
	assert.Equal(t, py, 0)

}

func Test_Intersections(t *testing.T) {

	wire1 := ParseWire("R8,U5,L5,D3")
	wire2 := ParseWire("U7,R6,D4,L4")
	intersect1 := coordinate{x: 6, y: 5}
	intersect2 := coordinate{x: 3, y: 3}
	intersections := Intersections(wire1, wire2)
	assert.Equal(t, len(intersections), 2)
	assert.Equal(t, intersections[0], intersect1)
	assert.Equal(t, intersections[1], intersect2)

}

func Test_FindShortestVector(t *testing.T) {

	lines1 := ParseWire("R8,U5,L5,D3")
	lines2 := ParseWire("U7,R6,D4,L4")
	intersections := Intersections(lines1, lines2)
	nearest := FindShortestVector(intersections)
	distance := ManhattanDistance(nearest)
	assert.Equal(t, distance, 6)

	lines1 = ParseWire("R75,D30,R83,U83,L12,D49,R71,U7,L72")
	lines2 = ParseWire("U62,R66,U55,R34,D71,R55,D58,R83")
	intersections = Intersections(lines1, lines2)
	nearest = FindShortestVector(intersections)
	distance = ManhattanDistance(nearest)
	assert.Equal(t, distance, 159)

	lines1 = ParseWire("R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51")
	lines2 = ParseWire("U98,R91,D20,R16,D67,R40,U7,R15,U6,R7")
	intersections = Intersections(lines1, lines2)
	nearest = FindShortestVector(intersections)
	distance = ManhattanDistance(nearest)
	assert.Equal(t, distance, 135)

}
