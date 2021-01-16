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

In part 1, we denote black/white with a boolean value. True = black, false = white.
In part 2, it becomes convenient to keep track of how many neighboring tiles are black.
Therefore, we represent each tile with an integer where bit 0 represents whether the
tile is black or white (0=white, 1=black), and bits 1-3 represent how many black neighbors
it has.

+--------- bits 1-3 = 2, so tile has two black neighbors
|     +--- bit 0 = 1, so tile is black
|     |
0 1 0 1
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

func isBlack(tile int) bool {
	return tile%2 == 1
}
func isWhite(tile int) bool {
	return tile%2 == 0
}

func incrementNeighbors(c coord, tileset *map[coord]int) {
	(*tileset)[coord{c.x + 1, c.y}] += 2     // e
	(*tileset)[coord{c.x - 1, c.y}] += 2     // w
	(*tileset)[coord{c.x, c.y + 1}] += 2     // se
	(*tileset)[coord{c.x, c.y - 1}] += 2     // nw
	(*tileset)[coord{c.x - 1, c.y + 1}] += 2 // sw
	(*tileset)[coord{c.x + 1, c.y - 1}] += 2 // ne
}

func decrementNeighbors(c coord, tileset *map[coord]int) {
	neighbors := []coord{
		coord{c.x + 1, c.y},     // e
		coord{c.x - 1, c.y},     // w
		coord{c.x, c.y + 1},     // se
		coord{c.x, c.y - 1},     // nw
		coord{c.x - 1, c.y + 1}, // sw
		coord{c.x + 1, c.y - 1}, // ne
	}
	for _, neighbor := range neighbors {
		if (*tileset)[neighbor] != 0 {
			(*tileset)[neighbor] -= 2
			if (*tileset)[neighbor] == 0 {
				delete(*tileset, neighbor)
			}
		}
	}

}

func initializeTileset(txt string) map[coord]int {
	// bit 0: 0=white, 1=black
	// bits 1-3: number of black neighbors
	tileset := make(map[coord]int, 0)
	for _, line := range strings.Split(txt, "\n") {
		tile := move(line, coord{0, 0})
		if isBlack(tileset[tile]) {
			tileset[tile]--
			decrementNeighbors(tile, &tileset)
		} else {
			tileset[tile]++
			incrementNeighbors(tile, &tileset)
		}
	}
	for c, tile := range tileset {
		if tile == 0 {
			delete(tileset, c)
		}
	}
	return tileset
}

func getNextTileset(tileset *map[coord]int) map[coord]int {
	nextTileset := make(map[coord]int, 0)

	for c, tile := range *tileset {
		nextTileset[c] = tile
	}

	for c, tile := range *tileset {
		neighborCount := tile >> 1
		if isBlack(tile) && (neighborCount == 0 || neighborCount > 2) {
			nextTileset[c]-- // tile turns white
			decrementNeighbors(c, &nextTileset)
			if nextTileset[c] == 0 {
				delete(nextTileset, c)
			}
		} else if isWhite(tile) && neighborCount == 2 {
			nextTileset[c]++ // tile turns black
			incrementNeighbors(c, &nextTileset)
		}
	}

	return nextTileset
}

func countBlackTiles(tileset *map[coord]int) int {
	total := 0
	for _, v := range *tileset {
		if isBlack(v) {
			total++
		}
	}
	return total
}

func part1(txt string) int {
	tileset := make(map[coord]bool, 0) // false = white, true = black
	for _, line := range strings.Split(txt, "\n") {
		tile := move(line, coord{0, 0})
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
	tileset := initializeTileset(txt)
	for i := 0; i < 100; i++ {
		tileset = getNextTileset(&tileset)
	}
	return countBlackTiles(&tileset)
}

func main() {
	txt := readInput("input.txt")
	p1 := part1(txt)
	fmt.Println("[PART 1] Result:", p1)
	p2 := part2(txt)
	fmt.Println("[PART 2] Result:", p2)
}
