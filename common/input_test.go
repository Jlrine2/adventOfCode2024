package common

import (
	"reflect"
	"testing"
)

func TestToLines(t *testing.T) {
	input := `line1
line2
line3`

	got := ToLines(input)
	want := []string{"line1", "line2", "line3"}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("ToLines: want: %v, got: %v", want, got)
	}
}

func TestToCollumns(t *testing.T) {
	input := `a b c
a b c
a b c`

	got := ToCollumns(input)
	want := [][]string{{"a", "a", "a"}, {"b", "b", "b"}, {"c", "c", "c"}}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("ToCollumns: want: %v, got: %v", want, got)
	}
}
