# Recursion and Dynamic Programming

## Recursion

A method where the solution to a problem depends on solutions to smaller instances of the same problem.

- **Base Case**: The condition that stops the recursion.
- **Recursive Step**: The part where the function calls itself.

## Dynamic Programming (DP)

DP is an optimization over plain recursion. It's used when we have **overlapping subproblems**.

- **Top-Down (Memoization)**: Store the results of recursive calls in a map or array so we don't recompute them.
- **Bottom-Up**: Build the solution from the smallest subproblems up to the target.

## Common Patterns

1.  **Fibonacci**: $F(n) = F(n-1) + F(n-2)$.
2.  **Grid Paths**: Number of ways to reach $(r, c)$ is usually ways to $(r-1, c)$ + ways to $(r, c-1)$.
3.  **Knapsack Problem**: Maximize value given weight constraint.
