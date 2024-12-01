package main

import (
	"adventOfCode/common"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

func splitLeftRight(input string) ([]int, []int) {
	lines := common.ToLines(input)
	var left []int
	var right []int
	for _, line := range lines {
		splits := strings.Split(line, "   ")
		l, _ := strconv.Atoi(splits[0])
		r, _ := strconv.Atoi(splits[1])
		left = append(left, l)
		right = append(right, r)
	}
	sort.Ints(left)
	sort.Ints(right)
	return left, right
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
