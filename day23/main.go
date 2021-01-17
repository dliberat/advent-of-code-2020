package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type node struct {
	val  int
	next *node
}

func readInput(fname string) string {
	data, err := ioutil.ReadFile(fname)
	if err != nil {
		panic("Cannot read input data.")
	}
	s := string(data)
	return strings.TrimRight(s, "\n")
}

func parseInput(txt string) []node {
	nums := make([]node, len(txt))
	for i := 0; i < len(txt); i++ {
		val := int(txt[i]) - 48
		nums[i] = node{val: val}
	}
	for i := 0; i < len(nums); i++ {
		next := (i + 1) % len(nums)
		nums[i].next = &nums[next]
	}
	return nums
}

func findTail(head *node) *node {
	prev := head
	cur := head.next
	for cur.val != head.val {
		prev = cur
		cur = cur.next
	}
	return prev
}

// removeN nodes from the linked list starting with the node
// immediately following head. Returns a pointer
// to the first removed node, which is the head of a new list
func removeN(head *node, n int) *node {
	ret := head.next

	cur := head
	for i := 0; i < n; i++ {
		cur = cur.next
	}
	// cur is now the last node we need to extract
	head.next = cur.next
	cur.next = ret

	return ret
}

// insert node x immediately after the node with the value dest
// Returns true if the node was successfully inserted.
// Returns false if the value dest could not be found
func insert(head, x *node, dest int) bool {
	headVal := head.val
	cur := head
	for cur.val != dest {
		cur = cur.next
		if cur.val == headVal {
			return false
		}
	}
	xtail := findTail(x)
	xtail.next = cur.next
	cur.next = x
	return true
}

func move(head *node, nCups int) *node {
	// Pick up 3 cups that are immediately clockwise of current cup
	pickup := removeN(head, 3)

	// Select destination cup (label equal to current cup minus 1)
	destinationCupVal := head.val - 1
	if destinationCupVal < 1 {
		destinationCupVal += nCups
	}

	// If this would select one of the cups that was just picked up,
	// keep subtracting 1 until it finds a cup that wasn't picked up.
	// Place the picked up cups immediately clockwise of the current cup
	for !insert(head, pickup, destinationCupVal) {
		destinationCupVal--
		if destinationCupVal < 1 {
			destinationCupVal += nCups
		}
	}

	// New current cup is the cup immediately clockwise of the current cup
	return head.next
}

func collectCups(head *node) []int {
	for head.val != 1 {
		head = head.next
	}
	head = head.next
	nums := []int{1}
	for head.val != 1 {
		nums = append(nums, head.val)
		head = head.next
	}
	return nums
}

func part1(txt string) string {
	numCups := len(txt)
	cups := parseInput(txt)
	head := &cups[0]

	for i := 0; i < 100; i++ {
		head = move(head, numCups)
	}

	collected := collectCups(head)
	collected = collected[1:] // cup #1 is not needed for the output
	s := ""
	for _, num := range collected {
		s = fmt.Sprintf("%s%d", s, num)
	}
	return s
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
