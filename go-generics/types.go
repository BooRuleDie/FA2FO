package main

// Numeric is a generic interface that represents all numeric types including their defined aliases
type Numeric interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64
}

// CustomInt is an alias type for int
type CustomInt int

// NumberContainer is a generic struct that holds a numeric value and associated metadata
type NumberContainer[T Numeric] struct {
	Value T     // The numeric value
	Meta  string // Associated metadata
}

// NumberMap is a generic map type with comparable keys and numeric values
type NumberMap[K comparable, V Numeric] map[K]V
