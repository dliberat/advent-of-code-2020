package main

import (
	"fmt"
	"testing"
)

func TestCountOccupied(t *testing.T) {
	txt := `#.##
#L##
L.#.
#L#.`
	seats := makeSeatMap(txt)
	cnt := seats.countOccupied()
	if cnt != 9 {
		t.Errorf("9 != %d", cnt)
	}
}

func TestCountOccupied02(t *testing.T) {
	txt := `#.##..
#L##..
L.#..#
#L#...`
	seats := makeSeatMap(txt)
	cnt := seats.countOccupied()
	if cnt != 10 {
		t.Errorf("10 != %d", cnt)
	}
}

func TestCountOccupied03(t *testing.T) {
	txt := `#.L#.L#.L#
#LLLLLL.LL
L.L.L..#..
##L#.#L.L#
L.L#.LL.L#
#.LLLL#.LL
..#.L.....
LLL###LLL#
#.LLLLL#.L
#.L#LL#.L#`
	seats := makeSeatMap(txt)
	cnt := seats.countOccupied()
	if cnt != 26 {
		t.Errorf("26 != %d", cnt)
	}
}

func TestCountOccupiedNeighbors01(t *testing.T) {
	txt := `#.##
#L##
L.#.
#L#.`
	seats := makeSeatMap(txt)
	expected := [][]int{
		[]int{1, 4, 3, 3},
		[]int{1, 5, 4, 4},
		[]int{2, 5, 3, 4},
		[]int{0, 3, 1, 2},
	}
	for i, row := range expected {
		for j, val := range row {
			actual := seats.countOccupiedNeighbors(i, j)
			if actual != val {
				t.Errorf("(%d, %d) %d != %d", i, j, val, actual)
			}
		}
	}
}
func TestCountOccupiedNeighbors02(t *testing.T) {
	txt := `L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL`
	seats := makeSeatMap(txt)
	for i, row := range seats.rows {
		for j := range row {
			actual := seats.countOccupiedNeighbors(i, j)
			if actual != 0 {
				t.Errorf("(%d, %d) 0 != %d", i, j, actual)
			}
		}
	}
}
func TestCountOccupiedNeighbors03(t *testing.T) {
	txt := `#.##..
#L##..
L.#...
#L#...`
	seats := makeSeatMap(txt)
	expected := [][]int{
		[]int{1, 4, 3, 3, 2, 0},
		[]int{1, 5, 4, 4, 2, 0},
		[]int{2, 5, 3, 4, 1, 0},
		[]int{0, 3, 1, 2, 0, 0},
	}
	for i, row := range expected {
		for j, val := range row {
			actual := seats.countOccupiedNeighbors(i, j)
			if actual != val {
				t.Errorf("(%d, %d) %d != %d", i, j, val, actual)
			}
		}
	}
}

func TestSeatMapEqual(t *testing.T) {
	txtA := `#.##
#L##
L.#.
#L#.`
	txtB := `#.##
#L##
L.#.
#L#.`
	txtC := `#.##
#L##
L...
#L#.`
	a := makeSeatMap(txtA)
	b := makeSeatMap(txtB)
	c := makeSeatMap(txtC)
	if !a.equal(&b) {
		t.Error("a should be equal to b")
		return
	}
	if !b.equal(&a) {
		t.Error("b should be equal to a")
		return
	}
	if a.equal(&c) {
		t.Error("a should not be equal to c (difference in line 3 of 4)")
		return
	}
	if b.equal(&c) {
		t.Error("a should not be equal to c (difference in line 3 of 4)")
		return
	}
	if c.equal(&a) {
		t.Error("c should not be equal to a (difference in line 3 of 4)")
		return
	}
	if c.equal(&b) {
		t.Error("c should not be equal to b (difference in line 3 of 4)")
		return
	}
}

func TestGetNextIter01(t *testing.T) {
	txt := `L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL`
	expectedTxt := `#.##.##.##
#######.##
#.#.#..#..
####.##.##
#.##.##.##
#.#####.##
..#.#.....
##########
#.######.#
#.#####.##`
	initialState := makeSeatMap(txt)
	nextIter := initialState.getNextIter()
	expected := makeSeatMap(expectedTxt)
	if !nextIter.equal(&expected) {
		t.Error("Next iteration does not match the expected seat layout.")

		for r := 0; r < len(expected.rows); r++ {
			for c := 0; c < len(expected.rows[0]); c++ {
				fmt.Print(string(nextIter.rows[r][c]))
				if nextIter.rows[r][c] != expected.rows[r][c] {
					fmt.Println("")
					return
				}
			}
		}
		fmt.Println("")
	}
}

