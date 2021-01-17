package main

import "testing"

func TestParseInput(t *testing.T) {
	nodes := parseInput("186524973")
	expected := []int{1, 8, 6, 5, 2, 4, 9, 7, 3}
	head := &nodes[0]
	if head.val != expected[0] {
		t.Errorf("%d != %d", expected[0], head.val)
		return
	}
	expected = expected[1:]

	current := head.next
	for current != head {
		if current.val != expected[0] {
			t.Errorf("%d != %d", expected[0], current.val)
			return
		}
		current = current.next
		expected = expected[1:]
	}
}

func TestRemoveN(t *testing.T) {
	nodes := parseInput("186524973")
	// remove nodes 8,6,5
	one := &nodes[0]
	eight := removeN(one, 3)

	two := one.next
	four := two.next
	nine := four.next
	seven := nine.next
	three := seven.next
	looparound := three.next
	if one.val != 1 || two.val != 2 || four.val != 4 || nine.val != 9 || seven.val != 7 || three.val != 3 || looparound != one {
		t.Errorf("Incorrect node removal: [%d, %d, %d, %d, %d, %d]", one.val, two.val, four.val, nine.val, seven.val, three.val)
		return
	}

	six := eight.next
	five := six.next
	looparound = five.next
	if eight.val != 8 || six.val != 6 || five.val != 5 || looparound != eight {
		t.Errorf("Incorrect nodes removed: [%d, %d, %d]", eight.val, six.val, five.val)
	}

}

func TestFindTail(t *testing.T) {
	nodes := parseInput("186524973")
	head := &nodes[0]
	tail := findTail(head)
	if tail.val != 3 {
		t.Errorf("Tail should be 3, got %d", tail.val)
	}
}

func TestInsert01(t *testing.T) {
	nodes := parseInput("13")
	toInsert := parseInput("2")
	destHead := &nodes[0]
	insertHead := &toInsert[0]

	res := insert(destHead, insertHead, 1)
	if !res {
		t.Error("Node should have been successfully inserted.")
		return
	}
	if destHead.val != 1 {
		t.Errorf("List head should be 1")
		return
	}
	two := destHead.next
	if two.val != 2 {
		t.Errorf("Second node in list should be 2")
		return
	}
	three := two.next
	if three.val != 3 {
		t.Errorf("Third node in list should be 3")
		return
	}
	looparound := three.next
	if looparound != destHead {
		t.Errorf("List should loop back around to 1")
	}
}

func assertListEqual(head *node, list []int) bool {
	for len(list) > 0 {
		if head.val != list[0] {
			return false
		}
		list = list[1:]
		head = head.next
	}
	return true
}

func TestMove01(t *testing.T) {
	nodes := parseInput("389125467")
	three := &nodes[0]
	expected := []int{3, 2, 8, 9, 1, 5, 4, 6, 7}
	newCurrentCup := move(three, 9)
	if !assertListEqual(three, expected) {
		t.Error("node order does not match expected")
		return
	}
	if newCurrentCup.val != 2 {
		t.Errorf("New current cup should be 2, but got %d", newCurrentCup.val)
	}
}

func TestMove02(t *testing.T) {
	nodes := parseInput("289154673")
	two := &nodes[0]
	expected := []int{2, 5, 4, 6, 7, 8, 9, 1, 3}
	newCurrentCup := move(two, 9)
	if !assertListEqual(two, expected) {
		t.Error("node order does not match expected")
		return
	}
	if newCurrentCup.val != 5 {
		t.Errorf("New current cup should be 5, but got %d", newCurrentCup.val)
	}
}

func TestPart1(t *testing.T) {
	txt := "389125467"
	res := part1(txt)
	if res != "67384529" {
		t.Errorf("67384529 != res")
	}
}
