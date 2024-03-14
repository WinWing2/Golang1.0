package main

import "math"

// Считает разницу меж квадратом суммы и суммой квадратов всех чисел от 0 до N, проверять на 10, = 2640
func euler_6(limit int) int {
	squareOfSums := 0
	sumOfSquares := 0
	for i := 0; i <= limit; i++ {
		squareOfSums += int(math.Pow(float64(i), float64(2)))
		sumOfSquares += i
	}
	result := int(math.Pow(float64(sumOfSquares), 2)) - squareOfSums
	return result
}
