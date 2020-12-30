package main

import "testing"

func TestParseInput(t *testing.T) {
	input := `1721
979
366
299
675
1456`
	expected := []int{1721, 979, 366, 299, 675, 1456}
	expenseList := input2int(input)
	if len(expenseList) != len(expected) {
		t.Errorf("Expected list of length %d but got %d", len(expected), len(expenseList))
		return
	}
	for i := range expected {
		if expenseList[i] != expected[i] {
			t.Errorf("String value %d was incorrectly read as %d", expected[i], expenseList[i])
		}
	}
}

func TestPart1Integration01(t *testing.T) {
	input := `1721
979
366
299
675
1456`
	result := part1(input)
	if result != 514579 {
		t.Errorf("Expected 514579 but got %d", result)
	}
}

func TestPart2Integration01(t *testing.T) {
	input := `1721
979
366
299
675
1456`
	result := part2(input)
	if result != 241861950 {
		t.Errorf("Expected 241861950 but got %d", result)
	}
}
