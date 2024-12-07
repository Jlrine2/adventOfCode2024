package main

import (
	"adventOfCode/common"
	"slices"
	"sync"
)

const day = 6

func getStart(grid [][]int32) common.Point {
	for i, row := range grid {
		for j, _ := range row {
			if grid[i][j] == '^' {
				return common.Point{X: j, Y: i}
			}

		}
	}
	return common.Point{X: 0, Y: 0}
}

func getVisited(grid [][]int32) map[common.Point]bool {
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
	return visited
}

func Part1(input string) int {
	grid := common.ToGrid(input)
	return len(getVisited(grid))
}

func obstructionHereWouldLoop(g [][]int32, o common.Point, wouldLoop chan common.Point, wg *sync.WaitGroup) bool {
	defer wg.Done()
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
			wouldLoop <- o
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

func getObstructedRowsCols(grid [][]int32) (map[int]bool, map[int]bool) {
	oRows := make(map[int]bool)
	oCols := make(map[int]bool)
	for i, row := range grid {
		for j, cell := range row {
			if cell == '#' {
				oRows[i] = true
				oCols[j] = true
			}
		}
	}
	return oRows, oCols
}

func adjColOrRowObsructed(p common.Point, oRows map[int]bool, oCols map[int]bool) bool {
	if oRows[p.Y-1] || oRows[p.Y+1] || oCols[p.X-1] || oCols[p.X+1] {
		return true
	}
	return false
}

func Part2(input string) int {
	grid := common.ToGrid(input)
	oRows, oCols := getObstructedRowsCols(grid)
	visited := getVisited(grid)
	var wg sync.WaitGroup
	wouldLoop := make(chan common.Point, len(visited))
	for p := range visited {
		if !adjColOrRowObsructed(p, oRows, oCols) {
			continue
		}
		wg.Add(1)
		go obstructionHereWouldLoop(grid, p, wouldLoop, &wg)
	}
	wg.Wait()
	return len(wouldLoop)
}

func main() {
	common.AocMain(day, Part1, Part2)
}
