package main

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

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
	bag["ORE"] = 1000000000000
	// for ok := true; ok; ok = (bag["ORE"] > 0) {
		createMaterial(Material{"FUEL", 100000000}, bag, reactions)
	// }
	fmt.Println("Result: ", bag)
	assert.Equal(t, bag["FUEL"], 82892753)

}

