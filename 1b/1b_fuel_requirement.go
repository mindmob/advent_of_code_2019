package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func main() {

	inputFile := "input"

	reader, err := ioutil.ReadFile(inputFile)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(reader), "\n")

	sumOfFuel := 0.0
	for _, modulemass := range lines {
		mass, _ := strconv.ParseFloat(modulemass, 64)
		if len(strings.TrimSpace(modulemass)) != 0 {
			sumOfFuel += calculateModuleFuel(mass)
		}
	}
	fmt.Printf("Gesamtsume: %.0f", sumOfFuel)

}

func calculateModuleFuel(moduleMass float64) float64 {
	fuel := math.Floor(moduleMass/3.0) - 2.0
	//	fmt.Printf("Modulemass %f , calcFuel : %f \n", moduleMass, fuel)
	if fuel < 0.0 {
		return 0.0
	}
	return fuel + calculateModuleFuel(fuel)
}
