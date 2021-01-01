package main

import (
	"testing"
)

func TestXmasReadValue(t *testing.T) {
	preambleLength := 5
	x := makeXmas(preambleLength)
	values := []int{0, 1, 2, 3, 4}
	for _, val := range values {
		x.readValue(val)
	}
	/*
		    0 1 2 3 4
		  +----------
		0 | 0 1 2 3 4
		1 |   2 3 4 5
		2 |     4 5 6
		3 |       6 7
		4 |         8
	*/
	expected := [][]int{
		[]int{0, 1, 2, 3, 4},
		[]int{0, 2, 3, 4, 5},
		[]int{0, 0, 4, 5, 6},
		[]int{0, 0, 0, 6, 7},
		[]int{0, 0, 0, 0, 8},
	}
	for i := 0; i < preambleLength; i++ {
		for j := 0; j < preambleLength; j++ {
			if x.table[i][j] != expected[i][j] {
				t.Error("Incorrect table setup after preamble.")
				x.printTable()
			}
		}
	}
}

func TestXmasReadValue02(t *testing.T) {
	preambleLength := 2
	x := makeXmas(preambleLength)
	values := []int{0, 1, 2, 3, 4}
	for _, val := range values {
		x.readValue(val)
	}
	/*
		After first 2 numbers (0, 1)
			    0 1
			  +----
		->  0 | 0 1
			1 |   2

		After third number (2)
			    2 1
			  +----
			2 | 4 3
		->	1 |   2

		After fourth number(3)
			    2 3
			  +----
		->  2 | 4 5
			3 |   6

		After fifth number (4)
			    4 3
			  +----
			4 | 8 7
		->	3 |   6
	*/
	expected := [][]int{
		[]int{8, 7},
		[]int{0, 6},
	}
	for i := 0; i < preambleLength; i++ {
		for j := 0; j < preambleLength; j++ {
			if x.table[i][j] != expected[i][j] {
				t.Error("Incorrect table setup")
				x.printTable()
			}
		}
	}
}

func TestPart1Integration01(t *testing.T) {
	input := []int{35, 20, 15, 25, 47, 40, 62, 55, 65, 95, 102, 117, 150, 182, 127, 219, 299, 277, 309, 576}
	preambleLength := 5
	res := part1(input, preambleLength)
	if res != 127 {
		t.Errorf("127 != %d", res)
	}
}

func TestPart2Integration01(t *testing.T) {
	input := []int{35, 20, 15, 25, 47, 40, 62, 55, 65, 95, 102, 117, 150, 182, 127, 219, 299, 277, 309, 576}
	res := part2(input, 127)
	if res != 62 {
		t.Errorf("62 != %d", res)
	}
}
