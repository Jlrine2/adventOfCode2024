package main

import (
	"adventOfCode/common"
	"slices"
)

const day = 1

func splitLeftRight(input string) ([]int, []int) {
	ints := common.Ints(input)
	collumns := common.Unzip(ints, 2)
	slices.Sort(collumns[0])
	slices.Sort(collumns[1])
	return collumns[0], collumns[1]
}

func Part1(input string) int {
	left, right := splitLeftRight(input)
	t := 0
	for i, a := range left {
		d := a - right[i]
		if d < 0 {
			d *= -1
		}
		t += d
	}
	return t
}

func Part2(input string) int {
	left, right := splitLeftRight(input)
	t := 0
	counts := map[int]int{}
	for _, a := range right {
		counts[a] = counts[a] + 1
	}
	for _, a := range left {
		t += a * counts[a]
	}
	return t
}

func main() {
	common.AocMain(day, Part1, Part2)
}
