package main

import (
	"adventOfCode/common"
)

const day = 8

func getAntennas(grid [][]int32) map[int32][]common.Point {
	antennas := make(map[int32][]common.Point)
	for i, _ := range grid {
		for j, _ := range grid[i] {
			if grid[i][j] != '.' {
				antennas[grid[i][j]] = append(antennas[grid[i][j]], common.Point{X: i, Y: j})
			}
		}
	}
	return antennas
}

func getExteriorAntiNodesForPair(a common.Point, b common.Point) (common.Point, common.Point) {
	xDistance := a.X - b.X
	yDistance := a.Y - b.Y
	return common.Point{X: a.X + xDistance, Y: a.Y + yDistance}, common.Point{X: b.X - xDistance, Y: b.Y - yDistance}
}

func Part1(input string) int {
	grid := common.ToGrid(input)
	antennas := getAntennas(grid)
	uniqueAntinodes := make(map[common.Point]bool)
	for _, v := range antennas {
		pairs := common.Pairs(v)
		for _, pair := range pairs {
			a, b := getExteriorAntiNodesForPair(pair[0], pair[1])
			if !(a.X < 0 || a.X >= len(grid) || a.Y < 0 || a.Y >= len(grid[0])) {
				uniqueAntinodes[a] = true
			}
			if !(b.X < 0 || b.X >= len(grid) || b.Y < 0 || b.Y >= len(grid[0])) {
				uniqueAntinodes[b] = true
			}
		}
	}
	return len(uniqueAntinodes)
}

func getResonantAntiNodesForPair(a common.Point, b common.Point, maxX int, maxY int) []common.Point {
	resonantAntiNodes := []common.Point{a, b}
	xDistance := a.X - b.X
	yDistance := a.Y - b.Y
	for i := 1; true; i++ {
		if (i > maxX) && (i > maxY) {
			break
		}
		resonantAntiNodes = append(resonantAntiNodes, common.Point{X: a.X + xDistance*i, Y: a.Y + yDistance*i}, common.Point{X: b.X - xDistance*i, Y: b.Y - yDistance*i})

	}
	return resonantAntiNodes
}

func Part2(input string) int {
	grid := common.ToGrid(input)
	antennas := getAntennas(grid)
	uniqueAntinodes := make(map[common.Point]bool)
	for _, v := range antennas {
		pairs := common.Pairs(v)
		for _, pair := range pairs {
			resonantPairs := getResonantAntiNodesForPair(pair[0], pair[1], len(grid), len(grid[0]))
			for _, a := range resonantPairs {
				if !(a.X < 0 || a.X >= len(grid) || a.Y < 0 || a.Y >= len(grid[0])) {
					uniqueAntinodes[a] = true
				}
			}
		}
	}
	common.Log(uniqueAntinodes)
	return len(uniqueAntinodes)
}

func main() {
	common.AocMain(day, Part1, Part2)
}
