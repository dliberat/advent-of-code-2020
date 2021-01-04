package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

type mask struct {
	ones  int64 // bitmask with 1s set where 1s appear in the original string
	zeros int64 // bitmask with 1s set where 0s appear in the original string
	xes   []int // indices (zero-based) of the digits where Xs appear in the original string
}

// applyPart1 applies the mask as specified in the directions for part 1 of the problem.
// Wherever the bitmask specifies a 1, overwrite n with a 1
// Wherever the bitmask specifies a 0, overwrite n with a 0
// Wherever the bitmask specifies an X, leave n unchanged
func (m *mask) applyPart1(n int64) int64 {
	n = m.overwrite1s(n)
	return m.overwrite0s(n)
}

func (m *mask) overwrite1s(n int64) int64 {
	return n | m.ones
}

func (m *mask) overwrite0s(n int64) int64 {
	return n &^ m.zeros
}

type memop struct {
	m       mask
	address int64
	value   int64
}

// getTargetAddresses returns a slice of the memory addresses that
// should be written to by the memory operation m.
// This method is only applicable to part 2 of the puzzle.
func (m *memop) getTargetAddresses() []int64 {
	addresses := make([]int64, 0)
	n := m.m.overwrite1s(m.address)
	perm(n, m.m.xes, &addresses)
	return addresses
}

// perm populates the perms slice with permutations of n with each bit at
// the specified floatIndices set to either 0 or 1.
// Example: n = 11 (01011), floatIndices = [0, 4]
// Result:
// Indices 0 and 4 are set to 1s and 0s
// indices:    4 3 2 1 0
//             |       |
// outputs: -  0 1 0 1 1
//          -  1 1 0 1 1
//          -  0 1 0 1 0
//          -  1 1 0 1 0
func perm(n int64, floatIndices []int, perms *[]int64) {
	if len(floatIndices) == 0 {
		*perms = append(*perms, n)
		return
	}

	currentIndex := floatIndices[0]
	floatIndices = floatIndices[1:]

	// zero branch
	newTarget := n &^ (1 << currentIndex)
	perm(newTarget, floatIndices, perms)

	// one branch
	newTarget = n | (1 << (currentIndex))
	perm(newTarget, floatIndices, perms)
}

func makeMask(code string) mask {
	m := mask{}
	m.xes = make([]int, 0)
	n := 0
	for i := len(code) - 1; i >= 0; i-- {
		if code[i] == '1' {
			m.ones += int64(math.Pow(2, float64(n)))
		} else if code[i] == '0' {
			m.zeros += int64(math.Pow(2, float64(n)))
		} else {
			m.xes = append(m.xes, n)
		}
		n++
	}
	return m
}

func readInput(fname string) string {
	data, err := ioutil.ReadFile(fname)
	if err != nil {
		panic("Cannot read input data.")
	}
	s := string(data)
	return strings.TrimRight(s, "\n")
}

func parseMemOp(lhs, rhs string) (int64, int64) {
	lhs = strings.Replace(lhs, "mem[", "", 1)
	lhs = strings.Replace(lhs, "]", "", 1)
	address, err := strconv.Atoi(lhs)
	if err != nil {
		panic("bad input:" + lhs + " = " + rhs)
	}

	value, err := strconv.Atoi(rhs)
	if err != nil {
		panic("bad input:" + lhs + " = " + rhs)
	}
	return int64(address), int64(value)
}

func parseProgramPart2(txt string) []memop {
	lines := strings.Split(txt, "\n")
	memops := make([]memop, 0)
	activeMask := makeMask("X")
	for _, line := range lines {
		split := strings.Split(line, " = ")
		lhs, rhs := split[0], split[1]
		if lhs == "mask" {
			activeMask = makeMask(rhs)
		} else {
			addr, arg := parseMemOp(lhs, rhs)
			memops = append(memops, memop{m: activeMask, address: addr, value: arg})
		}
	}
	return memops
}

func part1(txt string) int64 {
	mem := make(map[int64]int64, 0)
	lines := strings.Split(txt, "\n")
	activeMask := makeMask("X")

	for _, line := range lines {
		split := strings.Split(line, " = ")
		lhs, rhs := split[0], split[1]
		if lhs == "mask" {
			activeMask = makeMask(rhs)
		} else {
			addr, arg := parseMemOp(lhs, rhs)
			mem[addr] = activeMask.applyPart1(arg)
		}
	}

	total := int64(0)
	for _, val := range mem {
		total += val
	}
	return total
}

func part2(txt string) int64 {
	operations := parseProgramPart2(txt)
	targetAddresses := make(map[int64]int64, 0)

	for i := len(operations) - 1; i >= 0; i-- {
		addrs := operations[i].getTargetAddresses()
		for _, addr := range addrs {
			_, ok := targetAddresses[addr]
			if !ok {
				targetAddresses[addr] = operations[i].value
			}
		}
	}

	total := int64(0)
	for _, val := range targetAddresses {
		total += val
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
