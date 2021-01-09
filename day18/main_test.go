package main

import "testing"

func TestParse01(t *testing.T) {
	txt := "1"
	expr := parse(txt)
	expr = expr.exp[0]
	if expr.raw != "1" {
		t.Errorf("Incorrect raw value: '%s'", expr.raw)
	}
	if expr.value != 1 {
		t.Errorf("Incorrect calculated value: '%d'", expr.value)
	}
	if len(expr.exp) != 0 {
		t.Errorf("Incorrect number of subexpressions: %d", len(expr.exp))
	}
	if len(expr.ops) != 0 {
		t.Errorf("Incorrect number of operations: %d", len(expr.ops))
	}
}

func TestParse02(t *testing.T) {
	txt := "(5)"
	expr := parse(txt)
	expr = expr.exp[0].exp[0]
	if expr.raw != "5" {
		t.Errorf("Incorrect raw value: '%s'", expr.raw)
	}
	if expr.value != 5 {
		t.Errorf("Incorrect calculated value: '%d'", expr.value)
	}
	if len(expr.exp) != 0 {
		t.Errorf("Incorrect number of subexpressions: %d", len(expr.exp))
	}
	if len(expr.ops) != 0 {
		t.Errorf("Incorrect number of operations: %d", len(expr.ops))
	}
}

func TestParse03(t *testing.T) {
	txt := "1 + 1"
	expr := parse(txt)
	if expr.raw != "1 + 1" {
		t.Errorf("Incorrect raw value: '%s'", expr.raw)
	}
	if len(expr.exp) != 2 {
		t.Errorf("Incorrect number of subexpressions: %d", len(expr.exp))
	}
	if len(expr.ops) != 1 {
		t.Errorf("Incorrect number of operations: %d", len(expr.ops))
	}
	var oper opAdd = "+"
	if expr.ops[0] != oper {
		t.Errorf("Incorrectly parsed operator")
	}
}

func TestParse04(t *testing.T) {
	txt := "1 + ((2 * 3) + 4)"

	expr := parse(txt) // 1 + ((2 * 3) + 4)
	if len(expr.exp) != 2 {
		t.Errorf("Incorrect number of subexpressions: %s", expr.raw)
		return
	}
	expr = expr.exp[1] // (2 * 3) + 4
	if len(expr.exp) != 2 {
		t.Errorf("Incorrect number of subexpressions: %s", expr.raw)
		return
	}

	expr = expr.exp[0] // 2 * 3
	if len(expr.exp) != 2 {
		t.Errorf("Incorrect number of subexpressions: %s", expr.raw)
		return
	}

	lhs := expr.exp[0]
	if lhs.value != 2 {
		t.Errorf("Incorrect value for left hand side: %d", lhs.value)
	}
	rhs := expr.exp[1]
	if rhs.value != 3 {
		t.Errorf("Incorrect value for right hand side: %d", rhs.value)
	}
}

func TestReduce01(t *testing.T) {
	txt := "1"
	expected := 1
	ex := parse(txt)
	res := ex.reduce()
	if res != expected {
		t.Errorf("Bad reduction: %d != %d", expected, res)
	}
}

func TestReduce02(t *testing.T) {
	txt := "1 + 1"
	expected := 2
	ex := parse(txt)
	res := ex.reduce()
	if res != expected {
		t.Errorf("Bad reduction: %d != %d", expected, res)
	}
}

func TestReduce03(t *testing.T) {
	txt := "2 * (1 + 1)"
	expected := 4
	ex := parse(txt)
	res := ex.reduce()
	if res != expected {
		t.Errorf("Bad reduction: %d != %d", expected, res)
	}
}

func TestReduce04(t *testing.T) {
	txt := "(1 + 1) * 3"
	expected := 6
	ex := parse(txt)
	res := ex.reduce()
	if res != expected {
		t.Errorf("Bad reduction: %d != %d", expected, res)
	}
}

func TestReducePart201(t *testing.T) {
	txt := "2 * 3 + (4 * 5)"
	expected := 46
	ex := parse(txt)
	res := ex.reducePart2()
	if res != expected {
		t.Errorf("Bad reduction: %d != %d", expected, res)
	}
}

func TestReducePart202(t *testing.T) {
	txt := "1 + (2 * 3) + (4 * (5 + 6))"
	expected := 51
	ex := parse(txt)
	res := ex.reducePart2()
	if res != expected {
		t.Errorf("Bad reduction: %d != %d", expected, res)
	}
}

func TestCalculate01(t *testing.T) {
	txt := []string{
		"1 + 2 * 3 + 4 * 5 + 6",
		"1 + (2 * 3) + (4 * (5 + 6))",
		"2 * 3 + (4 * 5)",
		"5 + (8 * 3 + 9 + 3 * 4 * 3)",
		"5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))",
		"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2",
	}
	expected := []int{
		71,
		51,
		26,
		437,
		12240,
		13632,
	}

	for i := range txt {
		parsed := parse(txt[i])
		ex := expected[i]
		res := parsed.reduce()
		if res != ex {
			t.Errorf("%s = %d != %d", txt[i], res, ex)
		}
	}
}

func TestCalculatePart2(t *testing.T) {
	txt := []string{
		"1 + (2 * 3) + (4 * (5 + 6))",
		"2 * 3 + (4 * 5)",
		"5 + (8 * 3 + 9 + 3 * 4 * 3)",
		"5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))",
		"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2",
	}
	expected := []int{
		51,
		46,
		1445,
		669060,
		23340,
	}

	for i := range txt {
		parsed := parse(txt[i])
		ex := expected[i]
		res := parsed.reducePart2()
		if res != ex {
			t.Errorf("%s = %d != %d", txt[i], res, ex)
		}
	}
}
