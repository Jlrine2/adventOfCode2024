package main

import (
	"adventOfCode/common"
	"slices"
)

const day = 6

func getStart(g [][]int32) common.Point {
	for i, row := range g {
		for j, _ := range row {
			if g[i][j] == '^' {
				return common.Point{X: j, Y: i}
			}

		}
	}
	return common.Point{X: 0, Y: 0}
}

func Part1(input string) int {
	grid := common.ToGrid(input)
	visited := make(map[common.Point]bool)
	p := getStart(grid)
	directions := []string{"up", "right", "down", "left"}
	currentDirection := 0
	for true {
		_, inBounds := p.Get(grid)
		if !inBounds {
			break
		}
		visited[p] = true
		next := p.Go(directions[currentDirection])
		nextValue, _ := next.Get(grid)
		if nextValue == '#' {
			currentDirection++
			currentDirection = currentDirection % 4
		} else {
			p = next
		}
	}
	return len(visited)
}

func obstructionHereWouldLoop(g [][]int32, o common.Point) bool {
	visited := make(map[common.Point][]string)
	p := getStart(g)
	directions := []string{"up", "right", "down", "left"}
	currentDirection := 0
	for true {
		_, inBounds := p.Get(g)
		if !inBounds {
			return false
		}
		next := p.Go(directions[currentDirection])
		nextValue, _ := next.Get(g)
		if slices.Contains(visited[p], directions[currentDirection]) {
			return true
		}
		visited[p] = append(visited[p], directions[currentDirection])
		if nextValue == '#' || next == o {
			currentDirection++
			currentDirection = currentDirection % 4
		} else {
			p = next
		}
	}
	return false
}

func Part2(input string) int {
	grid := common.ToGrid(input)
	s := 0
	for i, row := range grid {
		for j, _ := range row {
			if grid[i][j] == '^' {
				continue
			}
			if obstructionHereWouldLoop(grid, common.Point{X: j, Y: i}) {
				s++
			}
		}
	}
	return s
}

func main() {
	common.AocMain(day, Part1, Part2)
}
