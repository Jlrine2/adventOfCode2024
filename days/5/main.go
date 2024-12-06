package main

import (
	"adventOfCode/common"
)

const day = 5

type rule struct {
	First  int
	Second int
}

func checkRule(rule rule, pages []int) bool {
	hasFirst := false
	hasSecond := false
	for _, page := range pages {
		if page == rule.First {
			if hasSecond {
				return false
			}
			hasFirst = true
		}
		if page == rule.Second {
			if hasFirst {
				return true
			}
			hasSecond = true
		}
	}
	return !(hasFirst && hasSecond)
}

func Part1(input string) int {
	lines := common.ToLines(input)
	var rules []rule

	s := 0
	doneRules := false
	for _, line := range lines {
		if line == "" {
			doneRules = true
			continue
		}
		nums := common.Ints(line)
		if !doneRules {
			rules = append(rules, rule{First: nums[0], Second: nums[1]})
			continue
		}
		isOrdered := true
		for _, rule := range rules {
			if !checkRule(rule, nums) {
				isOrdered = false
			}
		}
		if isOrdered {

			m := nums[len(nums)/2]
			s += m
		}
	}
	return s
}

func sort(pages []int, rules []rule) {
	for true {
		pass := true
		for _, rule := range rules {
			if !checkRule(rule, pages) {
				pass = false
				for i := 0; i < len(pages); i++ {
					if pages[i] == rule.Second {
						pages[i], pages[i+1] = pages[i+1], pages[i]
					}
					if pages[i] == rule.First {
						break
					}
				}
			}
		}
		if pass {
			return
		}
	}
}

func Part2(input string) int {
	lines := common.ToLines(input)
	var rules []rule

	s := 0
	doneRules := false
	for _, line := range lines {
		if line == "" {
			doneRules = true
			continue
		}
		nums := common.Ints(line)
		if !doneRules {
			rules = append(rules, rule{First: nums[0], Second: nums[1]})
			continue
		}
		isOrdered := true
		for _, rule := range rules {
			if !checkRule(rule, nums) {
				isOrdered = false
			}
		}
		if !isOrdered {
			sort(nums, rules)
			l := len(nums)
			if l%2 != 0 {
				l--
			}
			m := nums[l/2]
			common.Log(m, nums)
			s += m
		}
	}
	return s
}

func main() {
	common.AocMain(day, Part1, Part2)
}
