package main

import (
	"fmt"
	"runtime"
	"time"
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
	// Test Cases
	type testCase struct {
		name     string
		setup    func() (*Node, *Node)
		expected bool
	}

	testCases := []testCase{
		{
			"Route A -> C",
			func() (*Node, *Node) {
				a := &Node{Name: "A"}
				b := &Node{Name: "B"}
				c := &Node{Name: "C"}
				d := &Node{Name: "D"}
				e := &Node{Name: "E"}

				a.Neighbors = []*Node{b, d}
				b.Neighbors = []*Node{c}
				d.Neighbors = []*Node{e}
				return a, c
			},
			true,
		},
		{
			"Route C -> A",
			func() (*Node, *Node) {
				a := &Node{Name: "A"}
				b := &Node{Name: "B"}
				c := &Node{Name: "C"}
				d := &Node{Name: "D"}
				e := &Node{Name: "E"}

				a.Neighbors = []*Node{b, d}
				b.Neighbors = []*Node{c}
				d.Neighbors = []*Node{e}
				return c, a
			},
			false,
		},
	}

	// Profiling
	fmt.Println("\n--- Profiling ---")
	var m1, m2 runtime.MemStats
	runtime.ReadMemStats(&m1)
	start := time.Now()

	for _, tc := range testCases {
		startNode, endNode := tc.setup()
		result := RouteBetweenNodes(startNode, endNode)
		
		status := "FAIL"
		if result == tc.expected {
			status = "PASS"
		}
		fmt.Printf("%s: '%s' -> '%v' (Expected: '%v')\n", status, tc.name, result, tc.expected)
	}

	duration := time.Since(start)
	runtime.ReadMemStats(&m2)
	memUsage := m2.TotalAlloc - m1.TotalAlloc

	fmt.Printf("Input Length: %d\n", len(testCases))
	fmt.Printf("Execution Time: %v\n", duration)
	fmt.Printf("Memory Usage: %d bytes\n", memUsage)
}
