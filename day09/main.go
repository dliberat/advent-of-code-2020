package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func readInput(fname string) string {
	data, err := ioutil.ReadFile(fname)
	if err != nil {
		panic("Cannot read input data.")
	}
	s := string(data)
	return strings.TrimRight(s, "\n")
}

/*
Represent the most recent `preambleLength` numbers as a vertical column.
The horizontal rows of the array represent the result of summing the nth value
with the value at index 0, 1, 2, etc.
Because of the symmetry of the table, it is only necessary to calculate the
top right triangle.

                       |     table     |
         | values |    | 0 | 1 | 2 | 3 | ...
         +--------+    +---+---+---+---+----
         |    3   |    | 6 | 5 | 3 | 4 | ...
index -> |    2   |    | x | 4 | 2 | 2 | ...
         |    0   |    | x | x | 0 | 1 | ...
         |    1   |    | x | x | x | 2 | ...

Then, as new numbers come into the list and old numbers are evicted,
we only need to update the sums corresponding to the row and column
of the current index.
*/
type xmas struct {
	index         int
	length        int
	isInitialized bool
	values        []int
	table         [][]int
}

func makeXmas(preambleLength int) xmas {
	tbl := make([][]int, preambleLength)
	for i := range tbl {
		tbl[i] = make([]int, preambleLength)
	}

	return xmas{
		index:         0,
		length:        preambleLength,
		isInitialized: false,
		values:        make([]int, preambleLength),
		table:         tbl,
	}
}

func (x *xmas) increment() {
	x.index = (x.index + 1) % x.length
}

func (x *xmas) printTable() {
	for _, row := range x.table {
		for _, num := range row {
			fmt.Print(num, "\t")
		}
		fmt.Println("")
	}
}

func (x *xmas) readValue(val int) {
	x.values[x.index] = val

	// First, select the row at the current index, and upload
	// all values starting from the diagonal rightward
	for i := x.index; i < x.length; i++ {
		x.table[x.index][i] = x.values[x.index] + x.values[i]
	}

	// Now update the entire column at the current index from
	// the top of the table down to the current index row
	for j := 0; j <= x.index; j++ {
		x.table[j][x.index] = x.values[x.index] + x.values[j]
	}

	x.increment()
	if x.index == 0 {
		// finished reading the preamble and is ready to start processing
		x.isInitialized = true
	}
}

func (x *xmas) getValidValues() map[int]bool {
	valid := make(map[int]bool, 0)
	for i := 0; i < x.length; i++ {
		for j := i; j < x.length; j++ {
			value := x.table[i][j]
			valid[value] = true
		}
	}
	return valid
}

func (x *xmas) isValidNextNumber(n int) bool {
	if !x.isInitialized {
		return true
	}
	valid := x.getValidValues()
	return valid[n]
}

func input2slice(txt string) []int {
	lines := strings.Split(txt, "\n")
	numbers := make([]int, len(lines))
	for i := range lines {
		n, err := strconv.Atoi(lines[i])
		if err != nil {
			panic(fmt.Sprintf("Invalid input at line %d", i))
		}
		numbers[i] = n
	}
	return numbers
}

func part1(input []int, preambleLength int) int {
	x := makeXmas(preambleLength)
	for _, n := range input {
		if !x.isValidNextNumber(n) {
			return n
		}
		x.readValue(n)
	}
	return 0
}

func getSumOfMinAndMax(seq []int) int {
	min := seq[0]
	max := seq[0]
	for i := range seq {
		if seq[i] < min {
			min = seq[i]
		}
		if seq[i] > max {
			max = seq[i]
		}
	}
	return min + max
}

func part2(input []int, targetNum int) int {
	for i := 0; i < len(input); i++ {
		sum := input[i]

		if sum == targetNum {
			continue
		}

		for j := i + 1; j < len(input); j++ {
			sum += input[j]
			if sum == targetNum {
				return getSumOfMinAndMax(input[i : j+1])
			}
			if sum > targetNum {
				break
			}
		}
	}
	return 0
}

func main() {
	txt := readInput("input.txt")
	nums := input2slice(txt)
	p1 := part1(nums, 25)
	fmt.Println("[PART 1] Result:", p1)
	p2 := part2(nums, p1)
	fmt.Println("[PART 2] Result:", p2)
}
