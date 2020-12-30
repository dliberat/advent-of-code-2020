package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

type boardingPass struct {
	code int
}

func (bp *boardingPass) row() int {
	return bp.code >> 3
}

func (bp *boardingPass) col() int {
	return bp.code & 7
}

func (bp *boardingPass) sid() int {
	return bp.row()*8 + bp.col()
}

func (bp *boardingPass) toString() string {
	code := bp.code
	s := make([]byte, 10)
	for i := 9; i >= 7; i-- {
		if code%2 == 1 {
			s[i] = 'R'
		} else {
			s[i] = 'L'
		}
		code = code >> 1
	}
	for i := 6; i >= 0; i-- {
		if code%2 == 1 {
			s[i] = 'B'
		} else {
			s[i] = 'F'
		}
		code = code >> 1
	}
	return string(s)
}

func makeBoardingPass(txt string) boardingPass {
	code := 0
	for i := 0; i < len(txt); i++ {
		code = code << 1
		if txt[i] == 'B' || txt[i] == 'R' {
			code++
		}
	}
	return boardingPass{code: code}
}

func readInput(fname string) string {
	data, err := ioutil.ReadFile(fname)
	if err != nil {
		panic("Cannot read input data.")
	}
	s := string(data)
	return strings.TrimRight(s, "\n")
}

func part1(txt string) int {
	passes := strings.Split(txt, "\n")
	max := 0
	for _, pass := range passes {
		p := makeBoardingPass(pass)
		if p.sid() > max {
			max = p.sid()
		}
	}
	return max
}

func part2(txt string) int {
	split := strings.Split(txt, "\n")
	sids := make([]int, len(split))
	for i, p := range split {
		pass := makeBoardingPass(p)
		sids[i] = pass.sid()
	}
	sort.Ints(sids)

	// find the gap in a sequence of ints
	for i := 1; i < len(sids)-1; i++ {
		if sids[i]-sids[i-1] == 2 {
			return sids[i] - 1
		}
	}
	return 0
}

func main() {
	txt := readInput("input.txt")
	p1 := part1(txt)
	fmt.Println("[PART 1] Highest ID on a boarding pass:", p1)
	p2 := part2(txt)
	fmt.Println("[PART 2] Your seat ID is:", p2)
}
