package common

import (
	"reflect"
	"testing"
)

func TestInts(t *testing.T) {
	s := "ab1cd23egh456"

	ints := Ints(s)
	want := []int{1, 23, 456}
	if !reflect.DeepEqual(ints, want) {
		t.Errorf("Ints: got %#v want %#v", ints, want)
	}

	s = "ab-12cde23-47bda"

	ints = Ints(s)
	want = []int{-12, 23, -47}
	if !reflect.DeepEqual(ints, want) {
		t.Errorf("Ints: got %#v want %#v", ints, want)
	}
}
