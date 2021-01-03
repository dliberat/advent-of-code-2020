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

// parseInputPart1 takes the puzzle input as a string and returns
// an int representing the estimated earliest time you can depart,
// and a slice of ints containing the IDs of the shuttle buses in service.
func parseInputPart1(txt string) (int, []int) {
	lines := strings.Split(txt, "\n")
	eta, err := strconv.Atoi(lines[0])
	if err != nil {
		panic("Bad input.")
	}

	ids := make([]int, 0)
	buses := strings.Split(lines[1], ",")
	for _, bus := range buses {
		id, err := strconv.Atoi(bus)
		if err == nil {
			ids = append(ids, id)
		}
	}
	return eta, ids
}

func part1(txt string) int {
	eta, ids := parseInputPart1(txt)

	for i := eta; ; i++ {
		for _, id := range ids {
			if i%id == 0 {
				timeToWait := i - eta
				return id * timeToWait
			}
		}
	}
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
