package main

import (
	"adventOfCode/common"
	"slices"
)

const day = 2

func decreasingSafe(nums []int) bool {
	last := nums[0]
	for _, num := range nums[1:] {
		if num >= last {
			return false
		}
		if last-num > 3 {
			return false
		}
		if last-num < 1 {
			return false
		}
		last = num
	}
	return true
}

func isSafe(nums []int) bool {
	a := decreasingSafe(nums)
	slices.Reverse(nums)
	b := decreasingSafe(nums)
	return a || b
}

func Part1(input string) int {
	lines := common.ToLines(input)
	safe := 0
	for _, line := range lines {
		nums := common.Ints(line)
		if isSafe(nums) {
			safe += 1
		}
	}
	return safe
}

func isSafeDapener(nums []int) bool {
	if isSafe(nums) {
		return true
	}
	for i := 0; i < len(nums); i++ {
		newNums := make([]int, len(nums))
		copy(newNums, nums)
		newNums = append(newNums[:i], newNums[i+1:]...)
		if isSafe(newNums) {
			return true
		}
	}
	return false
}

func Part2(input string) int {
	lines := common.ToLines(input)
	safe := 0
	for _, line := range lines {
		nums := common.Ints(line)
		s := isSafeDapener(nums)
		if s {
			safe += 1
		}
	}
	return safe
}

func main() {
	common.AocMain(day, Part1, Part2)
}
