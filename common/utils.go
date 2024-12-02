package common

import (
	"fmt"
	"regexp"
	"strconv"
)

func Must[T any](t T, err error) {
	if err != nil {
		panic(fmt.Errorf("%w", err))
	}
}

func Ints(s string) []int {
	r, _ := regexp.Compile(`\d+`)
	var ints []int
	for _, match := range r.FindAll([]byte(s), -1) {
		i, _ := strconv.Atoi(string(match))
		ints = append(ints, i)
	}
	return ints
}
