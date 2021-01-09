package main

import "testing"

func TestParseRuleSet(t *testing.T) {
	txt := "class: 1-3 or 5-7\nrow: 6-11 or 33-44\nseat: 13-40 or 45-50"

	ruleset := parseRuleset(txt)
	if len(ruleset) != 3 {
		t.Errorf("Expected 3 rules but got %d", len(ruleset))
		return
	}

	names := []string{"class", "row", "seat"}
	for i := range ruleset {
		if ruleset[i].name != names[i] {
			t.Errorf("Expected name = %s but got %s", names[i], ruleset[i].name)
			return
		}
	}

	ranges := []rng{
		rng{from: 1, to: 3},
		rng{from: 5, to: 7},
		rng{from: 6, to: 11},
		rng{from: 33, to: 44},
		rng{from: 13, to: 40},
		rng{from: 45, to: 50},
	}
	j := 0
	for _, entry := range ruleset {
		for _, rangedata := range entry.ranges {
			if rangedata.from != ranges[j].from || rangedata.to != ranges[j].to {
				t.Errorf("%v != %v", ranges[j], rangedata)
				return
			}
			j++
		}
	}

}

func intSliceEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestParseTickets01(t *testing.T) {
	txt := "nearby tickets:\n7,3,47\n40,4,50\n55,2,20\n38,6,12"

	tix := parseTickets(txt)
	if len(tix) != 4 {
		t.Errorf("Expected 4 tickets, but got %d", len(tix))
	}

	expected := []ticket{
		[]int{7, 3, 47},
		[]int{40, 4, 50},
		[]int{55, 2, 20},
		[]int{38, 6, 12},
	}
	for i := range tix {
		if !intSliceEqual(tix[i], expected[i]) {
			t.Errorf("%v != %v", tix[i], expected[i])
			return
		}
	}
}

func TestParseTickets02(t *testing.T) {
	txt := "your ticket:\n7,1,14"

	tix := parseTickets(txt)
	if len(tix) != 1 {
		t.Errorf("Expected 1 ticket, but got %d", len(tix))
	}

	expected := []ticket{[]int{7, 1, 14}}
	for i := range tix {
		if !intSliceEqual(tix[i], expected[i]) {
			t.Errorf("%v != %v", tix[i], expected[i])
			return
		}
	}
}

func TestTicketScanErrorRate01(t *testing.T) {
	ruleData := "class: 1-3 or 5-7\nrow: 6-11 or 33-44\nseat: 13-40 or 45-50"
	ticketsData := "7,3,47\n40,4,50\n55,2,20\n38,6,12"
	expectedErrRates := []int{0, 4, 55, 12}

	rules := parseRuleset(ruleData)
	tix := parseTickets(ticketsData)
	for i := range tix {
		actual, _ := ticketScanErrorRate(tix[i], &rules)
		if actual != expectedErrRates[i] {
			t.Errorf("%d != %d", expectedErrRates[i], actual)
			return
		}
	}

}

func TestFilterInvalidTix(t *testing.T) {
	ruleData := "class: 1-3 or 5-7\nrow: 6-11 or 33-44\nseat: 13-40 or 45-50"
	ticketsData := "7,3,47\n40,4,50\n55,2,20\n38,6,12"
	expected := []int{7, 3, 47}
	rules := parseRuleset(ruleData)
	tix := parseTickets(ticketsData)
	res := filterInvalidTickets(tix, &rules)
	if len(res) > 1 {
		t.Error("Did not filter all invalid tickets.")
		return
	} else if len(res) < 1 {
		t.Error("Filtered too many tickets.")
		return
	}

	validTicket := res[0]
	if !intSliceEqual(expected, validTicket) {
		t.Errorf("%v != %v", expected, validTicket)
	}
}

func TestPart1Integration01(t *testing.T) {
	txt := `class: 1-3 or 5-7
row: 6-11 or 33-44
seat: 13-40 or 45-50

your ticket:
7,1,14

nearby tickets:
7,3,47
40,4,50
55,2,20
38,6,12`
	res := part1(txt)
	if res != 71 {
		t.Errorf("71 != %d", res)
	}
}

func TestInvalidateValidityTable01(t *testing.T) {
	vr := makeValidityTable(4, 3)
	/*  before
	T T T
	T T T
	T T T
	T T T*/
	vr.invalidate(1, 2)
	/*  after
	T T T
	T T F
	T T T
	T T T*/
	for i := range vr {
		for j, val := range vr[i] {
			if i == 1 && j == 2 {
				if val {
					t.Errorf("(1,2) should be false. %v", vr)
					return
				}
			} else if !val {
				t.Errorf("All values other than (1,2) should be true. %v", vr)
				return
			}
		}
	}
}

func TestInvalidateValidityTable02(t *testing.T) {
	vr := makeValidityTable(4, 3)
	/*  before
	T T T
	T T T
	T T T
	T T T*/
	vr.invalidate(1, 2)
	vr.invalidate(1, 0)
	/*  after
	T T T
	F T F
	T T T
	T T T
	rule index 1 is for sure linked to row 1,
	so the other rows should be updated accordingly:
	T F T
	F T F
	T F T
	T F T*/
	expected := []bool{true, false, true, false, true, false, true, false, true, true, false, true}
	for i := range vr {
		for _, val := range vr[i] {
			if val != expected[0] {
				t.Errorf("validityTable does not match expected values. %v", vr)
				return
			}
			expected = expected[1:]
		}
	}
}

func TestInvalidTicketExists(t *testing.T) {
	myrng1 := rng{from: 1, to: 4}
	myrng2 := rng{from: 10, to: 14}
	r := rule{name: "class", ranges: []rng{myrng1, myrng2}}
	ticketsData := "1,1,1\n1,6,1\n4,11,14\n11,3,12"
	tix := parseTickets(ticketsData)
	if invalidTicketExists(&tix, 0, r) {
		t.Errorf("Field index 0 (1,1,4,11) should be valid for all tickets.")
	}
	if !invalidTicketExists(&tix, 1, r) {
		t.Errorf("Field index 1 should be invalid for ticket 1 (6 is outside of allowed ranges)")
	}
	if invalidTicketExists(&tix, 2, r) {
		t.Errorf("Field index 2 (1,1,14,12) should be valid for all tickets.")
	}

}

func TestProcessOfElimination(t *testing.T) {
	ruleData := "class: 0-1 or 4-19\nrow: 0-5 or 8-19\nseat: 0-13 or 16-19"
	ticketData := "11,12,13\n3,9,18\n15,1,5\n5,14,9"
	rules := parseRuleset(ruleData)
	tix := parseTickets(ticketData)
	vt := makeValidityTable(3, 3)
	vt = processOfElimination(vt, &tix, rules)
	/*expected:
	field 0 must be row
	field 1 must be class
	field 2 must be seat

	0 F T F
	1 T F F
	2 F F T
	*/
	expected := []bool{false, true, false, true, false, false, false, false, true}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if vt[i][j] != expected[0] {
				t.Errorf("Incorrect validty table value. %v", vt)
				return
			}
			expected = expected[1:]
		}
	}
}
