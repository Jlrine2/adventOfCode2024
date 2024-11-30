package common

import (
	"io"
	"os"
	"strings"
)

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
