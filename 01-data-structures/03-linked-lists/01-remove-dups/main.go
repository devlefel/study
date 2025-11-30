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

// RemoveDups removes duplicate nodes from an unsorted linked list.
// Write code to remove duplicates from an unsorted linked list.
//
// Examples:
// Input: 1 -> 2 -> 1 -> 3 -> nil, Output: 1 -> 2 -> 3 -> nil
// Input: 1 -> 1 -> 1 -> nil, Output: 1 -> nil
// Input: 1 -> 2 -> 3 -> nil, Output: 1 -> 2 -> 3 -> nil
func RemoveDups(head *Node) *Node {
	// TODO: Implement this function
	return head
}

// Helper to print list
func printList(head *Node) string {
	if head == nil {
		return "nil"
	}
	return fmt.Sprintf("%d -> %s", head.Data, printList(head.Next))
}

// Helper to create list
func createList(nums []int) *Node {
	if len(nums) == 0 {
		return nil
	}
	head := &Node{Data: nums[0]}
	current := head
	for i := 1; i < len(nums); i++ {
		current.Next = &Node{Data: nums[i]}
		current = current.Next
	}
	return head
}

// Expected Time Complexity: O(N)
// Expected Space Complexity: O(N) (with buffer) or O(1) (without buffer, O(N^2) time)
// Thresholds:
//   Time:   Low < 5µs,   Medium < 50µs,   High > 500µs
//   Memory: Low < 1KB,   Medium < 10KB,   High > 100KB

func main() {
	// Test Cases
	testCases := []struct {
		input    []int
		expected []int
	}{
		{[]int{1, 2, 1, 3}, []int{1, 2, 3}},
		{[]int{1, 1, 1}, []int{1}},
		{[]int{1, 2, 3}, []int{1, 2, 3}},
	}

	// Profiling
	fmt.Println("\n--- Profiling ---")
	var m1, m2 runtime.MemStats
	runtime.ReadMemStats(&m1)
	start := time.Now()

	for _, tc := range testCases {
		head := createList(tc.input)
		RemoveDups(head)

		// Verify
		current := head
		status := "PASS"
		i := 0
		for current != nil {
			if i >= len(tc.expected) || current.Data != tc.expected[i] {
				status = "FAIL"
				break
			}
			current = current.Next
			i++
		}
		if i != len(tc.expected) {
			status = "FAIL"
		}

		result := printList(head)
		fmt.Printf("%s: %v -> %s (Expected: %v)\n", status, tc.input, result, tc.expected)
	}

	duration := time.Since(start)
	runtime.ReadMemStats(&m2)
	memUsage := m2.TotalAlloc - m1.TotalAlloc

	fmt.Printf("Input Length: %d\n", len(testCases))
	fmt.Printf("Execution Time: %v\n", duration)
	fmt.Printf("Memory Usage: %d bytes\n", memUsage)
}
