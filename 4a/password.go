package main

import "strconv"

import "fmt"

func main() {

	input1 := 109165
	input2 := 576723
	count := 0
	for i := input1; i <= input2; i++ {
		asstring := strconv.Itoa(i)
		if OnlyAscending(asstring) && HasDoubleDigits(asstring) {
			count++
		}
	}
	fmt.Println("Number of matches: ", count)

}

func OnlyAscending(pass string) bool {

	current := 0
	for _, c := range pass {
		digit := int(c - '0')
		if digit < current {
			return false
		}
		current = digit
	}

	return true

}

func HasDoubleDigits(pass string) bool {

	current := -1
	for _, c := range pass {
		digit := int(c - '0')
		if digit == current {
			return true
		}
		current = digit
	}

	return false

}
