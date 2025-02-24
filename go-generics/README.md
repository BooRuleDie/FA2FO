# Understanding Generics in Go

Generics in Go, introduced in Go 1.18, allow you to write code that can work with multiple types while maintaining type safety. Instead of writing separate implementations for different types, you can write a single generic implementation that works across various types.

## Why Use Generics?

1. **Type Safety**: Generics provide compile-time type checking while allowing flexible type usage
2. **Code Reusability**: Write functions and data structures that work with multiple types
3. **Reduce Code Duplication**: Avoid writing similar code for different types
4. **Performance**: No runtime type assertions needed compared to using interface{}

## Generic Functions

```go
func Min[T constraints.Ordered](x, y T) T {
    if x < y {
        return x
    }
    return y
}

// Usage
minInt := Min[int](2, 3)
minFloat := Min[float64](2.1, 1.7)
```

## Generic Data Structures

### Maps
```go
// Generic map declaration
type Map[K comparable, V any] map[K]V

// Usage
scores := Map[string, int]{
    "Alice": 95,
    "Bob": 89,
}
```

### Slices
```go
// Generic slice type
type List[T any] []T

// Usage
numbers := List[int]{1, 2, 3}
```

### Structs
```go
// Generic struct
type Pair[T any] struct {
    First  T
    Second T
}

// Usage
point := Pair[int]{
    First:  10,
    Second: 20,
}
```

## Type Constraints

Go uses interfaces to define type constraints for generics:

```go
type Number interface {
    int | int64 | float64
}

func Sum[T Number](items []T) T {
    var sum T
    for _, item := range items {
        sum += item
    }
    return sum
}
```

## Best Practices

1. Use generics when you need to write similar code for multiple types
2. Consider type constraints carefully to ensure type safety
3. Don't overuse generics where simple concrete types would suffice
4. Use meaningful type parameter names (e.g., T, K, V)

Generics in Go provide a powerful way to write flexible, type-safe code while maintaining Go's emphasis on simplicity and performance.