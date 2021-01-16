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
	// if deck is empty before attempting to pop
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

	// if deck is empty after popping
	if d.count == 0 {
		d.tail = nil
	}

	return head
}

// score calculates the player's total score.
// The value of the card at position n is multiplied by n+1.
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

func (d *deck) hash() string {
	nums := d.toIntSlice()
	return fmt.Sprintf("%v", nums)
}

func (d *deck) cloneToDepth(n int) deck {
	if d.count < n {
		panic("Cloning beyond deck depth is not permitted")
	}
	newDeck := deck{}
	newNodes := make([]card, n)

	node := d.head
	for i := 0; i < n; i++ {
		newNodes[i] = *node
		node = node.next
	}

	for i := range newNodes {
		newDeck.addToBottom(&newNodes[i])
	}

	return newDeck
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

func recursiveCombat(p1, p2 *deck) int {
	memo := make(map[string]bool, 0)

	for p1.count > 0 && p2.count > 0 {
		// before either player deals a card, if there was a previous round in this game
		// that had exactly the same cards in the same order in the same players' decks,
		// the game instantly ends in a win for player 1.
		hash := p1.hash() + p2.hash()
		if memo[hash] {
			return 1
		}
		memo[hash] = true

		// players begin the round by each drawing the top card of their deck
		p1card := p1.popHead()
		p2card := p2.popHead()

		if p1.count >= p1card.val && p2.count >= p2card.val {

			// to recurse, each player creates a new deck by making a copy of the next
			// cards in their deck (the quantity of cards copied is equal to the number
			// on the card they drew to trigger the recursion)
			p1clone := p1.cloneToDepth(p1card.val)
			p2clone := p2.cloneToDepth(p2card.val)
			if recursiveCombat(&p1clone, &p2clone) == 1 {
				p1.addToBottom(p1card)
				p1.addToBottom(p2card)
			} else {
				p2.addToBottom(p2card)
				p2.addToBottom(p1card)
			}

		} else {
			// at least one player does not have enough cards to recurse.
			// Winner is the player with the higher-value card
			if p1card.val > p2card.val {
				p1.addToBottom(p1card)
				p1.addToBottom(p2card)
			} else {
				p2.addToBottom(p2card)
				p2.addToBottom(p1card)
			}
		}
	}

	if p1.count > p2.count {
		return 1
	}
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
	p1, p2 := parseInput(txt)
	winner := recursiveCombat(&p1, &p2)
	if winner == 1 {
		return p1.score()
	}
	return p2.score()
}

func main() {
	txt := readInput("input.txt")
	p1 := part1(txt)
	fmt.Println("[PART 1] Result:", p1)
	p2 := part2(txt)
	fmt.Println("[PART 2] Result:", p2)
}
