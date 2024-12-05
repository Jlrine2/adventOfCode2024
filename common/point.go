package common

import "strings"

func ToGrid(s string) [][]int32 {
	lines := ToLines(s)
	grid := make([][]int32, len(lines))
	for i, line := range lines {
		grid[i] = make([]int32, len(line))
		for j, cell := range line {
			grid[i][j] = cell
		}
	}
	return grid
}

type Point struct {
	X int
	Y int
}

func (p Point) Go(direction string) Point {
	x, y := p.X, p.Y
	if strings.Contains(direction, "up") {
		y--
	}
	if strings.Contains(direction, "down") {
		y++
	}
	if strings.Contains(direction, "left") {
		x--
	}
	if strings.Contains(direction, "right") {
		x++
	}
	return Point{x, y}
}

func (p *Point) Get(grid [][]int32) (int32, bool) {
	if p.X < 0 || p.X >= len(grid) {
		return 0, false
	}
	if p.Y < 0 || p.Y >= len(grid[p.X]) {
		return 0, false
	}
	return grid[p.X][p.Y], true
}