func TestGetVisibleSeats01(t *testing.T) {
	txt := `.......#.
...#.....
.#.......
.........
..#L....#
....#....
.........
#........
...#.....`
	seats := makeSeatMap(txt)
	visible := seats.findVisibleSeats(4, 3)
	if len(visible) != 8 {
		t.Errorf("8 != %d", len(visible))
		fmt.Println(visible)
	}
}

func TestGetVisibleSeats02(t *testing.T) {
	txt := `.............
.L.L.#.#.#.#.
.............`
	seats := makeSeatMap(txt)
	visible := seats.findVisibleSeats(1, 1)
	if len(visible) != 1 {
		t.Errorf("1 != %d", len(visible))
	}
	neighbor := visible[0]
	if neighbor.row != 1 || neighbor.col != 3 {
		t.Errorf("Expected neighbor at (1,3) but got (%d,%d)", neighbor.row, neighbor.col)
	}
}

func TestGetVisibleSeats03(t *testing.T) {
	txt := `.##.##.
#.#.#.#
##...##
...L...
##...##
#.#.#.#
.##.##.`
	seats := makeSeatMap(txt)
	visible := seats.findVisibleSeats(3, 3)
	if len(visible) != 0 {
		t.Errorf("0 != %d", len(visible))
		fmt.Println(visible)
	}
}

func TestGetVisibleSeats04(t *testing.T) {
	txt := `#.##.##.##
#######.##
#.#.#..#..
####.##.##
#.##.##.##
#.#####.##
..#.#.....
##########
#.######.#
#.#####.##`
	seats := makeSeatMap(txt)
	visible := seats.findVisibleSeats(0, 2)
	if len(visible) != 5 {
		t.Errorf("5 != %d", len(visible))
		fmt.Println(visible)
	}
}

func TestPart1Integration01(t *testing.T) {
	txt := `L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL`
	result := part1(txt)
	if result != 37 {
		t.Errorf("37 != %d", result)
	}
}

func TestGetNextIterPart201(t *testing.T) {
	txt := `L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL`
	expectedTxt := `#.##.##.##
#######.##
#.#.#..#..
####.##.##
#.##.##.##
#.#####.##
..#.#.....
##########
#.######.#
#.#####.##`
	initialState := makeSeatMap(txt)
	initialState.findAllVisibleSeats()
	nextIter := initialState.getNextIterPart2()
	expected := makeSeatMap(expectedTxt)
	if !nextIter.equal(&expected) {
		t.Error("Next iteration does not match the expected seat layout.")

		for r := 0; r < len(expected.rows); r++ {
			for c := 0; c < len(expected.rows[0]); c++ {
				fmt.Print(string(nextIter.rows[r][c]))
				if nextIter.rows[r][c] != expected.rows[r][c] {
					fmt.Println("")
					return
				}
			}
		}
		fmt.Println("")
	}
}

func TestGetNextIterPart202(t *testing.T) {
	txt := `#.##.##.##
#######.##
#.#.#..#..
####.##.##
#.##.##.##
#.#####.##
..#.#.....
##########
#.######.#
#.#####.##`
	expectedTxt := `#.LL.LL.L#
#LLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLL#
#.LLLLLL.L
#.LLLLL.L#`
	initialState := makeSeatMap(txt)
	initialState.findAllVisibleSeats()
	nextIter := initialState.getNextIterPart2()
	expected := makeSeatMap(expectedTxt)
	if !nextIter.equal(&expected) {
		t.Error("Next iteration does not match the expected seat layout.")

		for r := 0; r < len(expected.rows); r++ {
			for c := 0; c < len(expected.rows[0]); c++ {
				fmt.Print(string(nextIter.rows[r][c]))
				if nextIter.rows[r][c] != expected.rows[r][c] {
					fmt.Println("")
					return
				}
			}
		}
		fmt.Println("")
	}
}

