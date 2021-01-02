package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type coord struct {
	row int
	col int
}

type seatMap struct {
	rows            [][]byte
	allVisibleSeats map[coord][]coord
}

// countOccupied returns the total number of occupied seats ("#") in the seatMap
func (s *seatMap) countOccupied() int {
	total := 0
	for _, row := range s.rows {
		for _, c := range row {
			if c == '#' {
				total++
			}
		}
	}
	return total
}

// countOccupiedNeighbors returns the total number of occupied seats ("#")
// directly adjacent (horizontal, vertical, or diagonal) to the provided location
func (s *seatMap) countOccupiedNeighbors(row, col int) int {
	total := 0
	rFrom := max(0, row-1)
	rTo := min(row+2, len(s.rows))
	cFrom := max(0, col-1)
	cTo := min(col+2, len(s.rows[0]))
	for r := rFrom; r < rTo; r++ {
		for c := cFrom; c < cTo; c++ {
			if r == row && c == col {
				continue
			}
			if len(s.rows[r]) == 0 {
				fmt.Printf("r=%d, c=%d, rows[r]:%v\n", r, c, s.rows[r])
			}
			if s.rows[r][c] == '#' {
				total++
			}
		}
	}
	return total
}

// countOccupiedInSet returns the number of occupied seats ("#")
// among the provided coordinates
func (s *seatMap) countOccupiedInSet(visible []coord) int {
	count := 0
	for _, c := range visible {
		if s.rows[c.row][c.col] == '#' {
			count++
		}
	}
	return count
}

// equal returns true if the two seat maps are equivalent for the purposes of
// determining if any seats have changed state between iterations
func (s *seatMap) equal(other *seatMap) bool {
	if len(s.rows) != len((*other).rows) {
		return false
	}
	if len(s.rows[0]) != len((*other).rows[0]) {
		return false
	}
	for i := 0; i < len(s.rows); i++ {
		for j := 0; j < len(s.rows[0]); j++ {
			if s.rows[i][j] != (*other).rows[i][j] {
				return false
			}
		}
	}
	return true
}

func (s *seatMap) getNextIter() seatMap {
	cpy := seatMap{rows: make([][]byte, len(s.rows))}
	for r := range s.rows {
		cpy.rows[r] = make([]byte, len(s.rows[r]))
		copy(cpy.rows[r], s.rows[r])
	}

	for r := 0; r < len(s.rows); r++ {
		for c := 0; c < len(s.rows[0]); c++ {
			// floor never changes.
			if s.rows[r][c] == '.' {
				continue
			}
			adjacent := s.countOccupiedNeighbors(r, c)
			// if a seat is empty (L) and there are no occupied seats
			// adjacent to it, the seat becomes occupied
			if s.rows[r][c] == 'L' && adjacent == 0 {
				cpy.rows[r][c] = '#'
			}
			// If a seat is occupied (#) and four or more seats adjacent
			// to it are also occupied, the seat becomes empty.
			if s.rows[r][c] == '#' && adjacent >= 4 {
				cpy.rows[r][c] = 'L'
			}
		}
	}

	return cpy
}

func (s *seatMap) getNextIterPart2() seatMap {
	cpy := seatMap{rows: make([][]byte, len(s.rows))}
	for r := range s.rows {
		cpy.rows[r] = make([]byte, len(s.rows[r]))
		copy(cpy.rows[r], s.rows[r])
	}
	cpy.allVisibleSeats = s.allVisibleSeats

	for r := 0; r < len(s.rows); r++ {
		for c := 0; c < len(s.rows[0]); c++ {
			// floor never changes.
			if s.rows[r][c] == '.' {
				continue
			}
			loc := coord{row: r, col: c}
			adjacent := s.countOccupiedInSet(s.allVisibleSeats[loc])
			// if a seat is empty (L) and there are no occupied seats
			// adjacent to it, the seat becomes occupied
			if s.rows[r][c] == 'L' && adjacent == 0 {
				cpy.rows[r][c] = '#'
			}
			// If a seat is occupied (#) and five or more seats adjacent
			// to it are also occupied, the seat becomes empty.
			if s.rows[r][c] == '#' && adjacent >= 5 {
				cpy.rows[r][c] = 'L'
			}
		}
	}

	return cpy
}

