package main

import "testing"

func TestPart1Integration01(t *testing.T) {
	nums := []int{16, 10, 15, 5, 1, 11, 7, 19, 6, 12, 4}
	nums = append(nums, 0)           // add in the airplane terminal
	nums = append(nums, max(nums)+3) // add in the mobile device
	result := part1(nums)
	if result != 7*5 {
		t.Errorf("35 != %d", result)
	}
}

func TestPart1Integration02(t *testing.T) {
	nums := []int{28, 33, 18, 42, 31, 14, 46, 20, 48, 47,
		24, 23, 49, 45, 19, 38, 39, 11, 1, 32, 25, 35, 8,
		17, 7, 9, 4, 2, 34, 10, 3}

	nums = append(nums, 0)           // add in the airplane terminal
	nums = append(nums, max(nums)+3) // add in the mobile device
	result := part1(nums)
	if result != 22*10 {
		t.Errorf("220 != %d", result)
	}
}

func TestPart2Integration02(t *testing.T) {
	nums := []int{28, 33, 18, 42, 31, 14, 46, 20, 48, 47,
		24, 23, 49, 45, 19, 38, 39, 11, 1, 32, 25, 35, 8,
		17, 7, 9, 4, 2, 34, 10, 3}

	nums = append(nums, 0)           // add in the airplane terminal
	nums = append(nums, max(nums)+3) // add in the mobile device
	result := part2(nums)
	if result != 19208 {
		t.Errorf("19208 != %d", result)
	}
}

func TestPart2Integration01(t *testing.T) {
	nums := []int{16, 10, 15, 5, 1, 11, 7, 19, 6, 12, 4}

	nums = append(nums, 0)           // add in the airplane terminal
	nums = append(nums, max(nums)+3) // add in the mobile device
	result := part2(nums)
	if result != 8 {
		t.Errorf("8 != %d", result)
	}
}
