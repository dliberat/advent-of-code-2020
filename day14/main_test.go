package main

import (
	"strconv"
	"testing"
)

func TestMask01(t *testing.T) {
	txt := "XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X"
	m := makeMask(txt)

	nums := []int64{11, 101, 0}
	expected := []int64{73, 101, 64}
	for i := range nums {
		actual := m.applyPart1(nums[i])
		if actual != expected[i] {
			in := strconv.FormatInt(nums[i], 2)
			ex := strconv.FormatInt(expected[i], 2)
			ac := strconv.FormatInt(actual, 2)
			t.Errorf("Expected %s -> %s but got %s", in, ex, ac)
		}
	}
}

func TestPart1Integration01(t *testing.T) {
	txt := `mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
mem[8] = 11
mem[7] = 101
mem[8] = 0`

	res := part1(txt)
	if res != 165 {
		t.Errorf("165 != %d", res)
	}
}

func TestGetTargetAddresses(t *testing.T) {
	txt := `mask = 000000000000000000000000000000X1001X
mem[42] = 100`

	operations := parseProgramPart2(txt)
	targetAddresses := make(map[int64]int64, 0)
	for i := len(operations) - 1; i >= 0; i-- {
		addrs := operations[i].getTargetAddresses()
		for _, addr := range addrs {
			_, ok := targetAddresses[addr]
			if !ok {
				targetAddresses[addr] = operations[i].value
			}
		}
	}

	expected := []int{26, 27, 58, 59}
	for _, ex := range expected {
		val := targetAddresses[int64(ex)]
		if val != 100 {
			t.Errorf("targetAddresses[%d]: 100 != %d", ex, val)
		}
	}
}

func TestGetTargetAddresses02(t *testing.T) {
	txt := `mask = 00000000000000000000000000000000X0XX
mem[26] = 1`

	operations := parseProgramPart2(txt)
	targetAddresses := make(map[int64]int64, 0)
	for i := len(operations) - 1; i >= 0; i-- {
		addrs := operations[i].getTargetAddresses()
		for _, addr := range addrs {
			_, ok := targetAddresses[addr]
			if !ok {
				targetAddresses[addr] = operations[i].value
			}
		}
	}

	expected := []int{16, 17, 18, 19, 24, 25, 26, 27}
	for _, ex := range expected {
		val := targetAddresses[int64(ex)]
		if val != 1 {
			t.Errorf("targetAddresses[%d]: 1 != %d", ex, val)
		}
	}
}

func TestPart2Integration01(t *testing.T) {
	program := `mask = 000000000000000000000000000000X1001X
mem[42] = 100
mask = 00000000000000000000000000000000X0XX
mem[26] = 1`
	res := part2(program)
	if res != 208 {
		t.Errorf("208 != %d", res)
	}
}
