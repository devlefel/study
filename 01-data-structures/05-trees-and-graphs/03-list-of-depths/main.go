package main

import (
	"fmt"
	"runtime"
	"time"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// ListOfDepths creates a linked list of all the nodes at each depth of a binary tree.
// If you have a tree with depth D, you'll have D linked lists.
//
// Examples:
// Input:
//     1
//    / \
//   2   3
// Output: [1->nil, 2->3->nil]
func ListOfDepths(root *TreeNode) []*ListNode {
	// TODO: Implement this function
	return nil
}

// Expected Time Complexity: O(N)
// Expected Space Complexity: O(N)
// Thresholds:
//   Time:   Low < 10µs,   Medium < 100µs,   High > 1ms
//   Memory: Low < 1KB,   Medium < 10KB,   High > 100KB

func main() {
	// Test Cases
	type testCase struct {
		name     string
		setup    func() *TreeNode
		expected int // Number of levels
	}

	testCases := []testCase{
		{
			"Balanced Tree Depth 3",
			func() *TreeNode {
				return &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 3}}
			},
			3,
		},
	}

	// Profiling
	fmt.Println("\n--- Profiling ---")
	var m1, m2 runtime.MemStats
	runtime.ReadMemStats(&m1)
	start := time.Now()

	for _, tc := range testCases {
		root := tc.setup()
		lists := ListOfDepths(root)
		
		status := "FAIL"
		if len(lists) == tc.expected {
			status = "PASS"
		}
		fmt.Printf("%s: '%s' -> %d Levels (Expected: %d)\n", status, tc.name, len(lists), tc.expected)
	}

	duration := time.Since(start)
	runtime.ReadMemStats(&m2)
	memUsage := m2.TotalAlloc - m1.TotalAlloc

	fmt.Printf("Input Length: %d\n", len(testCases))
	fmt.Printf("Execution Time: %v\n", duration)
	fmt.Printf("Memory Usage: %d bytes\n", memUsage)
}

func createBalancedTree(start, end int) *TreeNode {
	if end < start {
		return nil
	}
	mid := (start + end) / 2
	node := &TreeNode{Val: mid}
	node.Left = createBalancedTree(start, mid-1)
	node.Right = createBalancedTree(mid+1, end)
	return node
}
