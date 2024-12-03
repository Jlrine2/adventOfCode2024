package main

import "testing"

const testInput = `line1
line2
`

const p1expected = 0
const p2expected = 1

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