func TestGetNextIterPart203(t *testing.T) {
	txt := `#.LL.LL.L#
#LLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLL#
#.LLLLLL.L
#.LLLLL.L#`
	expectedTxt := `#.L#.##.L#
#L#####.LL
L.#.#..#..
##L#.##.##
#.##.#L.##
#.#####.#L
..#.#.....
LLL####LL#
#.L#####.L
#.L####.L#`
	initialState := makeSeatMap(txt)
	initialState.findAllVisibleSeats()
	nextIter := initialState.getNextIterPart2()
	expected := makeSeatMap(expectedTxt)
	if !nextIter.equal(&expected) {
		t.Error("Next iteration does not match the expected seat layout.")

		for r := 0; r < len(expected.rows); r++ {
			for c := 0; c < len(expected.rows[0]); c++ {
				fmt.Print(string(nextIter.rows[r][c]))
				if nextIter.rows[r][c] != expected.rows[r][c] {
					fmt.Println("")
					return
				}
			}
		}
		fmt.Println("")
	}
}

func TestGetNextIterPart204(t *testing.T) {
	txt := `#.L#.##.L#
#L#####.LL
L.#.#..#..
##L#.##.##
#.##.#L.##
#.#####.#L
..#.#.....
LLL####LL#
#.L#####.L
#.L####.L#`
	expectedTxt := `#.L#.L#.L#
#LLLLLL.LL
L.L.L..#..
##LL.LL.L#
L.LL.LL.L#
#.LLLLL.LL
..L.L.....
LLLLLLLLL#
#.LLLLL#.L
#.L#LL#.L#`
	initialState := makeSeatMap(txt)
	initialState.findAllVisibleSeats()
	nextIter := initialState.getNextIterPart2()
	expected := makeSeatMap(expectedTxt)
	if !nextIter.equal(&expected) {
		t.Error("Next iteration does not match the expected seat layout.")

		for r := 0; r < len(expected.rows); r++ {
			for c := 0; c < len(expected.rows[0]); c++ {
				fmt.Print(string(nextIter.rows[r][c]))
				if nextIter.rows[r][c] != expected.rows[r][c] {
					fmt.Println("")
					return
				}
			}
		}
		fmt.Println("")
	}
}

func TestGetNextIterPart205(t *testing.T) {
	txt := `#.L#.L#.L#
#LLLLLL.LL
L.L.L..#..
##LL.LL.L#
L.LL.LL.L#
#.LLLLL.LL
..L.L.....
LLLLLLLLL#
#.LLLLL#.L
#.L#LL#.L#`
	expectedTxt := `#.L#.L#.L#
#LLLLLL.LL
L.L.L..#..
##L#.#L.L#
L.L#.#L.L#
#.L####.LL
..#.#.....
LLL###LLL#
#.LLLLL#.L
#.L#LL#.L#`
	initialState := makeSeatMap(txt)
	initialState.findAllVisibleSeats()
	nextIter := initialState.getNextIterPart2()
	expected := makeSeatMap(expectedTxt)
	if !nextIter.equal(&expected) {
		t.Error("Next iteration does not match the expected seat layout.")

		for r := 0; r < len(expected.rows); r++ {
			for c := 0; c < len(expected.rows[0]); c++ {
				fmt.Print(string(nextIter.rows[r][c]))
				if nextIter.rows[r][c] != expected.rows[r][c] {
					fmt.Println("")
					return
				}
			}
		}
		fmt.Println("")
	}
}

func TestGetNextIterPart206(t *testing.T) {
	txt := `#.L#.L#.L#
#LLLLLL.LL
L.L.L..#..
##L#.#L.L#
L.L#.#L.L#
#.L####.LL
..#.#.....
LLL###LLL#
#.LLLLL#.L
#.L#LL#.L#`
	expectedTxt := `#.L#.L#.L#
#LLLLLL.LL
L.L.L..#..
##L#.#L.L#
L.L#.LL.L#
#.LLLL#.LL
..#.L.....
LLL###LLL#
#.LLLLL#.L
#.L#LL#.L#`
	initialState := makeSeatMap(txt)
	initialState.findAllVisibleSeats()
	nextIter := initialState.getNextIterPart2()
	expected := makeSeatMap(expectedTxt)
	if !nextIter.equal(&expected) {
		t.Error("Next iteration does not match the expected seat layout.")

		for r := 0; r < len(expected.rows); r++ {
			for c := 0; c < len(expected.rows[0]); c++ {
				fmt.Print(string(nextIter.rows[r][c]))
				if nextIter.rows[r][c] != expected.rows[r][c] {
					fmt.Println("")
					return
				}
			}
		}
		fmt.Println("")
	}
}

func TestPart2Integration01(t *testing.T) {
	txt := `L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL`
	result := part2(txt)
	if result != 26 {
		t.Errorf("26 != %d", result)
	}
}
