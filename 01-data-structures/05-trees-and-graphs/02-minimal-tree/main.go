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

// MinimalTree creates a binary search tree with minimal height from a sorted (increasing order) array with unique integer elements.
//
// Examples:
// Input: [1, 2, 3] -> Root: 2, Left: 1, Right: 3 (Height 2)
// Input: [1, 2, 3, 4, 5, 6, 7] -> Root: 4, Left: 2, Right: 6 ... (Height 3)
func MinimalTree(arr []int) *TreeNode {
	// TODO: Implement this function
	return nil
}

// Helper to get height
func getHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := getHeight(root.Left)
	right := getHeight(root.Right)
	if left > right {
		return left + 1
	}
	return right + 1
}

// Expected Time Complexity: O(N)
// Expected Space Complexity: O(N) (Recursion stack + result tree)
// Thresholds:
//   Time:   Low < 10µs,   Medium < 100µs,   High > 1ms
//   Memory: Low < 1KB,   Medium < 10KB,   High > 100KB

func main() {
	// Test Cases
	// Input: [1, 2, 3, 4, 5, 6, 7]
	// Expected: Balanced Tree
	//       4
	//     /   \
	//    2     6
	//   / \   / \
	//  1   3 5   7
	input := []int{1, 2, 3, 4, 5, 6, 7}
	root := MinimalTree(input)
	
	// Verify height is minimal (ceil(log2(7+1)) = 3)
	status := "FAIL"
	if root != nil && root.Val == 4 && root.Left.Val == 2 && root.Right.Val == 6 {
		status = "PASS"
	}
	fmt.Printf("Test Case 1: %s (Root: %d)\n", status, root.Val)

	// Profiling
	fmt.Println("\n--- Profiling ---")
	largeInput := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		largeInput[i] = i
	}
	
	var m1, m2 runtime.MemStats
	runtime.ReadMemStats(&m1)
	start := time.Now()
	
	MinimalTree(largeInput)
	
	duration := time.Since(start)
	runtime.ReadMemStats(&m2)
	memUsage := m2.TotalAlloc - m1.TotalAlloc
	
	fmt.Printf("Input Size: 1000\n")
	fmt.Printf("Execution Time: %v\n", duration)
	fmt.Printf("Memory Usage: %d bytes\n", memUsage)
}
