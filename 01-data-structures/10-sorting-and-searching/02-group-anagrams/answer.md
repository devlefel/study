# Answer: Group Anagrams

## Problem

Write a method to sort an array of strings so that all the anagrams are next to each other.

## Solution Approach

We can use a **Hash Map** (or Dictionary) to group the strings.

1.  The key of the map should be the "sorted" version of the string (e.g., "cab" -> "abc").
2.  The value of the map should be a list of strings that match that key.
3.  Iterate through the array, sort each string, add it to the map.
4.  Iterate through the map and put the strings back into the array.

### Diagram

```text
Input: ["cat", "dog", "tac", "god", "act"]

"cat" -> Key "act". Map: {"act": ["cat"]}
"dog" -> Key "dgo". Map: {"act": ["cat"], "dgo": ["dog"]}
"tac" -> Key "act". Map: {"act": ["cat", "tac"], "dgo": ["dog"]}
"god" -> Key "dgo". Map: {"act": ["cat", "tac"], "dgo": ["dog", "god"]}
"act" -> Key "act". Map: {"act": ["cat", "tac", "act"], "dgo": ["dog", "god"]}

Result: ["cat", "tac", "act", "dog", "god"]
```

## Code Snippet

```go
import (
	"sort"
	"strings"
)

func GroupAnagrams(strs []string) []string {
	groups := make(map[string][]string)

	for _, s := range strs {
		key := sortString(s)
		groups[key] = append(groups[key], s)
	}

	var result []string
	for _, group := range groups {
		result = append(result, group...)
	}
	return result
}

func sortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
```

## Complexity Analysis

- **Time Complexity**: O(N \* K log K), where N is number of strings and K is max length of a string (sorting each string).
- **Space Complexity**: O(N \* K) to store the map.