func (s *seatMap) findVisibleSeats(row, col int) []coord {
	visible := make([]coord, 0)

	// horizontal left
	for i := col - 1; i >= 0; i-- {
		if s.rows[row][i] != '.' {
			visible = append(visible, coord{row: row, col: i})
			break
		}
	}
	// horizontal right
	for i := col + 1; i < len(s.rows[row]); i++ {
		if s.rows[row][i] != '.' {
			visible = append(visible, coord{row: row, col: i})
			break
		}
	}
	// vertical up
	for i := row - 1; i >= 0; i-- {
		if s.rows[i][col] != '.' {
			visible = append(visible, coord{row: i, col: col})
			break
		}
	}
	// vertical down
	for i := row + 1; i < len(s.rows); i++ {
		if s.rows[i][col] != '.' {
			visible = append(visible, coord{row: i, col: col})
			break
		}
	}
	// diagonal up left
	for i := 1; row-i >= 0 && col-i >= 0; i++ {
		if s.rows[row-i][col-i] != '.' {
			visible = append(visible, coord{row: row - i, col: col - i})
			break
		}
	}
	// diagonal up right
	for i := 1; row-i >= 0 && col+i < len(s.rows[0]); i++ {
		if s.rows[row-i][col+i] != '.' {
			visible = append(visible, coord{row: row - i, col: col + i})
			break
		}
	}
	// diagonal down left
	for i := 1; row+i < len(s.rows) && col-i >= 0; i++ {
		if s.rows[row+i][col-i] != '.' {
			visible = append(visible, coord{row: row + i, col: col - i})
			break
		}
	}
	// diagonal down right
	for i := 1; row+i < len(s.rows) && col+i < len(s.rows[0]); i++ {
		if s.rows[row+i][col+i] != '.' {
			visible = append(visible, coord{row: row + i, col: col + i})
			break
		}
	}

	return visible
}

func (s *seatMap) findAllVisibleSeats() map[coord][]coord {
	s.allVisibleSeats = make(map[coord][]coord, 0)
	for i := 0; i < len(s.rows); i++ {
		for j := 0; j < len(s.rows[0]); j++ {
			c := coord{row: i, col: j}
			s.allVisibleSeats[c] = s.findVisibleSeats(i, j)
		}
	}
	return s.allVisibleSeats
}

func (s *seatMap) print() {
	for _, row := range s.rows {
		fmt.Println(string(row))
	}
}

func makeSeatMap(txt string) seatMap {
	s := seatMap{}
	lines := strings.Split(txt, "\n")
	s.rows = make([][]byte, len(lines))
	for i, line := range lines {
		s.rows[i] = []byte(line)
	}
	return s
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func readInput(fname string) string {
	data, err := ioutil.ReadFile(fname)
	if err != nil {
		panic("Cannot read input data.")
	}
	s := string(data)
	return strings.TrimRight(s, "\n")
}

func part1(txt string) int {
	seatingMap := makeSeatMap(txt)
	for true {
		next := seatingMap.getNextIter()
		if seatingMap.equal(&next) {
			return seatingMap.countOccupied()
		}
		seatingMap = next
	}
	return -1
}

func part2(txt string) int {
	seatingMap := makeSeatMap(txt)
	seatingMap.findAllVisibleSeats()
	for true {
		next := seatingMap.getNextIterPart2()
		if seatingMap.equal(&next) {
			return seatingMap.countOccupied()
		}
		seatingMap = next
	}
	return -1
}

func main() {
	txt := readInput("input.txt")
	p1 := part1(txt)
	fmt.Println("[PART 1] Result:", p1)
	p2 := part2(txt)
	fmt.Println("[PART 2] Result:", p2)
}
