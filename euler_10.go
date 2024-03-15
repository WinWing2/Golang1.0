package main

// Ищет сумму простых чисел, которые <= limit
func euler_10(limit int) int {
	sum := 0
	simpleDividersSlice := findPrimeNumbers(limit)
	for i := range simpleDividersSlice {
		sum += simpleDividersSlice[i]
	}
	return sum
}
