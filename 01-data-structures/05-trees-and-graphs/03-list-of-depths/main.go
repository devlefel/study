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
	//       1
	//     /   \
	//    2     3
	//   /
	//  4
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 3}}
	
	lists := ListOfDepths(root)
	
	// Verify
	// Level 0: [1]
	// Level 1: [2, 3]
	// Level 2: [4]
	status := "FAIL"
	// The original ListOfDepths returns []*ListNode, so we check the Val field directly.
	// The provided verification logic seems to assume a different list structure (e.g., container/list.List).
	// Adapting to the existing []*ListNode return type.
	if len(lists) == 3 && lists[0] != nil && lists[0].Val == 1 &&
		lists[1] != nil && lists[1].Val == 2 && lists[1].Next != nil && lists[1].Next.Val == 3 &&
		lists[2] != nil && lists[2].Val == 4 {
		status = "PASS"
	}
	fmt.Printf("Test Case 1: %s (Levels: %d)\n", status, len(lists))

	// Profiling
	fmt.Println("\n--- Profiling ---")
	// Create a balanced tree of 1023 nodes (depth 10)
	largeRoot := createBalancedTree(0, 1023)
	
	var m1, m2 runtime.MemStats
	runtime.ReadMemStats(&m1)
	start := time.Now()
	
	ListOfDepths(largeRoot)
	
	duration := time.Since(start)
	runtime.ReadMemStats(&m2)
	memUsage := m2.TotalAlloc - m1.TotalAlloc
	
	fmt.Printf("Tree Size: 1023 Nodes\n")
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
