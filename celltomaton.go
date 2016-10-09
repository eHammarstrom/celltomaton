package main

import "github.com/initiumsrc/binary"

var rules map[int]int

func Get(initialRow []int, height int, rule int) [][]int {
	rules = constructRules(rule)

	matrix := make([][]int, height, height)
	matrix[0] = initialRow

	for i := 1; i < height; i++ {
		matrix[i] = generateRow(matrix[i-1])
	}

	return matrix
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
			current := []int{prev[i-1], prev[i], prev[i+1]}
			genArr[i] = rules[binary.BinaryArrayToInt(&current)]
		}
	}

	return genArr
}
