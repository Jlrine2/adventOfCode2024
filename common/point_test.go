package common

import "testing"

func TestPoint_Go(t *testing.T) {
	p := Point{1, 1}

	u := p.Go("up")
	want := Point{1, 0}
	if u != want {
		t.Errorf("p.Go(\"up\")=%v, want %#v", u, want)
	}
	d := p.Go("down")
	want = Point{1, 2}
	if d != want {
		t.Errorf("p.Go(\"down\")=%v, want %#v", d, want)
	}
	l := p.Go("left")
	want = Point{0, 1}
	if l != want {
		t.Errorf("p.Go(\"left\")=%v, want %#v", l, want)
	}
	r := p.Go("right")
	want = Point{2, 1}
	if r != want {
		t.Errorf("p.Go(\"right\")=%v, want %#v", r, want)
	}
	ul := p.Go("upleft")
	want = Point{0, 0}
	if ul != want {
		t.Errorf("p.Go(\"upleft\")=%v, want %#v", ul, want)
	}
	ur := p.Go("upright")
	want = Point{2, 0}
	if ur != want {
		t.Errorf("p.Go(\"upright\")=%v, want %#v", ur, want)
	}
	dl := p.Go("downleft")
	want = Point{0, 2}
	if dl != want {
		t.Errorf("p.Go(\"downleft\")=%v, want %#v", dl, want)
	}
	dr := p.Go("downright")
	want = Point{2, 2}
	if dr != want {
		t.Errorf("p.Go(\"downlight\")=%v, want %#v", dr, want)
	}
}
