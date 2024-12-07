package main

import "testing"

const testInput = `190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`

const p1expected = 3749
const p2expected = 11387

func TestPart1(t *testing.T) {
	res := Part1(testInput)
	if res != p1expected {
		t.Errorf("Part1 = %d; want %d", res, p1expected)
	}
}

func TestPart2(t *testing.T) {
	res := Part2(testInput)
	if res != p2expected {
		t.Errorf("Part2 = %d; want %d", res, p2expected)
	}
}
