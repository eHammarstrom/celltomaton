package main

import (
	"fmt"
	"github.com/initiumsrc/binary"
	//"math"
	"os"
)

func main() {
	args := os.Args[1:]

	for i := 0; i < len(args); i++ {
		fmt.Println(args[i])
	}
}

func constructRules(rule int) map[int]int {
	m := make(map[int]int)

	if rule%2 == 0 {
		//pos := int(math.Log2(float64(rule)))
	}

	return m
}
