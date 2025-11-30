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
	// Note: DeleteMiddleNode requires access to the node to be deleted.
	// We'll simulate this by creating a list and picking a node.
	type testCase struct {
		input    []int
		delIndex int // 0-based index of node to delete
		expected []int
	}
	testCases := []testCase{
		{[]int{1, 2, 3, 4, 5}, 2, []int{1, 2, 4, 5}}, // Delete 3
		{[]int{1, 2, 3}, 1, []int{1, 3}},             // Delete 2
	}

	// Profiling
	fmt.Println("\n--- Profiling ---")
	var m1, m2 runtime.MemStats
	runtime.ReadMemStats(&m1)
	start := time.Now()

	for _, tc := range testCases {
		// Create list
		var head *Node
		if len(tc.input) > 0 {
			head = &Node{Data: tc.input[0]}
			curr := head
			var nodeToDelete *Node
			if tc.delIndex == 0 {
				nodeToDelete = head
			}
			for i := 1; i < len(tc.input); i++ {
				newNode := &Node{Data: tc.input[i]}
				curr.Next = newNode
				curr = newNode
				if i == tc.delIndex {
					nodeToDelete = newNode
				}
			}

			if nodeToDelete != nil {
				DeleteMiddleNode(nodeToDelete)
			}
		}

		// Verify
		status := "PASS"
		curr := head
		i := 0
		for curr != nil {
			if i >= len(tc.expected) || curr.Data != tc.expected[i] {
				status = "FAIL"
				break
			}
			curr = curr.Next
			i++
		}
		if i != len(tc.expected) {
			status = "FAIL"
		}

		result := printList(head)
		fmt.Printf("%s: %v (del index %d) -> %s (Expected: %v)\n", status, tc.input, tc.delIndex, result, tc.expected)
	}

	duration := time.Since(start)
	runtime.ReadMemStats(&m2)
	memUsage := m2.TotalAlloc - m1.TotalAlloc

	fmt.Printf("Input Length: %d\n", len(testCases))
	fmt.Printf("Execution Time: %v\n", duration)
	fmt.Printf("Memory Usage: %d bytes\n", memUsage)
}
