package main

import "testing"

func TestMakeBoardingPass(t *testing.T) {
	inputs := []string{"BFFFBBFRRR", "FFFBBBFRRR", "BBFFBBFRLL"}
	for _, input := range inputs {
		bp := makeBoardingPass(input)
		output := bp.toString()
		if input != output {
			t.Errorf("Expected %s but got %s", input, output)
		}
	}
}

func TestBoardingPassRow(t *testing.T) {
	inputs := []string{"BFFFBBFRRR", "FFFBBBFRRR", "BBFFBBFRLL"}
	expected := []int{70, 14, 102}

	for i, input := range inputs {
		bp := makeBoardingPass(input)
		actual := bp.row()
		if actual != expected[i] {
			t.Errorf("Expected %d, got %d", expected[i], actual)
		}
	}
}

func TestBoardingPassCol(t *testing.T) {
	inputs := []string{"BFFFBBFRRR", "FFFBBBFRRR", "BBFFBBFRLL"}
	expected := []int{7, 7, 4}
	for i, input := range inputs {
		bp := makeBoardingPass(input)
		actual := bp.col()
		if actual != expected[i] {
			t.Errorf("Expected %d, got %d", expected[i], actual)
		}
	}
}

func TestBoardingPassSeatId(t *testing.T) {
	inputs := []string{"BFFFBBFRRR", "FFFBBBFRRR", "BBFFBBFRLL"}
	expected := []int{567, 119, 820}
	for i, input := range inputs {
		bp := makeBoardingPass(input)
		actual := bp.sid()
		if actual != expected[i] {
			t.Errorf("Expected %d, got %d", expected[i], actual)
		}
	}
}
