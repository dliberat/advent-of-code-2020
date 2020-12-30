package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type rule struct {
	min  int
	max  int
	char byte
}

type dbEntry struct {
	r  rule
	pw string
}

func (entry *dbEntry) isValidPart1() bool {
	countDict := countChars(entry.pw)
	count := countDict[entry.r.char]
	return count >= entry.r.min && count <= entry.r.max
}

func (entry *dbEntry) isValidPart2() bool {
	charAtMin := entry.pw[entry.r.min-1]
	charAtMax := entry.pw[entry.r.max-1]
	charAtMinMatches := charAtMin == entry.r.char
	charAtMaxMatches := charAtMax == entry.r.char
	return charAtMinMatches != charAtMaxMatches // XOR
}

// makeRule takes the "rule" portion of the database entry
// (everything up until the ":") and returns the corresponding rule struct
func makeRule(txt string) rule {
	split := strings.Split(txt, " ")
	nums, char := split[0], split[1]
	splitNums := strings.Split(nums, "-")
	min, err := strconv.Atoi(splitNums[0])
	if err != nil {
		panic("Cannot convert string to number")
	}
	max, err := strconv.Atoi(splitNums[1])
	if err != nil {
		panic("Cannot convert string to number")
	}

	return rule{min: min, max: max, char: char[0]}
}

func readInput(fname string) string {
	data, err := ioutil.ReadFile(fname)
	if err != nil {
		panic("Cannot read input data.")
	}
	s := string(data)
	return strings.TrimRight(s, "\n")
}

func input2database(database string) []dbEntry {
	entries := strings.Split(database, "\n")
	db := make([]dbEntry, len(entries))
	for i, line := range entries {
		split := strings.Split(line, ": ")
		left, pw := split[0], split[1]
		r := makeRule(left)
		db[i] = dbEntry{r: r, pw: pw}
	}
	return db
}

func countChars(s string) map[byte]int {
	count := make(map[byte]int, 0)
	for i := range s {
		b := s[i]
		count[b] = count[b] + 1
	}
	return count
}

func part1(database string) int {
	db := input2database(database)
	validCounter := 0
	for _, entry := range db {
		if entry.isValidPart1() {
			validCounter++
		}
	}
	return validCounter
}

func part2(database string) int {
	db := input2database(database)
	validCounter := 0
	for _, entry := range db {
		if entry.isValidPart2() {
			validCounter++
		}
	}
	return validCounter
}

func main() {
	txt := readInput("input.txt")
	p1 := part1(txt)
	fmt.Println("[PART 1] There are", p1, "valid passwords in the database.")

	p2 := part2(txt)
	fmt.Println("[PART 2] There are", p2, "valid passwords in the database.")
}
