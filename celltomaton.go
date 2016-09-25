package main

import (
	"fmt"
	"github.com/initiumsrc/binary"
	"os"
)

var rules map[int]int

func main() {
	args := os.Args[1:]

	for i := 0; i < len(args); i++ {
		fmt.Println(args[i])
	}

	rules = constructRules(3)

	matrix := make([][]int, 10, 10)

	initialRow := make([]int, 20, 20)
	initialRow[5] = 1
	initialRow[0] = 1
	initialRow[1] = 1

	matrix[0] = initialRow
	matrix[1] = generateRow(initialRow)

	fmt.Println(matrix[0])
	fmt.Println(matrix[1])
}

func constructRules(rule int) map[int]int {
	if rule > 255 || rule < 0 {
		panic("Rules are defined between 0 and 255")
	}

	m := make(map[int]int)

	ruleArr := binary.PaddedBinaryArray(binary.IntToBinaryArray(rule), 8)

	for i := 0; i < 8; i++ {
		m[i] = ruleArr[len(ruleArr)-1-i]
	}

	return m
}

func generateRow(prev []int) []int {
	genArr := make([]int, len(prev), len(prev))

	for i := 0; i < len(prev); i++ {
		if i == 0 {
			first := []int{prev[len(prev)-1], prev[i], prev[i+1]}
			genArr[i] = rules[binary.BinaryArrayToInt(&first)]
		} else if i == len(prev)-1 {
			last := []int{prev[i-1], prev[i], prev[0]}
			genArr[i] = rules[binary.BinaryArrayToInt(&last)]
		} else {
			current := prev[i-1 : i+1]
			genArr[i] = rules[binary.BinaryArrayToInt(&current)]
		}
	}

	return genArr
}
