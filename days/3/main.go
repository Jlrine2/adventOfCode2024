package main

import (
	"adventOfCode/common"
	"regexp"
)

const day = 3

func Part1(input string) int {
	r, _ := regexp.Compile(`mul\(\d+,\d+\)`)
	matches := r.FindAllString(input, -1)
	s := 0
	for _, match := range matches {
		nums := common.Ints(match)
		s += nums[0] * nums[1]
	}
	common.Log(matches)
	return s
}

func Part2(input string) int {
	r, _ := regexp.Compile(`(mul\(\d+,\d+\))|(do\(\))|don't(\(\))`)
	matches := r.FindAllString(input, -1)
	s := 0
	enabled := true
	for _, match := range matches {
		if match == "do()" {
			enabled = true
			continue
		}
		if match == "don't()" {
			enabled = false
			continue
		}
		if enabled {
			nums := common.Ints(match)
			s += nums[0] * nums[1]
		}
	}
	return s
}

func main() {
	common.AocMain(day, Part1, Part2)
}
