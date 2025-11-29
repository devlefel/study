package main

import (
	"fmt"
	"runtime"
	"time"
)

type Node struct {
	Data int
	Next *Node
}

// DeleteMiddleNode deletes a node in the middle (i.e., any node but the first and last node, not necessarily the exact middle)
// of a singly linked list, given only access to that node.
//
// Examples:
// Input: Node c from the list a->b->c->d->e -> Output: nothing is returned, but the new list looks like a->b->d->e
func DeleteMiddleNode(node *Node) bool {
	// TODO: Implement this function
	return false
}

// Helper to print list
func printList(head *Node) string {
	if head == nil {
		return "nil"
	}
	return fmt.Sprintf("%s -> %s", string(rune(head.Data)), printList(head.Next))
}

// Expected Time Complexity: O(1) (given the node)
// Expected Space Complexity: O(1)
// Thresholds:
//   Time:   Low < 1µs,   Medium < 10µs,   High > 100µs
//   Memory: Low < 1KB,   Medium < 1KB,   High > 1KB

func main() {
	// Test Cases
	// 1 -> 2 -> 3 -> 4 -> 5
	node3 := &Node{Data: 3, Next: &Node{Data: 4, Next: &Node{Data: 5}}}
	head := &Node{Data: 1, Next: &Node{Data: 2, Next: node3}}
	
	// Delete node 3
	DeleteMiddleNode(node3)
	
	// Verify: 1 -> 2 -> 4 -> 5
	// Check if node 2 points to 4
	status := "FAIL"
	if head.Next.Next.Data == 4 {
		status = "PASS"
	}
	fmt.Printf("Test Case 1: %s\n", status)

	// Profiling
	fmt.Println("\n--- Profiling ---")
	// Setup dummy node to delete
	dummy := &Node{Data: 10, Next: &Node{Data: 20}}
	
	var m1, m2 runtime.MemStats
	runtime.ReadMemStats(&m1)
	start := time.Now()
	
	DeleteMiddleNode(dummy)
	
	duration := time.Since(start)
	runtime.ReadMemStats(&m2)
	memUsage := m2.TotalAlloc - m1.TotalAlloc
	
	fmt.Printf("Operation: Delete Node\n")
	fmt.Printf("Execution Time: %v\n", duration)
	fmt.Printf("Memory Usage: %d bytes\n", memUsage)
}
