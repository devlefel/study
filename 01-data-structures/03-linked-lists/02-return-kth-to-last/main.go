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

// ReturnKthToLast finds the kth to last element of a singly linked list.
// Implement an algorithm to find the kth to last element of a singly linked list.
// Assume k = 1 is the last element, k = 2 is the second to last, etc.
//
// Examples:
// Input: 1 -> 2 -> 3 -> 4 -> 5, k = 2 -> Output: 4
// Input: 1 -> 2 -> 3, k = 1 -> Output: 3
// Input: 1 -> 2 -> 3, k = 3 -> Output: 1
func ReturnKthToLast(head *Node, k int) int {
	// TODO: Implement this function
	return -1
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

func main() {
	// Test Cases
	testCases := []struct {
		input    []int
		k        int
		expected int
	}{
		{[]int{1, 2, 3, 4, 5}, 2, 4},
		{[]int{1, 2, 3}, 1, 3},
		{[]int{1, 2, 3}, 3, 1},
	}

	// Profiling
	fmt.Println("\n--- Profiling ---")
	largeHead := &Node{Data: 0}
	curr := largeHead
	for i := 0; i < 1000; i++ {
		curr.Next = &Node{Data: i}
		curr = curr.Next
	}

	var m1, m2 runtime.MemStats
	runtime.ReadMemStats(&m1)
	start := time.Now()

	for _, tc := range testCases {
		head := createList(tc.input)
		result := ReturnKthToLast(head, tc.k)
		
		status := "FAIL"
		if result == tc.expected {
			status = "PASS"
		}
		fmt.Printf("%s: %v (k=%d) -> %d (Expected: %d)\n", status, tc.input, tc.k, result, tc.expected)
	}

	duration := time.Since(start)
	runtime.ReadMemStats(&m2)
	memUsage := m2.TotalAlloc - m1.TotalAlloc

	fmt.Printf("Input Length: %d\n", len(testCases))
	fmt.Printf("Execution Time: %v\n", duration)
	fmt.Printf("Memory Usage: %d bytes\n", memUsage)
}
