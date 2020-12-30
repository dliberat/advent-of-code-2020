package main

import "testing"

func TestTreeLocations(t *testing.T) {
	input := `..###
#.#..`

	locations := treeLocations(input)
	if len(locations) != 5 {
		t.Errorf("Expected 5 trees, got %d", len(locations))
	}
	a := coord{y: 0, x: 2}
	b := coord{y: 0, x: 3}
	c := coord{y: 0, x: 4}
	d := coord{y: 1, x: 0}
	e := coord{y: 1, x: 2}
	foundA := false
	foundB := false
	foundC := false
	foundD := false
	foundE := false
	for _, loc := range locations {
		if loc == a {
			foundA = true
		}
		if loc == b {
			foundB = true
		}
		if loc == c {
			foundC = true
		}
		if loc == d {
			foundD = true
		}
		if loc == e {
			foundE = true
		}
	}
	if !(foundA && foundB && foundC && foundD && foundE) {
		t.Error("Failed to find all trees")
	}
}

func TestPart1Integration01(t *testing.T) {
	input := `..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#`

	sslope := slope{right: 3, down: 1}
	result := part1(input, sslope)
	if result != 7 {
		t.Errorf("Expected to hit 7 trees, but got %d", result)
	}
}

func TestPart2Integration01(t *testing.T) {
	input := `..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#`

	result := part2(input)
	if result != 336 {
		t.Errorf("Expected a final product of 336 but got %d", result)
	}
}
