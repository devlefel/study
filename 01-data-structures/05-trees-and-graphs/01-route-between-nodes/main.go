package main

import (
	"fmt"
)

type Node struct {
	Name      string
	Neighbors []*Node
	Visited   bool // Helper for traversal
}

// RouteBetweenNodes determines if there is a route between two nodes in a directed graph.
// Given a directed graph, design an algorithm to find out whether there is a route between two nodes.
//
// Examples:
// A -> B -> C
// Route(A, C) -> true
// Route(C, A) -> false
func RouteBetweenNodes(start, end *Node) bool {
	// TODO: Implement this function (BFS is usually better for shortest path/existence)
	return false
}

func main() {
	// Create graph
	// A -> B -> C
	// A -> D
	// D -> E
	a := &Node{Name: "A"}
	b := &Node{Name: "B"}
	c := &Node{Name: "C"}
	d := &Node{Name: "D"}
	e := &Node{Name: "E"}

	a.Neighbors = []*Node{b, d}
	b.Neighbors = []*Node{c}
	d.Neighbors = []*Node{e}

	// Test A -> C (True)
	if RouteBetweenNodes(a, c) {
		fmt.Println("PASS: Route A -> C exists")
	} else {
		fmt.Println("FAIL: Route A -> C should exist")
	}

	// Reset visited
	a.Visited, b.Visited, c.Visited, d.Visited, e.Visited = false, false, false, false, false

	// Test C -> A (False)
	if !RouteBetweenNodes(c, a) {
		fmt.Println("PASS: Route C -> A does not exist")
	} else {
		fmt.Println("FAIL: Route C -> A should not exist")
	}
}
