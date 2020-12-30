package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type coord struct {
	x int
	y int
}

type slope struct {
	right int
	down  int
}

func readInput(fname string) string {
	data, err := ioutil.ReadFile(fname)
	if err != nil {
		panic("Cannot read input data.")
	}
	s := string(data)
	return strings.TrimRight(s, "\n")
}

func mapDimensions(txt string) (int, int) {
	lines := strings.Split(txt, "\n")
	width := len(lines[0])
	height := len(lines)
	return width, height
}

func treeLocations(txt string) []coord {
	locations := make([]coord, 0)
	lines := strings.Split(txt, "\n")
	tree := []rune("#")[0]
	for y := range lines {
		for x, val := range lines[y] {
			if val == tree {
				locations = append(locations, coord{x: x, y: y})
			}
		}
	}
	return locations
}

func contains(seq []coord, c coord) bool {
	for _, x := range seq {
		if x == c {
			return true
		}
	}
	return false
}

func part1(txt string, s slope) int {
	trees := treeLocations(txt)
	w, h := mapDimensions(txt)

	currentX := 0
	currentY := 0
	treesHit := 0

	for currentY < h {
		currentCoord := coord{x: currentX, y: currentY}
		if contains(trees, currentCoord) {
			treesHit++
		}
		currentX = (currentX + s.right) % w
		currentY += s.down
	}
	return treesHit
}

func part2(txt string) int {
	slopes := []slope{
		slope{right: 1, down: 1},
		slope{right: 3, down: 1},
		slope{right: 5, down: 1},
		slope{right: 7, down: 1},
		slope{right: 1, down: 2},
	}
	treeHits := 1
	for _, s := range slopes {
		treeHits *= part1(txt, s)
	}
	return treeHits
}

func main() {
	txt := readInput("input.txt")
	sslope := slope{right: 3, down: 1}
	p1 := part1(txt, sslope)
	fmt.Println("[PART 1] Number of trees hit:", p1)

	p2 := part2(txt)
	fmt.Println("[PART 2] Final product:", p2)
}
