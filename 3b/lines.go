package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type line interface {
	end() coordinate
	start() coordinate
	Length() int
	Direction() coordinate
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

func (l *simpleline) Length() int {
	return l.length
}

func (l *simpleline) Direction() coordinate {
	return l.direction
}

func main() {

	input1 := "R1009,D117,L888,D799,L611,U766,L832,U859,L892,D79,R645,U191,L681,D787,R447,D429,L988,U536,L486,D832,R221,D619,R268,D545,L706,U234,L528,D453,R493,D24,L688,U658,L74,D281,R910,D849,L5,U16,R935,D399,L417,U609,R22,D782,L432,D83,L357,D982,L902,U294,L338,U102,R342,D621,R106,U979,L238,U158,R930,D948,L700,D808,R445,U897,R980,U227,L466,D416,R244,U396,R576,U157,R548,U795,R709,U550,R137,U212,L977,U786,L423,D792,R391,D974,R390,U771,R270,D409,L917,D9,R412,D699,L170,D276,L912,U710,R814,U656,R4,D800,R596,U970,L194,U315,L845,D490,L303,U514,L675,D737,L880,D86,L253,D525,R861,D5,R424,D113,L764,D900,R485,D421,R125,U684,R53,U96,L871,U260,R456,U378,L448,D450,L903,D482,R750,U961,R264,D501,R605,D367,R550,U642,R228,U164,L343,U868,R595,D318,R452,U845,L571,D281,R49,D889,L481,U963,R182,U358,R454,U267,L790,D252,R455,D188,L73,U256,L835,D816,R503,U895,L259,U418,R642,U818,L187,U355,R772,U466,R21,U91,R707,D349,L200,U305,R931,D982,L334,D416,L247,D935,L326,U449,L398,D914,R602,U10,R762,D944,L639,D141,L457,U579,L198,U527,R750,U167,R816,D753,R850,D281,L712,D583,L172,D254,L544,D456,R966,U839,R673,D479,R730,D912,R992,D969,R766,U205,R477,D719,R172,D735,R998,D687,R698,D407,R172,U945,R199,U348,L256,D876,R580,U770,L483,D437,R353,D214,R619,U541,R234,D962,R842,U639,R520,D354,L279,D15,R42,U138,L321,D376,L628,D893,L670,D574,L339,U298,L321,D120,L370,U408,L333,D353,L263,D79,R535,D487,R113,D638,R623,D59,L508,D866,R315,U166,L534,U927,L401,D626,L19,D994,L778,D317,L936,U207,L768,U948,R452,U165,R864,D283,L874"
	input2 := "L995,D93,L293,U447,L793,D605,R497,D155,L542,D570,R113,D779,L510,U367,L71,D980,R237,U290,L983,U49,R745,U182,L922,D174,L189,D629,R315,D203,R533,U72,L981,D848,L616,U654,R445,D864,R526,D668,L678,U378,L740,D840,L202,D429,R136,D998,L116,D554,L893,U759,R617,U942,R999,U582,L220,U447,R895,D13,R217,U743,L865,U950,R91,D381,R662,D518,L798,D637,L213,D93,L231,D185,R704,U581,L268,U773,R405,U862,R796,U73,L891,U553,L952,U450,R778,D868,R329,D669,L182,U378,L933,D83,R574,U807,R785,D278,R139,D362,R8,U546,R651,U241,L462,D309,L261,D307,L85,U701,L913,U271,R814,U723,L777,D256,R417,U814,L461,U652,R198,D747,R914,U520,R806,U956,L771,D229,R984,U685,R663,D812,R650,U214,R839,U574,L10,U66,R644,D371,L917,D819,L73,D236,R277,U611,R390,U723,L129,D496,L552,D451,R584,U105,L805,U165,R179,D372,L405,D702,R14,U332,L893,D419,R342,D146,R907,D672,L316,U257,L903,U919,L942,U771,R879,U624,L280,U150,L320,U220,R590,D242,R744,U291,R562,U418,L898,U66,L564,U495,R837,D555,L739,D780,R409,D122,L426,D857,R937,D600,R428,D592,R727,U917,R256,D680,L422,U630,L14,U240,R617,D664,L961,D554,L302,U925,L376,D187,L700,D31,L762,U397,L554,D217,R679,D683,R680,D572,R54,D164,L940,D523,R140,U52,L506,D638,R331,D415,R389,D884,R410,D62,R691,U665,R889,U864,L663,D690,R487,U811,L190,U780,L758,U267,R155,D344,L133,D137,R93,D229,L729,U878,L889,D603,R288,U890,R251,U531,L249,D995,R863,D257,R655,D311,R874,U356,L833,U151,L741,U246,R694,D899,L48,U915,L900,U757,L861,U402,R971,U537,R460,D844,R54,U956,L151,U74,R892,U248,R677,D881,R99,D931,R427"
	wires1 := ParseWire(input1)
	wires2 := ParseWire(input2)
	res := FindShortestIntersection(wires1, wires2)
	fmt.Println("Nearest intersection: ", res)

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
	l.coord = coordinate{x: 0, y: 0}
	return &l
}

func NewLineWithCoordinate(length int, direction byte, px, py int) *simpleline {
	l := NewLine(length, direction)
	l.coord = coordinate{x: px, y: py}
	return l
}

func NewLineFromString(lineAsString string) *simpleline {
	direction := lineAsString[0]
	length, _ := strconv.Atoi(lineAsString[1:len(lineAsString)])
	l := NewLine(length, direction)
	return l
}

func ParseWire(wireAsString string) []line {
	wire := make([]line, 0)
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

func Intersections(wire1 []line, wire2 []line) []coordinate {
	intersections := make([]coordinate, 0)
	for _, l1 := range wire1 {
		for _, l2 := range wire2 {
			doesIntersect, px, py := get_line_intersection(l1, l2)
			if doesIntersect && (px != 0 || py != 0) {
				intersections = append(intersections, coordinate{x: px, y: py})
			}
		}
	}
	return intersections
}

func FindShortestIntersection(wire1, wire2 []line) int {

	intersections := Intersections(wire1, wire2)
	res := 999999999999999
	for _, i := range intersections {
		d1 := CountStepsToCoordinate(wire1, i)
		d2 := CountStepsToCoordinate(wire2, i)
		val := d1 + d2
		if val < res {
			res = val
		}
	}
	return res

}

func CountStepsToCoordinate(wire []line, coord coordinate) int {

	steps := 0
	for _, l := range wire {
		if IsPointOnLine(l, coord) {
			offset := Minus(coord, l.start())
			steps += Length(offset)
			return steps
		}
		steps += l.Length()
	}
	return -1

}

func IsPointOnLine(line line, coord coordinate) bool {

	if line.start().x != coord.x && line.start().y != coord.y {
		return false
	}

	if line.start() == coord {
		return true
	}

	p := Minus(coord, line.start())
	normal := Normal(p)
	if normal != line.Direction() {
		return false
	}

	if Length(p) > line.Length() {
		return false
	}

	return true

}

func FindShortestVector(intersections []coordinate) coordinate {

	min := 999999999999999
	var minCoord coordinate
	for _, intersection := range intersections {
		l := ManhattanDistance(intersection)
		if l < min {
			min = l
			minCoord = intersection
		}
	}
	return minCoord
}

func get_line_intersection(line1, line2 line) (bool, int, int) {
	var s02_x, s02_y, s10_x, s10_y, s32_x, s32_y, s_numer, t_numer, denom, t, i_x, i_y float32
	s10_x = float32(line1.end().x) - float32(line1.start().x)
	s10_y = float32(line1.end().y) - float32(line1.start().y)
	s32_x = float32(line2.end().x) - float32(line2.start().x)
	s32_y = float32(line2.end().y) - float32(line2.start().y)

	denom = s10_x*s32_y - s32_x*s10_y
	if denom == 0 {
		return false, 0.0, 0.0
	} // Collinear
	denomPositive := denom > 0

	s02_x = float32(line1.start().x) - float32(line2.start().x)
	s02_y = float32(line1.start().y) - float32(line2.start().y)
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
	i_x = float32(line1.start().x) + (t * s10_x)
	i_y = float32(line1.start().y) + (t * s10_y)

	return true, int(i_x), int(i_y)
}

func ManhattanDistance(p coordinate) int {
	return int(math.Abs(float64(p.x)) + math.Abs(float64(p.y)))
}

func Length(p coordinate) int {
	return int(math.Sqrt(float64((p.x * p.x) + (p.y * p.y))))
}

func Minus(p coordinate, q coordinate) coordinate {
	return coordinate{p.x - q.x, p.y - q.y}
}

func Plus(p coordinate, q coordinate) coordinate {
	return coordinate{p.x + q.x, p.y + q.y}
}

func Mult(p coordinate, x int) coordinate {
	return coordinate{p.x * x, p.y * x}
}

func Div(p coordinate, x int) coordinate {
	return coordinate{p.x / x, p.y / x}
}

func Abs(p coordinate) coordinate {
	return coordinate{int(math.Abs(float64(p.x))), int(math.Abs(float64(p.y)))}
}

func Normal(p coordinate) coordinate {
	l := int(Length(p))
	if l == 0 {
		return coordinate{0, 0}
	}
	return coordinate{p.x / l, p.y / l}
}
