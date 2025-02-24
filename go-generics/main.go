package main

import "fmt"

func main() {
	// Example with integers
	intSum := Sum(1, 2, 3, 4, 5)
	fmt.Printf("Sum of integers: %v\n", intSum)

	// Example with floats
	floatSum := Sum(1.5, 2.7, 3.2, 4.9, 5.1)
	fmt.Printf("Sum of floats: %v\n", floatSum)

	// Example with uint8
	uint8Sum := Sum(uint8(1), uint8(2), uint8(3))
	fmt.Printf("Sum of uint8: %v\n", uint8Sum)

	// Example with int64
	int64Sum := Sum(int64(100), int64(200), int64(300))
	fmt.Printf("Sum of int64: %v\n", int64Sum)

	// Example with CustomInt
	customIntSum := Sum(CustomInt(10), CustomInt(20), CustomInt(30))
	fmt.Printf("Sum of CustomInt: %v\n", customIntSum)

	// Test with int
	nc1 := NumberContainer[int]{Value: 42, Meta: "test"}
	fmt.Printf("NumberContainer with int: %+v\n", nc1)

	// Test with float64
	nc2 := NumberContainer[float64]{Value: 3.14, Meta: "pi"}
	fmt.Printf("NumberContainer with float64: %+v\n", nc2)

	// Test with string keys and int values
	nm1 := NumberMap[string, int]{"one": 1, "two": 2}
	fmt.Printf("NumberMap with string keys and int values: %v\n", nm1)

	// Test with int keys and float64 values
	nm2 := NumberMap[int, float64]{1: 1.1, 2: 2.2}
	fmt.Printf("NumberMap with int keys and float64 values: %v\n", nm2)
}
