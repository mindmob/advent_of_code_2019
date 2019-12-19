package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Instruction func(program []int, ip *int)

func main() {

	functionMap := map[int]Instruction{
		1: add,
		2: mult,
		3: input,
		4: output,
	}
	// rawProgram := "1,1,1,4,99,5,6,0,99"
	rawProgram := "3,225,1,225,6,6,1100,1,238,225,104,0,101,14,135,224,101,-69,224,224,4,224,1002,223,8,223,101,3,224,224,1,224,223,223,102,90,169,224,1001,224,-4590,224,4,224,1002,223,8,223,1001,224,1,224,1,224,223,223,1102,90,45,224,1001,224,-4050,224,4,224,102,8,223,223,101,5,224,224,1,224,223,223,1001,144,32,224,101,-72,224,224,4,224,102,8,223,223,101,3,224,224,1,223,224,223,1102,36,93,225,1101,88,52,225,1002,102,38,224,101,-3534,224,224,4,224,102,8,223,223,101,4,224,224,1,223,224,223,1102,15,57,225,1102,55,49,225,1102,11,33,225,1101,56,40,225,1,131,105,224,101,-103,224,224,4,224,102,8,223,223,1001,224,2,224,1,224,223,223,1102,51,39,225,1101,45,90,225,2,173,139,224,101,-495,224,224,4,224,1002,223,8,223,1001,224,5,224,1,223,224,223,1101,68,86,224,1001,224,-154,224,4,224,102,8,223,223,1001,224,1,224,1,224,223,223,4,223,99,0,0,0,677,0,0,0,0,0,0,0,0,0,0,0,1105,0,99999,1105,227,247,1105,1,99999,1005,227,99999,1005,0,256,1105,1,99999,1106,227,99999,1106,0,265,1105,1,99999,1006,0,99999,1006,227,274,1105,1,99999,1105,1,280,1105,1,99999,1,225,225,225,1101,294,0,0,105,1,0,1105,1,99999,1106,0,300,1105,1,99999,1,225,225,225,1101,314,0,0,106,0,0,1105,1,99999,108,226,677,224,1002,223,2,223,1006,224,329,1001,223,1,223,1007,226,226,224,1002,223,2,223,1006,224,344,101,1,223,223,1008,226,226,224,102,2,223,223,1006,224,359,1001,223,1,223,107,226,677,224,1002,223,2,223,1005,224,374,101,1,223,223,1107,677,226,224,102,2,223,223,1006,224,389,101,1,223,223,108,677,677,224,102,2,223,223,1006,224,404,1001,223,1,223,1108,677,226,224,102,2,223,223,1005,224,419,101,1,223,223,1007,677,226,224,1002,223,2,223,1006,224,434,101,1,223,223,1107,226,226,224,1002,223,2,223,1006,224,449,101,1,223,223,8,677,226,224,102,2,223,223,1006,224,464,1001,223,1,223,1107,226,677,224,102,2,223,223,1005,224,479,1001,223,1,223,1007,677,677,224,102,2,223,223,1005,224,494,1001,223,1,223,1108,677,677,224,102,2,223,223,1006,224,509,101,1,223,223,1008,677,677,224,102,2,223,223,1005,224,524,1001,223,1,223,107,226,226,224,1002,223,2,223,1005,224,539,101,1,223,223,7,226,226,224,102,2,223,223,1005,224,554,101,1,223,223,1108,226,677,224,1002,223,2,223,1006,224,569,1001,223,1,223,107,677,677,224,102,2,223,223,1005,224,584,101,1,223,223,7,677,226,224,1002,223,2,223,1005,224,599,101,1,223,223,108,226,226,224,1002,223,2,223,1005,224,614,101,1,223,223,1008,677,226,224,1002,223,2,223,1005,224,629,1001,223,1,223,7,226,677,224,102,2,223,223,1005,224,644,101,1,223,223,8,677,677,224,102,2,223,223,1005,224,659,1001,223,1,223,8,226,677,224,102,2,223,223,1006,224,674,1001,223,1,223,4,223,99,226"
	runProgram(convertInput(rawProgram), functionMap)
}

func runProgram(program []int, functionMap map[int]Instruction) {

	ip := 0
	for true {

		token := parseOpCode(program[ip])
		if token == 99 {
			fmt.Println("Endtoken Detected at :", ip)
			// fmt.Println("Result: ", program)
			return
		}
		function, ok := functionMap[token]
		if ok {
			function(program, &ip)
			// fmt.Println("Result: ", program)
			fmt.Println("new ip: ", ip)
		} else {
			fmt.Println("ERROR Invalid Token at Index ", ip, ": ", token)
			return
		}
	}
}

func mult(program []int, ip *int) {
	pMode := parseParamaterModes(program[*ip], 3)
	p1 := peek(program, *ip+1, pMode[0])
	p2 := peek(program, *ip+2, pMode[1])
	result := p1 * p2
	program[program[*ip+3]] = result
	*ip += 4
	fmt.Println("mult ", p1, " * ", p2, " = ", result, " => ", program[*ip+3])
}

func add(program []int, ip *int) {
	pMode := parseParamaterModes(program[*ip], 3)
	p1 := peek(program, *ip+1, pMode[0])
	p2 := peek(program, *ip+2, pMode[1])
	result := p1 + p2
	program[program[*ip+3]] = result
	*ip += 4
	fmt.Println("add ", p1, " + ", p2, " = ", result, " => ", program[*ip+3])
}

func input(program []int, ip *int) {
	optCode := 0
	fmt.Println("Input[", program[*ip+1], "]:", *ip)
	fmt.Scanf("%d", &optCode)
	program[program[*ip+1]] = optCode
	*ip += 2
}

func output(program []int, ip *int) {
	fmt.Println("Output: ", program[program[*ip+1]])
	*ip += 2
}

func parseOpCode(opCode int) int {
	if opCode < 100 {
		return opCode
	}
	instruction := strconv.Itoa(opCode)
	code, _ := strconv.Atoi(instruction[len(instruction)-2 : len(instruction)])
	return code
}

func parseParamaterModes(opCode int, paramCount int) []int {
	modes := make([]int, paramCount)
	for i := 1; i <= paramCount; i++ {
		modes[i-1] = getTheDigit(opCode/100, i)
	}
	return modes
}

// Returns the nth digit of a number.
// Examples:
// - 1st digit of 123 is 3.
// - 2nd digit of 123 is 2
// - 3rd digit of 123 is 1
// - 4th digit of 123 is 0
// (there are no more digits, so it returns 0 by default)
func getTheDigit(number int, digit int) int {
	return int((float64(number) / math.Pow(10, float64(digit)-1))) % 10
}

func peek(program []int, adress int, mode int) int {
	switch mode {
	case 0:
		return program[program[adress]]
	case 1:
		return program[adress]
	default:
		fmt.Println("ERROR Illegal parameter mode ", mode, " trying to read adress ", adress)
	}
	return 0
}

func convertInput(input string) []int {
	pgmAsString := strings.Split(input, ",")
	program := make([]int, 0)
	for _, number := range pgmAsString {
		i, _ := strconv.Atoi(number)
		program = append(program, i)
	}
	return program
}
