package main

import "slices"

// Ищет наибольший простой делитель, проверять на 600851475143
func euler_3(x int) int {
	return slices.Max(findSimpleDividers(x))
}
