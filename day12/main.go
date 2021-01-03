package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type ship struct {
	pos      []int
	waypoint []int
}

// makeShip creates a ship located at position (x, y) with a waypoint at (wx, wy)
func makeShip(x, y, wx, wy int) ship {
	// The waypoint starts 10 units east and 1 unit north relative to the ship.
	return ship{
		pos:      []int{x, y},
		waypoint: []int{wx, wy},
	}
}

// movePart1 is the movement logic for Part 1 of the puzzle
func (s *ship) movePart1(instr string) {
	action := instr[0]
	value, err := strconv.Atoi(instr[1:])
	if err != nil {
		panic("bad instruction: " + instr)
	}
	if action == 'N' {
		scaled := scale([]int{0, -1}, value)
		s.pos = add(s.pos, scaled)
	} else if action == 'S' {
		scaled := scale([]int{0, 1}, value)
		s.pos = add(s.pos, scaled)
	} else if action == 'W' {
		scaled := scale([]int{-1, 0}, value)
		s.pos = add(s.pos, scaled)
	} else if action == 'E' {
		scaled := scale([]int{1, 0}, value)
		s.pos = add(s.pos, scaled)
	} else if action == 'L' {
		s.waypoint = rotateL(s.waypoint, value)
	} else if action == 'R' {
		s.waypoint = rotateR(s.waypoint, value)
	} else if action == 'F' {
		scaled := scale(s.waypoint, value)
		s.pos = add(s.pos, scaled)
	}
}

// movePart2 is the movement logic for Part 2 of the puzzle
func (s *ship) movePart2(instr string) {
	action := instr[0]
	value, err := strconv.Atoi(instr[1:])
	if err != nil {
		panic("bad instruction: " + instr)
	}
	if action == 'N' {
		s.waypoint = moveN(s.waypoint, value)
	} else if action == 'S' {
		s.waypoint = moveS(s.waypoint, value)
	} else if action == 'W' {
		s.waypoint = moveW(s.waypoint, value)
	} else if action == 'E' {
		s.waypoint = moveE(s.waypoint, value)
	} else if action == 'L' {
		s.waypoint = rotateL(s.waypoint, value)
	} else if action == 'R' {
		s.waypoint = rotateR(s.waypoint, value)
	} else if action == 'F' {
		scaled := scale(s.waypoint, value)
		s.pos = add(s.pos, scaled)
	}
}

func readInput(fname string) string {
	data, err := ioutil.ReadFile(fname)
	if err != nil {
		panic("Cannot read input data.")
	}
	s := string(data)
	return strings.TrimRight(s, "\n")
}

// matmul multiplies vector v on the left side by matrix r
func matmul(r [][]int, v []int) []int {
	if len(r[0]) != len(v) {
		panic(fmt.Sprintf("incompatible matrix and vector sizes %d and %d", len(r[0]), len(v)))
	}
	output := make([]int, len(r))
	for i, row := range r {
		for j := range row {
			output[i] += row[j] * v[j]
		}
	}
	return output
}

// rotateR rotates the vector v by theta degrees clockwise,
// where theta is one of 90, 180, or 270
func rotateR(v []int, theta int) []int {
	var r [][]int
	if theta == 90 {
		r = [][]int{
			[]int{0, -1},
			[]int{1, 0},
		}
	} else if theta == 180 {
		r = [][]int{
			[]int{-1, 0},
			[]int{0, -1},
		}
	} else if theta == 270 {
		r = [][]int{
			[]int{0, 1},
			[]int{-1, 0},
		}
	} else {
		panic("not implemented")
	}
	return matmul(r, v)
}

// rotateL rotates the vector v by theta degrees counterclockwise,
// where theta is one of 90, 180, or 270
func rotateL(v []int, theta int) []int {
	if theta == 90 {
		return rotateR(v, 270)
	} else if theta == 270 {
		return rotateR(v, 90)
	} else {
		return rotateR(v, theta)
	}
}

// add vectors a and b
func add(a, b []int) []int {
	res := make([]int, len(a))
	for i := range a {
		res[i] = a[i] + b[i]
	}
	return res
}

// scale the vector v by x
func scale(v []int, x int) []int {
	y := make([]int, len(v))
	for i := range v {
		y[i] = v[i] * x
	}
	return y
}

func moveN(v []int, dist int) []int {
	x := []int{0, -dist}
	return add(v, x)
}
func moveS(v []int, dist int) []int {
	x := []int{0, dist}
	return add(v, x)
}
func moveW(v []int, dist int) []int {
	x := []int{-dist, 0}
	return add(v, x)
}
func moveE(v []int, dist int) []int {
	x := []int{dist, 0}
	return add(v, x)
}
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func part1(instructions []string) int {
	// The ship starts by facing east.
	s := makeShip(0, 0, 1, 0)
	for _, instr := range instructions {
		s.movePart1(instr)
	}
	return abs(s.pos[0]) + abs(s.pos[1])
}

func part2(instructions []string) int {
	// The waypoint starts 10 units east and 1 unit north relative to the ship.
	s := makeShip(0, 0, 10, -1)
	for _, instr := range instructions {
		s.movePart2(instr)
	}
	return abs(s.pos[0]) + abs(s.pos[1])
}

func main() {
	txt := readInput("input.txt")
	lines := strings.Split(txt, "\n")
	p1 := part1(lines)
	fmt.Println("[PART 1] Result:", p1)
	p2 := part2(lines)
	fmt.Println("[PART 2] Result:", p2)
}
