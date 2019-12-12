package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

type Coordinate struct {
	X, Y int
}

func main() {

	wirepath := make(map[Coordinate]int)

	inputFile := "input"

	reader, err := ioutil.ReadFile(inputFile)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(reader), "\n")

	for _, line := range lines {
		malen(wirepath, line)
	}
	//fmt.Printf("Gesamtsume: %.0f", sumOfFuel)

}

func malen(wirepath map[Coordinate]int, wiresString string) {

	currentPosition := Coordinate{0, 0}
	wires := strings.Split(wiresString, ",")
	for _, wire := range wires {
		direction := wire[0]
		length, _ := strconv.Atoi(wire[1:len(wire)])
		switch direction {
		case 'U':
			pos := 0
			for pos < length {
				newCoordinate := Coordinate{currentPosition.X, currentPosition.Y + pos}
				numberOfWires, ok := wirepath[newCoordinate]
				if ok {
					wirepath[newCoordinate] = numberOfWires + 1
				} else {
					wirepath[newCoordinate] = 1
				}
			}
			break
		case 'D':
			break
		case 'L':
			break
		case 'R':
			break
		}
	}

}
