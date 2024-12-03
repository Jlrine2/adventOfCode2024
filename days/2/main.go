package main

import (
	"adventOfCode/common"
	"slices"
)

const day = 2

func isSafe(nums []int) bool {
	accending := nums[0] < nums[1]
	min := nums[0] - nums[1]
	if min < 0 {
		min *= -1
	}
	max := min
	for i := 0; i < len(nums)-1; i++ {
		if accending != (nums[i] < nums[i+1]) {
			return false
		}
		diff := nums[i+1] - nums[i]
		if !accending {
			diff *= -1
		}
		if diff < min {
			min = diff
		}
		if diff > max {
			max = diff
		}
	}
	return 1 <= min && 3 >= max
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
		newNums = slices.Delete(newNums, i, i+1)
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
