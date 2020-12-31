package main

import (
	"fmt"
	"io/ioutil"
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

func countAnyYeses(txt string) map[byte]int {
	counts := make(map[byte]int, 0)

	individuals := strings.Split(txt, "\n")
	for _, individual := range individuals {
		for i := 0; i < len(individual); i++ {
			char := individual[i]
			counts[char] = counts[char] + 1
		}
	}
	return counts
}

func countAllYeses(txt string) int {
	counts := make(map[byte]int, 0)

	individuals := strings.Split(txt, "\n")
	for _, individual := range individuals {
		for i := 0; i < len(individual); i++ {
			char := individual[i]
			counts[char] = counts[char] + 1
		}
	}

	allYeses := 0
	for _, val := range counts {
		if val == len(individuals) {
			allYeses++
		}
	}
	return allYeses
}

func part1(txt string) int {
	groups := strings.Split(txt, "\n\n")
	total := 0
	for _, group := range groups {
		counts := countAnyYeses(group) // how many times each question got a yes answer
		total += len(counts)
	}
	return total
}

func part2(txt string) int {
	groups := strings.Split(txt, "\n\n")
	total := 0
	for _, group := range groups {
		cnt := countAllYeses(group) // how many times each question got a yes answer
		total += cnt
	}
	return total
}

func main() {
	txt := readInput("input.txt")
	p1 := part1(txt)
	fmt.Println("[PART 1] Result:", p1)
	p2 := part2(txt)
	fmt.Println("[PART 2] Result:", p2)
}
