package main

func Sum[T Numeric](nums ...T) T {
	var sum T
	for _, num := range nums {
		sum += num
	}
	return sum
}
