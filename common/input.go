package common

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func AocMain(day int, part1 func(string) int, part2 func(string) int) {
	input := ReadInput(fmt.Sprintf("days/%d/input.txt", day))
	p1start := time.Now()
	p1 := part1(input)
	p1elapsed := time.Since(p1start)

	p2start := time.Now()
	p2 := part2(input)
	p2elapsed := time.Since(p2start)

	fmt.Print("---RESULTS---\n\n")
	fmt.Printf("Part1:\nsolution: %#v\ntime: %#v\n\n", p1, p1elapsed.String())
	fmt.Printf("Part2:\nsolution: %#v\ntime: %#v\n", p2, p2elapsed.String())
}

func ReadInput(path string) string {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			panic(err)
		}
	}(f)
	data, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func ToLines(s string) []string {
	return strings.Split(strings.ReplaceAll(s, "\r\n", "\n"), "\n")
}

func ToCollumns(s string) [][]string {
	lines := ToLines(s)
	numRows := len(strings.Fields(lines[0]))
	columns := make([][]string, numRows, numRows)
	for _, l := range lines {
		f := strings.Fields(l)
		for i, c := range f {
			columns[i] = append(columns[i], c)
		}
	}
	return columns
}
