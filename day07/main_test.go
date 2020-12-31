package main

import (
	"testing"

	"github.com/dliberat/algoimpl/go/graph"
)

func ruleParseTester(txt, expectedBagType string, expectedContents map[string]int, t *testing.T) {
	r := parseRule(txt)
	if r.bagType != expectedBagType {
		t.Errorf("bagType should be '%s', but got %s", expectedBagType, r.bagType)
	}
	if len(r.contents) != len(expectedContents) {
		t.Errorf("Light red bags should contain %d types of bags, but found %d", len(expectedContents), len(r.contents))
	}
	for k, v := range expectedContents {
		if r.contents[k] != v {
			t.Errorf("%s bags should be %d, but got %d", k, v, r.contents[k])
		}
	}
}

func TestParseRule01(t *testing.T) {
	txt := "light red bags contain 1 bright white bag, 2 muted yellow bags."
	expectedContents := map[string]int{
		"bright white": 1,
		"muted yellow": 2,
	}
	ruleParseTester(txt, "light red", expectedContents, t)
}

func TestParseRule02(t *testing.T) {
	txt := "dark orange bags contain 3 bright white bags, 4 muted yellow bags."
	expectedContents := map[string]int{
		"bright white": 3,
		"muted yellow": 4,
	}
	ruleParseTester(txt, "dark orange", expectedContents, t)
}

func TestParseRule03(t *testing.T) {
	txt := "bright white bags contain 1 shiny gold bag."
	expectedContents := map[string]int{
		"shiny gold": 1,
	}
	ruleParseTester(txt, "bright white", expectedContents, t)
}

func TestParseRule04(t *testing.T) {
	txt := "faded blue bags contain no other bags."
	expectedContents := map[string]int{}
	ruleParseTester(txt, "faded blue", expectedContents, t)
}

func TestMakeGraph(t *testing.T) {
	input := `light red bags contain 1 bright white bag, 2 muted yellow bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
bright white bags contain 1 shiny gold bag.
muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain 3 faded blue bags, 4 dotted black bags.
vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags.`
	rules := parseRuleList(input)
	var nodes map[string]graph.Node
	g := buildGraph(rules, &nodes)

	neighbors := g.Neighbors(nodes["light red"])
	if len(neighbors) != 2 {
		t.Errorf("Expected 2 neighbors but got %d", len(neighbors))
	}

	neighbors = g.Neighbors(nodes["bright white"])
	if len(neighbors) != 1 {
		t.Errorf("Expected 1 neighbor but got %d", len(neighbors))
	}

	neighbors = g.Neighbors(nodes["dotted black"])
	if len(neighbors) != 0 {
		t.Errorf("Expected 0 neighbors but got %d", len(neighbors))
	}
}

func TestPart1Integration01(t *testing.T) {
	input := `light red bags contain 1 bright white bag, 2 muted yellow bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
bright white bags contain 1 shiny gold bag.
muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain 3 faded blue bags, 4 dotted black bags.
vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags.`
	result := part1(input)
	if result != 4 {
		t.Errorf("Expected 4 paths to 'shiny gold' but got %d", result)
	}
}

func TestDfs01(t *testing.T) {
	input := `dotted black bags contain no other bags.`
	rules := parseRuleList(input)
	nodes := make(map[string]graph.Node, 0)
	g := buildGraph(rules, &nodes)

	startNode := nodes["dotted black"]
	totalBags := dfs(g, startNode, 1)
	totalBags-- // don't count the outermost bag
	if totalBags != 0 {
		t.Errorf("Dotted black bags contain no other bags, but got %d total bags", totalBags)
	}
}

func TestDfs02(t *testing.T) {
	input := `vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags.`

	rules := parseRuleList(input)
	nodes := make(map[string]graph.Node, 0)
	g := buildGraph(rules, &nodes)

	startNode := nodes["vibrant plum"]
	totalBags := dfs(g, startNode, 1)
	totalBags-- // don't count the outermost bag
	if totalBags != 11 {
		t.Errorf("5x faded blue + 6x dotted black = 11 but got %d", totalBags)
	}
}

func TestDfs03(t *testing.T) {
	input := `shiny gold bags contain 2 vibrant plum bags.
vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags.`

	rules := parseRuleList(input)
	nodes := make(map[string]graph.Node, 0)
	g := buildGraph(rules, &nodes)

	startNode := nodes["shiny gold"]
	totalBags := dfs(g, startNode, 1)
	totalBags-- // don't count the outermost bag
	if totalBags != 24 {
		t.Errorf("2 + 2*11 = 24 but got %d", totalBags)
	}
}

func TestDfs04(t *testing.T) {
	input := `shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain 3 faded blue bags, 4 dotted black bags.
vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags.`

	rules := parseRuleList(input)
	nodes := make(map[string]graph.Node, 0)
	g := buildGraph(rules, &nodes)

	startNode := nodes["shiny gold"]
	totalBags := dfs(g, startNode, 1)
	totalBags-- // don't count the outermost bag
	if totalBags != 32 {
		t.Errorf("1 + 1*7 + 2 + 2*11 = 32 but got %d", totalBags)
	}
}

func TestDfs05(t *testing.T) {
	input := `shiny gold bags contain 2 dark red bags.
dark red bags contain 2 dark orange bags.
dark orange bags contain 2 dark yellow bags.
dark yellow bags contain 2 dark green bags.
dark green bags contain 2 dark blue bags.
dark blue bags contain 2 dark violet bags.
dark violet bags contain no other bags.`

	rules := parseRuleList(input)
	nodes := make(map[string]graph.Node, 0)
	g := buildGraph(rules, &nodes)

	startNode := nodes["shiny gold"]
	totalBags := dfs(g, startNode, 1)
	totalBags-- // don't count the outermost bag
	if totalBags != 126 {
		t.Errorf("Expected 126 but got %d", totalBags)
	}
}
