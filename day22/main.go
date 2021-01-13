package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type card struct {
	val  int
	next *card
	prev *card
}

// deck contains an ordered list of cards. Because we play cards from the top downard,
// the .next pointers point toward the bottom of the deck.
// head                       tail
//    <--prev        next -->
type deck struct {
	head  *card
	tail  *card
	count int
}

func (d *deck) addToBottom(c *card) {
	c.next = nil
	c.prev = d.tail
	if d.tail != nil {
		d.tail.next = c
	} else {
		d.head = c
	}
	d.tail = c
	d.count++
}

func (d *deck) popHead() *card {
	// empty deck
	if d.head == nil {
		return d.head
	}

	head := d.head

	// second card in the deck moves up to top spot
	if head.next != nil {
		head.next.prev = nil
	}
	d.head = head.next

	// popped node does not belong in a deck anymore
	head.next = nil
	head.prev = nil

	// update count
	d.count--

	// empty deck
	if d.count == 0 {
		d.tail = nil
	}

	return head
}

func (d *deck) score() int {
	if d.tail == nil {
		return 0
	}
	total := 0
	i := 1
	node := d.tail
	for node != nil {
		total += node.val * i
		i++
		node = node.prev
	}
	return total
}

func (d *deck) toIntSlice() []int {
	nums := make([]int, 0)
	node := d.head
	for node != nil {
		nums = append(nums, node.val)
		node = node.next
	}
	return nums
}

func readInput(fname string) string {
	data, err := ioutil.ReadFile(fname)
	if err != nil {
		panic("Cannot read input data.")
	}
	s := string(data)
	return strings.TrimRight(s, "\n")
}

// parseInput into Player 1's deck and Player 2's deck
func parseInput(txt string) (deck, deck) {
	p1 := deck{}
	p2 := deck{}

	lines := strings.Split(txt, "\n")
	i := 0
	for ; lines[i] != ""; i++ {
		num, err := strconv.Atoi(lines[i])
		if err != nil {
			// empty line or "Player 1:"
			continue
		}
		c := card{val: num}
		p1.addToBottom(&c)
	}
	for ; i < len(lines); i++ {
		num, err := strconv.Atoi(lines[i])
		if err != nil {
			// empty line or "Player 2:"
			continue
		}
		c := card{val: num}
		p2.addToBottom(&c)
	}

	return p1, p2
}

func playRound(p1, p2 *deck) int {
	p1card := p1.popHead()
	p2card := p2.popHead()

	if p1card.val > p2card.val {
		p1.addToBottom(p1card)
		p1.addToBottom(p2card)
		return 1
	}
	p2.addToBottom(p2card)
	p2.addToBottom(p1card)
	return 2
}

func part1(txt string) int {
	p1, p2 := parseInput(txt)

	var winner int
	for p1.count > 0 && p2.count > 0 {
		winner = playRound(&p1, &p2)
	}
	if winner == 1 {
		return p1.score()
	}
	return p2.score()
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
