package main

import (
	"io/ioutil"
	"strings"
)

func main() {

	inputFileName := "input"

	reader, err := ioutil.ReadFile(inputFileName)
	if err != nil {
		panic(err)
	}

	wires := make([]line, 2)
	wiresAsString := strings.Split(string(reader), "\n")
	for _, wireString := range wiresAsString {
		wires = append(wires, NewLineFromString(wireString))
	}

}

/*
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
*/
