package main

import ("fmt")

func addEdge(graph map[string]map[string]bool, from string, to string) {
	edges := graph[from]
	if edges == nil {
		edges = make(map[string]bool)
		graph[from] = edges
	}
	edges[to] = true
}

func hasEdge(graph map[string]map[string]bool, from string, to string) bool {
	return graph[from][to]
}

func main() {
	g := make(map[string]map[string]bool)

	addEdge(g, "a", "b")
	fmt.Println(hasEdge(g, "a", "b"))
	fmt.Println(hasEdge(g, "a", "c"))
}
