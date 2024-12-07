package main

import (
	"adventOfCode/common"
	"strconv"
)

const day = 7

func canMakeTarget(target int, nums []int) bool {
	if len(nums) == 2 {
		res := nums[0]+nums[1] == target || nums[0]*nums[1] == target

		return res
	}
	return canMakeTarget(target, append([]int{nums[0] + nums[1]}, nums[2:]...)) || canMakeTarget(target, append([]int{nums[0] * nums[1]}, nums[2:]...))
}

func Part1(input string) int {
	lines := common.ToLines(input)
	s := 0
	for _, line := range lines {
		nums := common.Ints(line)
		target := nums[0]
		if canMakeTarget(target, nums[1:]) {
			s += target
		}
	}
	return s
}
func canMakeTargetWithConcat(target int, nums []int) bool {
	first := strconv.Itoa(nums[0])
	second := strconv.Itoa(nums[1])

	sum := nums[0] + nums[1]
	prod := nums[0] * nums[1]
	concat, _ := strconv.Atoi(first + second)
	if len(nums) == 2 {
		res := sum == target || prod == target || concat == target
		return res
	}
	return canMakeTargetWithConcat(target, append([]int{sum}, nums[2:]...)) ||
		canMakeTargetWithConcat(target, append([]int{prod}, nums[2:]...)) ||
		canMakeTargetWithConcat(target, append([]int{concat}, nums[2:]...))
}

func Part2(input string) int {
	lines := common.ToLines(input)
	s := 0
	for _, line := range lines {
		nums := common.Ints(line)
		target := nums[0]
		b := canMakeTargetWithConcat(target, nums[1:])
		if b {
			s += target
		}
	}
	return s
}

func main() {
	common.AocMain(day, Part1, Part2)
}
