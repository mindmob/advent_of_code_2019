package main

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func Test_1(t *testing.T) {
	input := `10 ORE => 10 A
	5 A => 5 B
	5 A => 5 C
	5 A, 5 B => 1 FUEL`

	reactions := parseReactions(input)
	bag := make(map[string]int)
	oreCount := 0
	createMaterial(Material{"FUEL", 1}, bag, reactions, &oreCount)
	fmt.Println("Result: ", bag)
	fmt.Println("oreCount: ", oreCount)
	assert.Equal(t, oreCount, 10)
}

func Test_2(t *testing.T) {
	input := `10 ORE => 10 A
	1 ORE => 1 B
	7 A, 1 B => 1 C
	7 A, 1 C => 1 D
	7 A, 1 D => 1 E
	7 A, 1 E => 1 FUEL`

	reactions := parseReactions(input)
	bag := make(map[string]int)
	oreCount := 0
	createMaterial(Material{"FUEL", 1}, bag, reactions, &oreCount)
	fmt.Println("Result: ", bag)
	fmt.Println("oreCount: ", oreCount)
	assert.Equal(t, oreCount, 31)
}

func Test_3(t *testing.T) {
	input := `9 ORE => 2 A
	8 ORE => 3 B
	7 ORE => 5 C
	3 A, 4 B => 1 AB
	5 B, 7 C => 1 BC
	4 C, 1 A => 1 CA
	2 AB, 3 BC, 4 CA => 1 FUEL`

	reactions := parseReactions(input)
	bag := make(map[string]int)
	oreCount := 0
	createMaterial(Material{"FUEL", 1}, bag, reactions, &oreCount)
	fmt.Println("Result: ", bag)
	fmt.Println("oreCount: ", oreCount)
	assert.Equal(t, oreCount, 165)

}

func Test_4(t *testing.T) {
	input := `157 ORE => 5 NZVS
	165 ORE => 6 DCFZ
	44 XJWVT, 5 KHKGT, 1 QDVJ, 29 NZVS, 9 GPVTF, 48 HKGWZ => 1 FUEL
	12 HKGWZ, 1 GPVTF, 8 PSHF => 9 QDVJ
	179 ORE => 7 PSHF
	177 ORE => 5 HKGWZ
	7 DCFZ, 7 PSHF => 2 XJWVT
	165 ORE => 2 GPVTF
	3 DCFZ, 7 NZVS, 5 HKGWZ, 10 PSHF => 8 KHKGT`

	reactions := parseReactions(input)
	bag := make(map[string]int)
	oreCount := 0
	createMaterial(Material{"FUEL", 1}, bag, reactions, &oreCount)
	fmt.Println("Result: ", bag)
	fmt.Println("oreCount: ", oreCount)
	assert.Equal(t, oreCount, 13312)

}


/*
func Test_2(t *testing.T) {
	input := `2 ORE => 3 A
    2 A => 1 B
    2 B => 1 FUEL`

	reactions := parseReactions(input)
	ore := howMutchDoINeed(Material{"FUEL", 1}, reactions)
	fmt.Println(ore)
	// fmt.Println(countOre(howMutchDoINeed(Material{"FUEL", 1}, reactions)))
}
*/
