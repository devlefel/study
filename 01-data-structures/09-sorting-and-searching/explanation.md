# Sorting and Searching

## Common Sorting Algorithms

- **Bubble Sort**: Swap adjacent elements. $O(N^2)$.
- **Selection Sort**: Find min, swap with start. $O(N^2)$.
- **Merge Sort**: Divide in half, sort halves, merge. $O(N \log N)$. Stable.
- **Quick Sort**: Pick pivot, partition around pivot, recurse. $O(N \log N)$ avg, $O(N^2)$ worst. Unstable.
- **Radix Sort**: Sort by digit. $O(KN)$.

## Searching

- **Binary Search**: Search in sorted array by checking middle. $O(\log N)$.

## Go Implementation

Go's `sort` package provides efficient sorting.

```go
sort.Ints(arr) // Sorts an int slice
sort.Strings(arr) // Sorts a string slice
sort.SearchInts(arr, x) // Binary search
```
