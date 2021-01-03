package main

import "testing"

func vectorEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestMatmul01(t *testing.T) {
	eye := [][]int{
		[]int{1, 0},
		[]int{0, 1},
	}
	testVals := [][]int{
		[]int{5, 4},
		[]int{-1, 3},
		[]int{0, 44},
		[]int{1, 1},
	}
	for i, vector := range testVals {
		result := matmul(eye, vector)
		if !vectorEqual(testVals[i], result) {
			t.Errorf("%v != %v", testVals[i], result)
		}
	}
}

func TestMatMul02(t *testing.T) {
	r := [][]int{
		[]int{2, 0},
		[]int{0, 2},
	}
	testVals := [][]int{
		[]int{5, 4},
		[]int{-1, 3},
		[]int{0, 44},
		[]int{1, 1},
	}
	expected := [][]int{
		[]int{10, 8},
		[]int{-2, 6},
		[]int{0, 88},
		[]int{2, 2},
	}
	for i := range testVals {
		result := matmul(r, testVals[i])
		if !vectorEqual(expected[i], result) {
			t.Errorf("%v != %v", expected[i], result)
		}
	}
}

func TestMatMul03(t *testing.T) {
	r := [][]int{
		[]int{1, 1},
		[]int{1, -1},
	}
	testVals := [][]int{
		[]int{20, 10},
		[]int{5, 3},
		[]int{2, 7},
		[]int{0, 0},
	}
	expected := [][]int{
		[]int{30, 10},
		[]int{8, 2},
		[]int{9, -5},
		[]int{0, 0},
	}
	for i := range testVals {
		result := matmul(r, testVals[i])
		if !vectorEqual(expected[i], result) {
			t.Errorf("%v != %v", expected[i], result)
		}
	}
}

func TestRotateR90(t *testing.T) {
	testVals := [][]int{
		[]int{1, 0},  // >
		[]int{-1, 0}, // <
		[]int{0, 1},  // v
		[]int{0, -1}, // ^
	}
	expected := [][]int{
		[]int{0, 1},  // v
		[]int{0, -1}, // ^
		[]int{-1, 0}, // <
		[]int{1, 0},  // >
	}
	for i := range testVals {
		result := rotateR(testVals[i], 90)
		if !vectorEqual(expected[i], result) {
			t.Errorf("%v != %v", expected[i], result)
		}
	}
}

func TestRotateL90(t *testing.T) {
	testVals := [][]int{
		[]int{1, 0},  // >
		[]int{-1, 0}, // <
		[]int{0, 1},  // v
		[]int{0, -1}, // ^
	}
	expected := [][]int{
		[]int{0, -1}, // ^
		[]int{0, 1},  // v
		[]int{1, 0},  // >
		[]int{-1, 0}, // <
	}
	for i := range testVals {
		result := rotateL(testVals[i], 90)
		if !vectorEqual(expected[i], result) {
			t.Errorf("%v != %v", expected[i], result)
		}
	}
}

func TestRotateR270(t *testing.T) {
	testVals := [][]int{
		[]int{1, 0},  // >
		[]int{-1, 0}, // <
		[]int{0, 1},  // v
		[]int{0, -1}, // ^
	}
	expected := [][]int{
		[]int{0, -1}, // ^
		[]int{0, 1},  // v
		[]int{1, 0},  // >
		[]int{-1, 0}, // <
	}
	for i := range testVals {
		result := rotateR(testVals[i], 270)
		if !vectorEqual(expected[i], result) {
			t.Errorf("%v != %v", expected[i], result)
		}
	}
}

func TestRotateL180(t *testing.T) {
	testVals := [][]int{
		[]int{1, 0},  // >
		[]int{-1, 0}, // <
		[]int{0, 1},  // v
		[]int{0, -1}, // ^
	}
	expected := [][]int{
		[]int{-1, 0}, // ^
		[]int{1, 0},  // v
		[]int{0, -1}, // >
		[]int{0, 1},  // <
	}
	for i := range testVals {
		result := rotateR(testVals[i], 180)
		if !vectorEqual(expected[i], result) {
			t.Errorf("%v != %v", expected[i], result)
		}
	}
}

func TestWaypointMoveF(t *testing.T) {
	s := makeShip(0, 0, 10, -1)
	s.movePart2("F10")
	expectedPos := []int{100, -10}
	if !vectorEqual(expectedPos, s.pos) {
		t.Errorf("%v != %v", expectedPos, s.pos)
	}
}

func TestWaypointMove(t *testing.T) {
	s := makeShip(0, 0, 10, -1)
	s.movePart2("F10")
	s.movePart2("N3")
	s.movePart2("F7")
	expectedPos := []int{170, -38}
	if !vectorEqual(expectedPos, s.pos) {
		t.Errorf("%v != %v", expectedPos, s.pos)
	}
}

func TestPart1Integration01(t *testing.T) {
	seq := []string{
		"F10",
		"N3",
		"F7",
		"R90",
		"F11",
	}

	result := part1(seq)
	if result != 25 {
		t.Errorf("25 != %d", result)
	}
}

func TestPart2Integration01(t *testing.T) {
	seq := []string{
		"F10",
		"N3",
		"F7",
		"R90",
		"F11",
	}

	result := part2(seq)
	if result != 286 {
		t.Errorf("286 != %d", result)
	}
}
