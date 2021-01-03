package main

import "testing"

func TestPart1Integration01(t *testing.T) {
	txt := `939
7,13,x,x,59,x,31,19`

	res := part1(txt)
	if res != 295 {
		t.Errorf("295 != %d", res)
	}
}
