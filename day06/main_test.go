package main

import "testing"

func TestPart1Integration01(t *testing.T) {
	input := `abc

a
b
c

ab
ac

a
a
a
a

b`
	result := part1(input)
	if result != 11 {
		t.Errorf("Expected 11, got %d", result)
	}
}

func TestPart2Integration01(t *testing.T) {
	input := `abc

a
b
c

ab
ac

a
a
a
a

b`
	result := part2(input)
	if result != 6 {
		t.Errorf("Expected 6, got %d", result)
	}
}
