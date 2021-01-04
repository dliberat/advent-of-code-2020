package main

import "testing"

func TestPart1Integration01(t *testing.T) {
	nums := []int{0, 3, 6}
	res := part1(nums)
	if res != 436 {
		t.Errorf("436 != %d", res)
	}
}

func TestPart1Integration02(t *testing.T) {
	nums := []int{1, 3, 2}
	res := part1(nums)
	if res != 1 {
		t.Errorf("1 != %d", res)
	}
}

func TestPart1Integration03(t *testing.T) {
	nums := []int{2, 1, 3}
	res := part1(nums)
	if res != 10 {
		t.Errorf("10 != %d", res)
	}
}

func TestPart1Integration04(t *testing.T) {
	nums := []int{1, 2, 3}
	res := part1(nums)
	if res != 27 {
		t.Errorf("27 != %d", res)
	}
}

func TestPart1Integration05(t *testing.T) {
	nums := []int{2, 3, 1}
	res := part1(nums)
	if res != 78 {
		t.Errorf("78 != %d", res)
	}
}

func TestPart1Integration06(t *testing.T) {
	nums := []int{3, 2, 1}
	res := part1(nums)
	if res != 438 {
		t.Errorf("438 != %d", res)
	}
}
func TestPart1Integration07(t *testing.T) {
	nums := []int{3, 1, 2}
	res := part1(nums)
	if res != 1836 {
		t.Errorf("1836 != %d", res)
	}
}
