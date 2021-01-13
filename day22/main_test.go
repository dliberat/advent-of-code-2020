package main

import "testing"

func TestAddToBottomOrdering(t *testing.T) {
	c1 := card{val: 1}
	c2 := card{val: 2}
	c3 := card{val: 3}

	d := deck{}
	d.addToBottom(&c1)
	if c1.next != nil || c1.prev != nil {
		t.Errorf("c1 should be alone in the deck")
		return
	}

	d.addToBottom(&c2)
	if c1.prev != nil || c1.next != &c2 || c2.prev != &c1 {
		t.Errorf("Deck should contain [1, 2]")
		return
	}

	d.addToBottom(&c3)
	if c1.prev != nil || c1.next != &c2 || c2.prev != &c1 || c2.next != &c3 || c3.prev != &c2 || c3.next != nil {
		t.Errorf("Deck should contains [1, 2, 3]")
		return
	}
}

func TestAddToBottomHeadTail(t *testing.T) {
	c1 := card{val: 1}
	c2 := card{val: 2}
	c3 := card{val: 3}
	d := deck{}
	d.addToBottom(&c1)
	if d.head != &c1 {
		t.Error("List head should be c1")
		return
	}
	if d.tail != &c1 {
		t.Error("List tail should be c1")
		return
	}
	d.addToBottom(&c2)
	if d.head != &c1 {
		t.Error("List head should be c1")
		return
	}
	if d.tail != &c2 {
		t.Error("List tail should be c2")
		return
	}
	d.addToBottom(&c3)
	if d.head != &c1 {
		t.Error("List head should be c1")
		return
	}
	if d.tail != &c3 {
		t.Error("List tail should be c3")
		return
	}
}

func TestAddToBottomCount(t *testing.T) {
	c1 := card{val: 1}
	c2 := card{val: 2}
	c3 := card{val: 3}
	d := deck{}
	if d.count != 0 {
		t.Errorf("0 != %d", d.count)
	}
	d.addToBottom(&c1)
	if d.count != 1 {
		t.Errorf("1 != %d", d.count)
	}
	d.addToBottom(&c2)
	if d.count != 2 {
		t.Errorf("2 != %d", d.count)
	}
	d.addToBottom(&c3)
	if d.count != 3 {
		t.Errorf("3 != %d", d.count)
	}
}

func TestPopHeadEmptyList(t *testing.T) {
	d := deck{}
	if d.popHead() != nil {
		t.Error("Should return nil if the deck is empty.")
	}
}

func TestPopHead(t *testing.T) {
	c1 := card{val: 1}
	c2 := card{val: 2}
	c3 := card{val: 3}
	d := deck{}
	d.addToBottom(&c1)
	d.addToBottom(&c2)
	d.addToBottom(&c3)

	head := d.popHead()
	if head != &c1 {
		t.Error("Should have popped c1")
		return
	}
	if d.head != &c2 {
		t.Error("c2 should be on top of the deck after removing c1")
	}
	if d.tail != &c3 {
		t.Errorf("c3 should be on bottom")
	}
	if head.prev != nil || head.next != nil {
		t.Error("c1 should no longer have pointers to any other cards")
	}
}

func TestPopHeadCount(t *testing.T) {
	c1 := card{val: 1}
	c2 := card{val: 2}
	c3 := card{val: 3}
	d := deck{}
	d.addToBottom(&c1)
	d.addToBottom(&c2)
	d.addToBottom(&c3)
	if d.count != 3 {
		t.Errorf("3 != %d", d.count)
		return
	}

	d.popHead()
	if d.count != 2 {
		t.Errorf("2 != %d", d.count)
		return
	}
	d.popHead()
	if d.count != 1 {
		t.Errorf("1 != %d", d.count)
		return
	}
	d.popHead()
	if d.count != 0 {
		t.Errorf("0 != %d", d.count)
		return
	}
	if d.head != nil {
		t.Error("Deck should not have a head")
		return
	}
	if d.tail != nil {
		t.Error("Deck should not have a tail")
	}
}

func intSliceEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestParseInput(t *testing.T) {
	txt := "Player 1:\n9\n2\n6\n3\n1\n\nPlayer 2:\n5\n8\n4\n7\n10"
	p1, p2 := parseInput(txt)
	expectedP1 := []int{9, 2, 6, 3, 1}
	expectedP2 := []int{5, 8, 4, 7, 10}
	p1cards := p1.toIntSlice()
	p2cards := p2.toIntSlice()
	if !intSliceEqual(expectedP1, p1cards) {
		t.Errorf("%v != %v", expectedP1, p1cards)
		return
	}
	if !intSliceEqual(expectedP2, p2cards) {
		t.Errorf("%v != %v", expectedP2, p2cards)
	}
}

func TestPlayRound01(t *testing.T) {
	txt := "Player 1:\n9\n2\n6\n3\n1\n\nPlayer 2:\n5\n8\n4\n7\n10"
	p1, p2 := parseInput(txt)
	expectedP1 := []int{2, 6, 3, 1, 9, 5}
	expectedP2 := []int{8, 4, 7, 10}

	roundWinner := playRound(&p1, &p2)
	if roundWinner != 1 {
		t.Errorf("Player 1 should have won the round")
	}

	p1cards := p1.toIntSlice()
	p2cards := p2.toIntSlice()
	if !intSliceEqual(expectedP1, p1cards) {
		t.Errorf("%v != %v", expectedP1, p1cards)
		return
	}
	if !intSliceEqual(expectedP2, p2cards) {
		t.Errorf("%v != %v", expectedP2, p2cards)
	}
}

func TestScore(t *testing.T) {
	nums := []int{3, 2, 10, 6, 8, 5, 9, 4, 7, 1}
	cards := make([]card, len(nums))
	for i, num := range nums {
		cards[i] = card{val: num}
	}
	d := deck{}
	for i := range cards {
		d.addToBottom(&(cards[i]))
	}
	score := d.score()
	if score != 306 {
		t.Errorf("306 != %d", score)
	}
}
