/*
Conway's Game of Life in 3D with an infinite grid.

A cell is represented using six bits

+----------------- bits 6-7: unused
|   +------------- bits 1-5: neighbor count
|   |         +--- bit 0: cell state (1: alive, 0: dead)
|   |         |
0 0 0 0 0 0 0 0

In this representation, an off cell with no active neighbors is 0.
Since cells with no neighbors do not change and do not affect other cells,
they do not need to be stored in the hashmap.
*/
package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type coord struct {
	x int
	y int
	z int
	w int
}

func (c *coord) toString() string {
	return fmt.Sprintf("(%d, %d, %d, %d)", c.x, c.y, c.z, c.w)
}

type cellmap struct {
	cells map[coord]int
	is4d  bool // switch between 3D and 4D for parts 1 and 2 of the question
}

// setCell turns a cell on (switch from dead to alive).
// Note: The cell MUST have been dead or else neighbor counts will be incorrect
func (cm *cellmap) setCell(x, y, z, w int) {
	minW, maxW := 0, 0
	if cm.is4d {
		minW, maxW = -1, 1
	}
	for ow := minW; ow <= maxW; ow++ {
		for ox := -1; ox <= 1; ox++ {
			for oy := -1; oy <= 1; oy++ {
				for oz := -1; oz <= 1; oz++ {
					loc := coord{x + ox, y + oy, z + oz, w + ow}

					if ox == 0 && oy == 0 && oz == 0 && ow == 0 {
						cm.cells[loc] |= 1 // turn cell on
					} else {
						cm.cells[loc] += 2 // add 1 to neighbor count
					}
				}
			}
		}
	}
}

// clearCell turns a cell off (switch from alive to dead).
// Note: The cell MUST have been alive or else neighbor counts will be incorrect
func (cm *cellmap) clearCell(x, y, z, w int) {
	minW, maxW := 0, 0
	if cm.is4d {
		minW, maxW = -1, 1
	}
	for ow := minW; ow <= maxW; ow++ {
		for ox := -1; ox <= 1; ox++ {
			for oy := -1; oy <= 1; oy++ {
				for oz := -1; oz <= 1; oz++ {
					loc := coord{x + ox, y + oy, z + oz, w + ow}

					if ox == 0 && oy == 0 && oz == 0 && ow == 0 {
						cm.cells[loc]-- // turn cell off
					} else {
						cm.cells[loc] -= 2 // subtract 1 from neighbor count
					}

					if cm.cells[loc] == 0 {
						delete(cm.cells, loc)
					}
				}
			}
		}
	}
}

func (cm *cellmap) getNextIter() cellmap {
	newCm := makeCellMap(cm.is4d)

	for loc, cell := range cm.cells {
		neighborCount := cell >> 1
		if cell%2 == 1 {
			// cell is active
			if neighborCount == 2 || neighborCount == 3 {
				newCm.setCell(loc.x, loc.y, loc.z, loc.w)
			}
		} else {
			// cell is inactive
			if neighborCount == 3 {
				newCm.setCell(loc.x, loc.y, loc.z, loc.w)
			}
		}
	}

	return newCm
}

func (cm *cellmap) countActive() int {
	total := 0
	for _, cell := range cm.cells {
		total += cell % 2
	}
	return total
}

func makeCellMap(is4d bool) cellmap {
	return cellmap{cells: make(map[coord]int, 0), is4d: is4d}
}

func readInput(fname string) string {
	data, err := ioutil.ReadFile(fname)
	if err != nil {
		panic("Cannot read input data.")
	}
	s := string(data)
	return strings.TrimRight(s, "\n")
}

func input2cellmap(txt string, is4d bool) cellmap {
	cm := makeCellMap(is4d)
	split := strings.Split(txt, "\n")
	rowoffset := len(split) / 2
	coloffset := len(split[0]) / 2
	for i, row := range split {
		for j, char := range row {
			if char == '#' {
				y := i - rowoffset
				x := j - coloffset
				cm.setCell(x, y, 0, 0)
			}
		}
	}
	return cm
}

func part1(txt string) int {
	cm := input2cellmap(txt, false)

	// boot sequence = six iterations
	cm = cm.getNextIter()
	cm = cm.getNextIter()
	cm = cm.getNextIter()
	cm = cm.getNextIter()
	cm = cm.getNextIter()
	cm = cm.getNextIter()
	return cm.countActive()
}

func part2(txt string) int {
	cm := input2cellmap(txt, true)

	// boot sequence = six iterations
	cm = cm.getNextIter()
	cm = cm.getNextIter()
	cm = cm.getNextIter()
	cm = cm.getNextIter()
	cm = cm.getNextIter()
	cm = cm.getNextIter()
	return cm.countActive()
}

func main() {
	txt := readInput("input.txt")
	p1 := part1(txt)
	fmt.Println("[PART 1] Result:", p1)
	p2 := part2(txt)
	fmt.Println("[PART 2] Result:", p2)
}
