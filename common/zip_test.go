package common

import (
	"reflect"
	"testing"
)

func TestZip(t *testing.T) {
	a := []int{1, 3, 5}
	b := []int{2, 4, 6}

	got := Zip(a, b)
	want := []int{1, 2, 3, 4, 5, 6}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Zip %#v, %#v, Got %#v, Want %#v", a, b, got, want)
	}
}

func TestUnzip(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6}

	got := Unzip(a, 2)
	want := [][]int{{1, 3, 5}, {2, 4, 6}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Unip %#v, Got %#v, Want %#v", a, got, want)
	}

	b := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	got = Unzip(b, 3)
	want = [][]int{{1, 4, 7}, {2, 5, 8}, {3, 6, 9}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Unip %#v, Got %#v, Want %#v", b, got, want)
	}

	c := []int{1, 2, 3}

	got = Unzip(c, 1)
	want = [][]int{{1, 2, 3}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Unip %#v, Got %#v, Want %#v", b, got, want)
	}
}

func TestZipUnzip(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	got := Zip(Unzip(a, 3)...)
	if !reflect.DeepEqual(got, a) {
		t.Errorf("ZipUnzip %#v, Got %#v, Want %#v", a, got, a)
	}
}

func TestUnzipZip(t *testing.T) {
	a := [][]int{{1, 2, 3}, {4, 5, 6}}

	got := Unzip(Zip(a...), 2)
	if !reflect.DeepEqual(got, a) {
		t.Errorf("ZipUnzip %#v, Got %#v, Want %#v", a, got, a)
	}
}
