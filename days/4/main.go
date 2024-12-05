package main

import (
	"adventOfCode/common"
)

const day = 4

func dirIs(p common.Point, dir string, value string, grid [][]int32) bool {
	for _, c := range value {
		cur, inBounds := p.Get(grid)
		if !inBounds {
			return false
		}
		if c != cur {
			return false
		}
		p = p.Go(dir)
	}
	return true
}

func Part1(input string) int {
	grid := common.ToGrid(input)
	s := 0
	for i, row := range grid {
		for j, _ := range row {
			p := common.Point{X: i, Y: j}
			for _, dir := range []string{
				"up",
				"down",
				"left",
				"right",
				"upleft",
				"upright",
				"downleft",
				"downright",
			} {
				if dirIs(p, dir, "XMAS", grid) {
					s++
				}
			}
		}
	}
	return s
}

func Part2(input string) int {
	grid := common.ToGrid(input)
	s := 0
	for i, row := range grid {
		for j, c := range row {
			p := common.Point{X: i, Y: j}
			if c == 'A' {
				ul := p.Go("upleft")
				dl := p.Go("downleft")
				if (dirIs(ul, "downright", "MAS", grid) || dirIs(ul, "downright", "SAM", grid)) &&
					(dirIs(dl, "upright", "MAS", grid) || dirIs(dl, "upright", "SAM", grid)) {
					s++
				}
			}
		}
	}
	return s
}

func main() {
	common.AocMain(day, Part1, Part2)
}
