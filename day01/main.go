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

func input2int(expenses string) []int {
	s := strings.Split(expenses, "\n")
	expenseList := make([]int, len(s))
	for i, str := range s {
		val, err := strconv.Atoi(str)
		if err != nil {
			panic(fmt.Sprintf("Bad input at line %d", i))
		}
		expenseList[i] = val
	}
	return expenseList
}

func part1(expenses string) int {
	expenseList := input2int(expenses)
	diffs := make(map[int]int, 0)
	for _, expense := range expenseList {
		diff := 2020 - expense
		if diffs[diff] != 0 {
			return expense * diff
		}
		diffs[expense] = diff
	}
	return 0
}

func part2(expenses string) int {
	expenseList := input2int(expenses)

	diffs := make([]map[int]int, len(expenseList))
	for i := range expenseList {
		diffs[i] = make(map[int]int)
	}

	for i := range expenseList {
		for j := range expenseList {
			if i == j {
				continue
			}

			diff := 2020 - expenseList[i] - expenseList[j]
			if diffs[i][diff] != 0 {
				return expenseList[j] * expenseList[i] * diff
			}
			diffs[i][expenseList[j]] = diff

		}
	}
	return 0
}

func main() {
	txt := readInput("input.txt")
	p1 := part1(txt)
	fmt.Println("[PART 1] Result:", p1)

	p2 := part2(txt)
	fmt.Println("[PART 2] Result:", p2)
}
