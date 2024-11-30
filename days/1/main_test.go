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
		t.Errorf("Solve(%q) = %d; want %d", testInput, res, p1expected)
	}
}

func TestPart2(t *testing.T) {
	res := Part2(testInput)
	if res != p2expected {
		t.Errorf("Solve(%q) = %d; want %d", testInput, res, p2expected)
	}
}
