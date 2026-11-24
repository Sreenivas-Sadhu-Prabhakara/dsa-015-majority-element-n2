package main

import "testing"

func TestMajority(t *testing.T) {
	cases := []struct {
		in     []int
		want   int
		wantOK bool
	}{
		{[]int{2, 2, 1, 1, 1, 2, 2}, 2, true},
		{[]int{3, 2, 3}, 3, true},
		{[]int{1}, 1, true},
		{[]int{1, 2}, 0, false},
		{[]int{1, 2, 3}, 0, false},
		{[]int{}, 0, false},
	}
	for _, c := range cases {
		v, ok := Majority(c.in)
		if ok != c.wantOK || (ok && v != c.want) {
			t.Errorf("Majority(%v) = (%d,%v), want (%d,%v)", c.in, v, ok, c.want, c.wantOK)
		}
	}
}
