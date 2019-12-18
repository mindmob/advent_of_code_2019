package main

import (
	"strconv"
	"strings"
)

type line interface {
	end() coordinate
	start() coordinate
}

type coordinate struct {
	x, y int
}

type simpleline struct {
	coord     coordinate
	direction coordinate
	length    int
}

func (l *simpleline) start() coordinate {
	return l.coord
}

func (l *simpleline) end() coordinate {
	return coordinate{
		x: l.coord.x + (l.direction.x * l.length),
		y: l.coord.y + (l.direction.y * l.length),
	}
}

func NewLine(length int, direction byte) *simpleline {
	l := simpleline{length: length}
	switch direction {
	case 'U':
		l.direction = coordinate{x: 0, y: 1}
	case 'D':
		l.direction = coordinate{x: 0, y: -1}
	case 'L':
		l.direction = coordinate{x: -1, y: 0}
	case 'R':
		l.direction = coordinate{x: 1, y: 0}
	}
	return &l
}

func NewLineFromString(lineAsString string) *simpleline {
	direction := lineAsString[0]
	length, _ := strconv.Atoi(lineAsString[1:len(lineAsString)])
	l := NewLine(length, direction)
	return l
}

func ParseWire(wireAsString string) []line {
	wire := make([]line, 1)
	currentPos := coordinate{x: 0, y: 0}

	lines := strings.Split(wireAsString, ",")
	for _, l := range lines {
		newLine := NewLineFromString(l)
		newLine.coord = currentPos
		currentPos = newLine.end()
		wire = append(wire, newLine)
	}

	return wire
}

/*
func Intersections(wire1 []line, wire2 []line) []coordinate {
	intersections := make([]coordinate, 0)

}
*/
func get_line_intersection(p0_x, p0_y, p1_x, p1_y, p2_x, p2_y, p3_x, p3_y int) (bool, int, int) {
	var s02_x, s02_y, s10_x, s10_y, s32_x, s32_y, s_numer, t_numer, denom, t, i_x, i_y float32
	s10_x = float32(p1_x) - float32(p0_x)
	s10_y = float32(p1_y) - float32(p0_y)
	s32_x = float32(p3_x) - float32(p2_x)
	s32_y = float32(p3_y) - float32(p2_y)

	denom = s10_x*s32_y - s32_x*s10_y
	if denom == 0 {
		return false, 0.0, 0.0
	} // Collinear
	denomPositive := denom > 0

	s02_x = float32(p0_x) - float32(p2_x)
	s02_y = float32(p0_y) - float32(p2_y)
	s_numer = s10_x*s02_y - s10_y*s02_x
	if (s_numer < 0) == denomPositive {
		return false, 0.0, 0.0 // No collision
	}

	t_numer = s32_x*s02_y - s32_y*s02_x
	if (t_numer < 0) == denomPositive {
		return false, 0.0, 0.0 // No collision
	}

	if ((s_numer > denom) == denomPositive) || ((t_numer > denom) == denomPositive) {
		return false, 0.0, 0.0 // No collision
	}
	// Collision detected
	t = t_numer / denom
	i_x = float32(p0_x) + (t * s10_x)
	i_y = float32(p0_y) + (t * s10_y)

	return true, int(i_x), int(i_y)
}
