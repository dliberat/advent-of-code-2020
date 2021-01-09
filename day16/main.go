package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type ticket []int

type validityTable [][]bool

func (vt *validityTable) invalidate(fieldIndex, ruleIndex int) {
	(*vt)[fieldIndex][ruleIndex] = false

	// did we already narrow down the candidate list to a single possibility?
	singleCandidateIndex := -1
	candidateCount := 0
	for i, isCandidate := range (*vt)[fieldIndex] {
		if isCandidate {
			singleCandidateIndex = i
			candidateCount++
			if candidateCount > 1 {
				break
			}
		}
	}

	// if so, no other fields can refer to that rule.
	if candidateCount == 1 {
		for i := range *vt {
			if i == fieldIndex {
				continue
			}
			(*vt)[i][singleCandidateIndex] = false
		}
	}
}

func (vt *validityTable) getRuleIndex(fieldIndex int) (int, error) {
	fieldRow := (*vt)[fieldIndex]
	count := 0
	validIndex := -1
	for i := range fieldRow {
		if fieldRow[i] {
			count++
			validIndex = i
		}
	}

	if count > 1 {
		return -1, errors.New("process of elimination has not ruled out all invalid rules")
	}
	if count < 1 {
		return -1, errors.New("no candidate rules remain. All candidates have been deemed invalid")
	}
	return validIndex, nil
}

func (vt *validityTable) verticalElimination() {
	// stop iterating once there is nothing more that
	// can be eliminated using this strategy
	changeCount := 1
	for changeCount > 0 {
		changeCount = 0

		// scan each column one at a time from top to bottom
		for col := 0; col < len((*vt)[0]); col++ {
			trueCount := 0
			trueIndex := -1
			for row := 0; row < len(*vt); row++ {
				if (*vt)[row][col] {
					trueCount++
					trueIndex = row
					if trueCount > 1 {
						break
					}
				}
			}

			// if only one row in the column is true, then we know for sure
			// that that row must correspond to the field at position col,
			// and therefore all other candidates for that row can be rejected
			if trueCount == 1 {
				for c := 0; c < len((*vt)[0]); c++ {
					if c == col {
						continue
					}
					if (*vt)[trueIndex][c] {
						changeCount++
						(*vt)[trueIndex][c] = false
					}
				}
			}
		}
	}
}

func (vt *validityTable) toString() string {
	var b strings.Builder
	for i := range *vt {
		for _, val := range (*vt)[i] {
			if val {
				fmt.Fprintf(&b, "T ")
			} else {
				fmt.Fprintf(&b, "- ")
			}
		}
		fmt.Fprintf(&b, "\n")
	}
	return b.String()
}

type rng struct {
	from int
	to   int
}

func (r *rng) isInRange(n int) bool {
	return n >= r.from && n <= r.to
}

type rule struct {
	name   string
	ranges []rng
}

func (r *rule) fieldIsValid(n int) bool {
	for _, rr := range r.ranges {
		if rr.isInRange(n) {
			return true
		}
	}
	return false
}

func readInput(fname string) string {
	data, err := ioutil.ReadFile(fname)
	if err != nil {
		panic("Cannot read input data.")
	}
	s := string(data)
	return strings.TrimRight(s, "\n")
}

func makeRange(txt string) rng {
	// example txt: "44-56"
	s := strings.Split(txt, "-")
	f, t := s[0], s[1]
	from, err := strconv.Atoi(f)
	if err != nil {
		panic("bad input: " + txt)
	}
	to, err := strconv.Atoi(t)
	if err != nil {
		panic("bad input: " + txt)
	}
	return rng{from: from, to: to}
}

func makeTicket(line string) ticket {
	s := strings.Split(line, ",")
	tik := make(ticket, len(s))
	for i, strVal := range s {
		numVal, err := strconv.Atoi(strVal)
		if err != nil {
			panic("bad 'your ticket' input:" + line)
		}
		tik[i] = numVal
	}
	return tik
}

func parseRuleset(ruleset string) []rule {
	ruleList := make([]rule, 0)
	for _, line := range strings.Split(ruleset, "\n") {
		if len(line) == 0 {
			continue
		}
		s := strings.Split(line, ": ")
		r := rule{name: s[0], ranges: make([]rng, 0)}
		s = strings.Split(s[1], " or ")
		for _, substr := range s {
			r.ranges = append(r.ranges, makeRange(substr))
		}
		ruleList = append(ruleList, r)
	}
	return ruleList
}

func parseTickets(txt string) []ticket {
	tix := make([]ticket, 0)
	for _, line := range strings.Split(txt, "\n") {
		if len(line) == 0 || line == "nearby tickets:" || line == "your ticket:" {
			continue
		}
		tix = append(tix, makeTicket(line))
	}
	return tix
}

