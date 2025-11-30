# Bit Manipulation

Bit manipulation involves manipulating individual bits of a number. It's common in low-level programming and optimization problems.

## Bitwise Operators

- `&` (AND): 1 if both bits are 1.
- `|` (OR): 1 if either bit is 1.
- `^` (XOR): 1 if bits are different.
- `~` (NOT): Inverts bits (0 -> 1, 1 -> 0).
- `<<` (Left Shift): Shifts bits to the left, filling with 0. Equivalent to multiplying by $2^k$.
- `>>` (Right Shift): Shifts bits to the right.

## Common Tricks

- **Get Bit**: `(num & (1 << i)) != 0`
- **Set Bit**: `num | (1 << i)`
- **Clear Bit**: `num & ^(1 << i)`
- **Update Bit**: Clear bit, then set to value.
- **Clear Bits MSB through i**: `num & ((1 << i) - 1)`
- **Clear Bits i through 0**: `num & ^((1 << (i + 1)) - 1)`

## Two's Complement

Negative numbers are stored using Two's Complement.

1.  Invert bits.
2.  Add 1.
    Example: -1 is all 1s (`111...1`).

## Go Implementation

Go supports all standard bitwise operators.

```go
x := 1 << 2 // 4 (0100)
y := x | 1  // 5 (0101)
```
