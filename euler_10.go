package main

// Ищет сумму простых чисел, которые <= limit
func euler_10(limit int) int {
	sum := 0
	var simpleDividersSlice []int
	for i := 1; i <= limit; i++ {
		simpleDividersSlice = findSimpleDividers(i)
		if len(simpleDividersSlice) == 1 {
			sum += i
		}
	}
	return sum
}
