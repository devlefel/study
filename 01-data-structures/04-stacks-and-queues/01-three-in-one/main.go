package main

import (
	"fmt"
	"runtime"
	"time"
)

// ThreeInOne describes how you could use a single array to implement three stacks.
// Implement a data structure that uses a single array to implement three stacks.
//
// Examples:
// stack.Push(0, 10) -> Stack 0 has [10]
// stack.Push(1, 20) -> Stack 1 has [20]
// stack.Push(0, 11) -> Stack 0 has [10, 11]
// stack.Pop(0) -> Returns 11, Stack 0 has [10]
type ThreeStacks struct {
	// TODO: Add fields
}

func NewThreeStacks(stackSize int) *ThreeStacks {
	// TODO: Implement
	return &ThreeStacks{}
}

func (ts *ThreeStacks) Push(stackNum int, value int) error {
	// TODO: Implement
	return nil
}

func (ts *ThreeStacks) Pop(stackNum int) (int, error) {
	// TODO: Implement
	return 0, nil
}

func (ts *ThreeStacks) Peek(stackNum int) (int, error) {
	// TODO: Implement
	return 0, nil
}

func (ts *ThreeStacks) IsEmpty(stackNum int) bool {
	// TODO: Implement
	return true
}

// Expected Time Complexity: O(1) for push/pop
// Expected Space Complexity: O(N) for the array
// Thresholds:
//   Time:   Low < 1µs,   Medium < 10µs,   High > 100µs
//   Memory: Low < 1KB,   Medium < 10KB,   High > 100KB

func main() {
	// Test Cases
	stack := NewThreeStacks(3)
	
	// Push to stack 0
	stack.Push(0, 10)
	stack.Push(0, 11)
	
	// Push to stack 1
	stack.Push(1, 20)
	
	// Pop
	val1, _ := stack.Pop(0) // 11
	val2, _ := stack.Pop(1) // 20
	
	status := "FAIL"
	if val1 == 11 && val2 == 20 {
		status = "PASS"
	}
	fmt.Printf("Test Case 1: %s (Got %d, %d)\n", status, val1, val2)

	// Profiling
	fmt.Println("\n--- Profiling ---")
	largeStack := NewThreeStacks(1000)
	
	var m1, m2 runtime.MemStats
	runtime.ReadMemStats(&m1)
	start := time.Now()
	
	for i := 0; i < 1000; i++ {
		largeStack.Push(0, i)
	}
	for i := 0; i < 1000; i++ {
		largeStack.Pop(0)
	}
	
	duration := time.Since(start)
	runtime.ReadMemStats(&m2)
	memUsage := m2.TotalAlloc - m1.TotalAlloc
	
	fmt.Printf("Operations: 1000 Push + 1000 Pop\n")
	fmt.Printf("Execution Time: %v\n", duration)
	fmt.Printf("Memory Usage: %d bytes\n", memUsage)
}
