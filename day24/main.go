/*
Every tile has six neighbors: east, southeast, southwest, west, northwest, and northeast.
These directions are given in your list, respectively, as e, se, sw, w, nw, and ne.

      /\ /\
     |  |  |
     /\ /\ /\
    |  |  |  |  --> x
	 \/ \/ \/
	  |  |  |
	   \/ \/
			 \
			  \
			   y

Tiles can be addressed using a 2-dimensional coordinate system.
With the axes as defined in the image above, the directions in the puzzle input
correspond to the following tiles relative to the center tile:
e  = x+1, y
w  = x-1, y
se = x, y+1
nw = x, y-1
sw = x-1, y+1
ne = x+1, y-1
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
}

func readInput(fname string) string {
	data, err := ioutil.ReadFile(fname)
	if err != nil {
		panic("Cannot read input data.")
	}
	s := string(data)
	return strings.TrimRight(s, "\n")
}

// move from the starting tile to a new tile using the directions given in string s.
// s is a sequence of directions 'e', 'w', 'se', 'sw', 'ne', or 'nw' with no delimiters.
func move(s string, start coord) coord {
	end := start
	for i := 0; i < len(s); i++ {
		if s[i] == 'e' {
			end.x++
		} else if s[i] == 'w' {
			end.x--
		} else if s[i] == 's' {
			if s[i+1] == 'e' {
				end.y++
			} else {
				end.x--
				end.y++
			}
			i++
		} else if s[i] == 'n' {
			if s[i+1] == 'e' {
				end.y--
				end.x++
			} else {
				end.y--
			}
			i++
		}
	}
	return end
}

func part1(txt string) int {
	tileset := make(map[coord]bool, 0) // false = white, true = black
	for _, line := range strings.Split(txt, "\n") {
		tile := move(line, coord{0, 0, 0})
		tileset[tile] = !tileset[tile]
	}

	// count black tiles
	total := 0
	for _, v := range tileset {
		if v {
			total++
		}
	}
	return total
}

func part2(txt string) int {
	return 0
}

func main() {
	txt := readInput("input.txt")
	p1 := part1(txt)
	fmt.Println("[PART 1] Result:", p1)
	p2 := part2(txt)
	fmt.Println("[PART 2] Result:", p2)
}
