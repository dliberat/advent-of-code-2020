package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type operand interface {
	operate(a, b int) int
}

type opAdd string

func (o opAdd) operate(a, b int) int {
	return a + b
}

type opMul string

func (o opMul) operate(a, b int) int {
	return a * b
}

type expression struct {
	raw   string
	value int
	exp   []expression
	ops   []operand
}

func (e *expression) reduce() int {
	if len(e.exp) == 0 {
		return e.value
	}
	values := make([]int, len(e.exp))
	for i, subexpr := range e.exp {
		values[i] = subexpr.reduce()
	}

	if len(values) != len(e.ops)+1 {
		panic(fmt.Sprintf("Value count and ops count don't match up. %d - 1 != %d", len(values), len(e.ops)))
	}

	acc := values[0]
	for i := 1; i < len(values); i++ {
		acc = e.ops[i-1].operate(acc, values[i])
	}

	return acc
}

func (e *expression) reducePart2() int {
	if len(e.exp) == 0 {
		return e.value
	}
	values := make([]int, len(e.exp))
	for i, subexpr := range e.exp {
		values[i] = subexpr.reducePart2()
	}

	if len(values) != len(e.ops)+1 {
		panic(fmt.Sprintf("Value count and ops count don't match up. %d - 1 != %d", len(values), len(e.ops)))
	}

	// this is a hack, but it computes the additions and
	// ignores multiplications and leaves the results in
	// the values slice
	var add opAdd = "+"
	valIndex := 0
	for i := 0; i < len(e.ops); i++ {
		op := e.ops[i]
		if op == add {
			values[valIndex] = values[valIndex] + values[valIndex+1]
			values = remove(values, valIndex+1)
		} else {
			valIndex++
		}
	}

	// And now we handle all the multiplications
	acc := 1
	for _, val := range values {
		acc *= val
	}

	return acc
}

func readInput(fname string) string {
	data, err := ioutil.ReadFile(fname)
	if err != nil {
		panic("Cannot read input data.")
	}
	s := string(data)
	return strings.TrimRight(s, "\n")
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func parse(txt string) expression {
	ex := expression{raw: txt}
	ex.exp = make([]expression, 0)
	ex.ops = make([]operand, 0)

	for i := 0; i < len(txt); i++ {
		char := txt[i : i+1]
		// whitespace
		if char == " " {
			continue
		}
		// operators
		if char == "+" {
			var o opAdd = "+"
			ex.ops = append(ex.ops, o)
		} else if char == "*" {
			var o opMul = "*"
			ex.ops = append(ex.ops, o)
		} else if char == "(" {
			// sub expressions
			bracketStack := 1
			for j := i + 1; j <= len(txt); j++ {
				nxtChar := txt[j : j+1]
				if nxtChar == "(" {
					bracketStack++
				} else if nxtChar == ")" {
					bracketStack--
				}

				if bracketStack == 0 {
					x := parse(txt[i+1 : j])
					ex.exp = append(ex.exp, x)
					i = j + 1
					break
				}
			}
			if bracketStack != 0 {
				panic("Unparseable expression: " + txt)
			}
		} else {
			// values
			val, err := strconv.Atoi(string(char))
			if err != nil {
				panic("Else is not a number: " + string(char))
			}
			x := expression{
				raw:   string(char),
				value: val,
			}
			ex.exp = append(ex.exp, x)
		}
	}

	return ex
}

func part1(txt string) int {
	lines := strings.Split(txt, "\n")

	total := 0
	for _, line := range lines {
		ex := parse(line)
		total += ex.reduce()
	}

	return total
}

func part2(txt string) int {
	lines := strings.Split(txt, "\n")

	total := 0
	for _, line := range lines {
		ex := parse(line)
		total += ex.reducePart2()
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