// parseInput into a slice of rules, a single ticket (your ticket),
// and a slice of tickets (nearby tickets)
func parseInput(txt string) ([]rule, ticket, []ticket) {
	split := strings.Split(txt, "\n\n")
	ruleset, yourTicket, nearbyTickets := split[0], split[1], split[2]

	ruleList := parseRuleset(ruleset)
	tik := parseTickets(yourTicket)[0]
	tix := parseTickets(nearbyTickets)
	return ruleList, tik, tix
}

func ticketScanErrorRate(t ticket, rules *[]rule) (int, bool) {
	errRate := 0
	isValidTicket := true
	for _, field := range t {
		isValid := false
		for _, r := range *rules {
			if r.fieldIsValid(field) {
				isValid = true
				break
			}
		}
		if !isValid {
			errRate += field
			isValidTicket = false
		}
	}
	return errRate, isValidTicket
}

func filterInvalidTickets(tickets []ticket, rules *[]rule) []ticket {
	validTix := make([]ticket, 0)
	for _, t := range tickets {
		_, isValidTicket := ticketScanErrorRate(t, rules)
		if isValidTicket {
			validTix = append(validTix, t)
		}
	}
	return validTix
}

func makeValidityTable(ticketLen, ruleLen int) validityTable {
	vt := make([][]bool, ticketLen)
	for i := range vt {
		vt[i] = make([]bool, ruleLen)
		for j := range vt[i] {
			vt[i][j] = true
		}
	}
	return vt
}

func invalidTicketExists(tix *[]ticket, i int, r rule) bool {
	for _, t := range *tix {
		if !r.fieldIsValid(t[i]) {
			return true
		}
	}
	return false
}

func processOfElimination(vt validityTable, tix *[]ticket, rules []rule) validityTable {
	for i := 0; i < len(vt); i++ {
		for j := 0; j < len(vt[i]); j++ {
			if !vt[i][j] {
				// if the entry at (i,j) has already been eliminated
				// there is no need to test it again
				continue
			}
			if invalidTicketExists(tix, i, rules[j]) {
				vt.invalidate(i, j)
			}
		}
	}
	return vt
}

func mapFieldsToRules(vt validityTable, rules []rule) map[int]rule {
	mapping := make(map[int]rule, 0)
	sanityCheck := make(map[int]bool, 0)
	for i := range vt {
		ruleIndex, err := vt.getRuleIndex(i)
		if err != nil {
			fmt.Println(vt.toString())
			panic(err)
		}
		mapping[i] = rules[ruleIndex]

		if sanityCheck[ruleIndex] {
			fmt.Println(vt)
			panic("multiple fields are been associated with the same rule")
		}
		sanityCheck[ruleIndex] = true
	}
	return mapping
}

func part1(txt string) int {
	rules, _, nearbyTix := parseInput(txt)
	errRate := 0
	for _, t := range nearbyTix {
		e, _ := ticketScanErrorRate(t, &rules)
		errRate += e
	}
	return errRate
}

func part2(txt string) int {
	/* Create a matrix whose rows represent field indices on tickets,
		 and columns represent candidate rules in the ruleset.
		 Initially, each field could potentially be linked to any rule,
		 so all fields are initialized to True.

	                rule
		ticket     index
		field     0  1  2
		  0       T  T  T
		  1       T  T  T
		  2       T  T  T
		  3       T  T  T

		Suppose we scan a ticket and notice that field 0 is incompatible
		with rule index 1. At that point, we set the corresponding entry to F.

		            rule
		ticket     index
		field     0  1  2
		0         T  F  T
		1         T  T  T
		2         T  T  T
		3         T  T  T

		In principle, we can continue this process of elimination until only
		a single true entry remains in each row. However, it's also possible
		to have a situation like this:

		            rule
		ticket     index
		field     0  1  2  3
		0         T  F  T  F
		1         F  T  T  F
		2         T  T  T  T
		3         F  T  T  F

		Here, there are no rows where we've completed the initial process of
		elimination. However, we can see that the only available candidate for
		rule 3 is field 2. We can therefore convert all other entries in row 2
		to False and continue the process of elimination.
	*/
	rules, myticket, nearbyTickets := parseInput(txt)
	allTickets := append(nearbyTickets, myticket)
	allTickets = filterInvalidTickets(allTickets, &rules)
	vt := makeValidityTable(len(myticket), len(rules))
	vt = processOfElimination(vt, &allTickets, rules)
	vt.verticalElimination()

	mapping := mapFieldsToRules(vt, rules)

	total := 1
	sanityCheck := 0
	for i, r := range mapping {
		if strings.Contains(r.name, "departure") {
			total *= myticket[i]
			sanityCheck++
		}
	}
	if sanityCheck != 6 {
		panic("expected only 6 fields with the word 'departure'")
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
