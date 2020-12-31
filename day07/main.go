package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/dliberat/algoimpl/go/graph"
)

type rule struct {
	bagType  string
	contents map[string]int
}

func readInput(fname string) string {
	data, err := ioutil.ReadFile(fname)
	if err != nil {
		panic("Cannot read input data.")
	}
	s := string(data)
	return strings.TrimRight(s, "\n")
}

func parseRuleList(txt string) []rule {
	lines := strings.Split(txt, "\n")
	rules := make([]rule, len(lines))
	for i, line := range lines {
		rules[i] = parseRule(line)
	}
	return rules
}

func parseRule(txt string) rule {
	split := strings.Split(txt, " bags contain ")
	lhs, rhs := split[0], split[1]
	contentRules := strings.Split(rhs, ", ")
	contents := make(map[string]int, 0)
	for _, contentRule := range contentRules {
		if contentRule == "no other bags." {
			break
		}
		color, num := parseContentRule(contentRule)
		contents[color] = num
	}
	return rule{bagType: lhs, contents: contents}
}

func parseContentRule(txt string) (string, int) {
	firstSpace := strings.Index(txt, " ")
	lastSpace := strings.LastIndex(txt, " ")
	num, err := strconv.Atoi(txt[0:firstSpace])
	if err != nil {
		panic(fmt.Sprintf("Cannot parse content rule: %s", txt))
	}
	color := txt[firstSpace+1 : lastSpace]
	return color, num
}

func buildGraph(rules []rule, nodes *map[string]graph.Node) *graph.Graph {
	g := graph.New(graph.Directed)
	*nodes = make(map[string]graph.Node, 0)
	for _, r := range rules {
		from, ok := (*nodes)[r.bagType]
		if !ok {
			from = g.MakeNode()
			(*nodes)[r.bagType] = from
		}

		for key, val := range r.contents {
			to, ok := (*nodes)[key]
			if !ok {
				to = g.MakeNode()
				(*nodes)[key] = to
			}
			err := g.MakeEdgeWeight(from, to, val)
			if err != nil {
				fmt.Println(err, "|", r.bagType, "->", key)
			}
		}
	}
	// Make references back to the string values
	for key, node := range *nodes {
		*node.Value = key
	}
	return g
}

func pathExists(from, to graph.Node, g *graph.Graph) bool {
	paths := g.DijkstraSearch(from)
	for _, path := range paths {
		if len(path.Path) < 1 {
			continue // disjoint
		}
		if path.Weight == 0 {
			continue // path to self
		}
		e := path.Path[len(path.Path)-1]
		if e.End == to {
			return true
		}
	}
	return false
}

func part1(txt string) int {
	rules := parseRuleList(txt)
	nodes := make(map[string]graph.Node, 0)
	g := buildGraph(rules, &nodes)

	targetNode := nodes["shiny gold"]
	totalPaths := 0

	for _, node := range nodes {
		if node == targetNode {
			continue
		}
		if pathExists(node, targetNode, g) {
			totalPaths++
		}
	}
	return totalPaths
}

func dfs(g *graph.Graph, currentNode graph.Node, factor int) int {
	neighbors := g.Neighbors(currentNode)
	if len(neighbors) == 0 {
		return factor
	}

	totalBags := 0

	for _, neighbor := range neighbors {
		edge, err := g.GetEdge(currentNode, neighbor)
		if err != nil {
			panic("there should always be an edge between a node and its neighbors")
		}
		bags := dfs(g, neighbor, factor*edge.Weight)
		totalBags += bags
	}

	return totalBags + factor
}

func part2(txt string) int {
	rules := parseRuleList(txt)
	nodes := make(map[string]graph.Node, 0)
	g := buildGraph(rules, &nodes)

	startNode := nodes["shiny gold"]
	totalBags := dfs(g, startNode, 1)
	return totalBags - 1 // don't count the outermost shiny gold bag
}

func main() {
	txt := readInput("input.txt")
	p1 := part1(txt)
	fmt.Println("[PART 1] Result:", p1)
	p2 := part2(txt)
	fmt.Println("[PART 2] Result:", p2)
}
