package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

var memo = make(map[int]int64, 0)

func readInput(fname string) string {
	data, err := ioutil.ReadFile(fname)
	if err != nil {
		panic("Cannot read input data.")
	}
	s := string(data)
	return strings.TrimRight(s, "\n")
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

func max(seq []int) int {
	m := seq[0]
	for i := 1; i < len(seq); i++ {
		if seq[i] > m {
			m = seq[i]
		}
	}
	return m
}

func remove(s []int, i int) []int {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func part1(nums []int) int {
	sort.Ints(nums)
	diff1 := 0
	diff3 := 0
	for i := 1; i < len(nums); i++ {
		diff := nums[i] - nums[i-1]
		if diff == 1 {
			diff1++
		}
		if diff == 3 {
			diff3++
		}
	}
	return diff1 * diff3
}

// rec requires a sorted list of ints as input
func rec(nums []int) int64 {

	val, ok := memo[nums[0]]
	if ok {
		return val
	}

	if len(nums) == 1 {
		return 1
	}

	possibleNextAdapterIndices := make([]int, 0)
	for i := 1; i < len(nums); i++ {
		diff := nums[i] - nums[0]
		if diff > 3 {
			break
		}
		possibleNextAdapterIndices = append(possibleNextAdapterIndices, i)
	}

	totalPerms := int64(0)
	for _, i := range possibleNextAdapterIndices {
		totalPerms += rec(nums[i:])
	}
	memo[nums[0]] = totalPerms
	return totalPerms
}

func part2(nums []int) int64 {
	sort.Ints(nums)
	memo = make(map[int]int64, 0)
	return rec(nums)
}

func main() {
	txt := readInput("input.txt")
	nums := input2slice(txt)
	nums = append(nums, 0)           // add in the airplane terminal
	nums = append(nums, max(nums)+3) // add in the handheld device
	p1 := part1(nums)
	fmt.Println("[PART 1] Result:", p1)
	p2 := part2(nums)
	fmt.Println("[PART 2] Result:", p2)
}
