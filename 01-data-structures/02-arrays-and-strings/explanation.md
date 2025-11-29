# Arrays and Strings

## Hash Tables

A hash table maps keys to values for highly efficient lookup.

- **Insertion**: $O(1)$ on average.
- **Lookup**: $O(1)$ on average.
- **Worst Case**: $O(N)$ if many collisions occur (rare with good hash functions).

**Go Implementation**: Go's built-in `map` type is a hash table.

```go
m := make(map[string]int)
m["key"] = 10
val, exists := m["key"]
```

## ArrayList (Slice in Go)

An array that resizes itself as needed.

- **Access**: $O(1)$.
- **Append**: $O(1)$ amortized.

**Go Implementation**: Go's `slice` acts like an ArrayList.

```go
s := []int{1, 2, 3}
s = append(s, 4) // Resizes if capacity is reached
```

## StringBuilder (strings.Builder in Go)

Concatenating strings in a loop is $O(N^2)$ because strings are immutable (a new copy is created each time). `StringBuilder` avoids this by keeping a resizable array of all strings, copying them back to a string only when necessary.

**Go Implementation**:

```go
var sb strings.Builder
for i := 0; i < 100; i++ {
    sb.WriteString("a")
}
result := sb.String()
```

## Common Techniques

1.  **Two Pointers**: One starting at the beginning, one at the end (or moving at different speeds).

    ```go
    func isPalindrome(s string) bool {
        left, right := 0, len(s)-1
        for left < right {
            if s[left] != s[right] {
                return false
            }
            left++
            right--
        }
        return true
    }
    ```

2.  **Sliding Window**: Maintain a window of elements that satisfies a condition.

    ```go
    func maxSubArraySum(arr []int, k int) int {
        maxSum, windowSum := 0, 0
        for i := 0; i < k; i++ {
            windowSum += arr[i]
        }
        maxSum = windowSum
        for i := k; i < len(arr); i++ {
            windowSum += arr[i] - arr[i-k]
            if windowSum > maxSum {
                maxSum = windowSum
            }
        }
        return maxSum
    }
    ```

3.  **Hash Map for Frequency**: Count occurrences of characters/elements.

    ```go
    func charFrequency(s string) map[rune]int {
        freq := make(map[rune]int)
        for _, char := range s {
            freq[char]++
        }
        return freq
    }
    ```
