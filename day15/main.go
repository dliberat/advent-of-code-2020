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

// memoryGame takes the starting values as a slice of ints and a target step number,
// and produces the number that would be spoken at the target step
func memoryGame(nums []int, target int) int {
	ages := make(map[int]int, 0)
	previousNum := nums[0] // step 1

	step := 2
	for i := 1; i < len(nums); i++ {
		ages[previousNum] = step - 1
		previousNum = nums[i]
		step++
	}

	for ; step <= target; step++ {
		v, ok := ages[previousNum]
		if !ok {
			// number has never been said before
			ages[previousNum] = step - 1
			previousNum = v
		} else {
			ages[previousNum] = step - 1
			previousNum = (step - 1) - v
		}
	}

	return previousNum
}

func part1(nums []int) int {
	return memoryGame(nums, 2020)
}

func part2(nums []int) int {
	return memoryGame(nums, 30_000_000)
}

func main() {
	txt := readInput("input.txt")
	s := strings.Split(txt, ",")
	nums := make([]int, len(s))
	for i := range s {
		val, err := strconv.Atoi(s[i])
		if err != nil {
			panic("bad input")
		}
		nums[i] = val
	}
	p1 := part1(nums)
	fmt.Println("[PART 1] Result:", p1)
	p2 := part2(nums)
	fmt.Println("[PART 2] Result:", p2)
}
