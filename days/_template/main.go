package main

import (
	"adventOfCode/common"
	"fmt"
	"time"
)

func Part1(input string) int {
	return 0
}

func Part2(input string) int {
	return 1
}

func main() {
	input := common.ReadInput("days/1/input.txt")
	p1start := time.Now()
	Part1(input)
	p1elapsed := time.Since(p1start)

	p2start := time.Now()
	Part2(input)
	p2elapsed := time.Since(p2start)
	fmt.Printf("Part1:\nsolution: %#v\ntime: %#v\n", Part1(input), p1elapsed.String())
	fmt.Printf("Part2:\nsolution: %#v\ntime: %#v\n", Part2(input), p2elapsed.String())
}
