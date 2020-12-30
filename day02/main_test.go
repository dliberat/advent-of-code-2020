package main

import "testing"

func TestPart1Integration01(t *testing.T) {
	input := `1-3 a: abcde
1-3 b: cdefg
2-9 c: ccccccccc`

	result := part1(input)
	if result != 2 {
		t.Errorf("Expected 2 valid passwords but found %d", result)
	}
}

func TestPart2Integration01(t *testing.T) {
	input := `1-3 a: abcde
1-3 b: cdefg
2-9 c: ccccccccc`

	result := part2(input)
	if result != 1 {
		t.Errorf("Expected 1 valid password but found %d", result)
	}
}
