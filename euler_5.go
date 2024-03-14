package main

import "math"

// Находит минимальное число, делящееся на все числа от 1 до N, проверять на 10
func euler_5(x int) int {
	product := 1
	currentNumberDivider := make(map[int]int)
	totalDividers := make(map[int]int)
	for i := 0; i <= x; i++ {
		dividersSlice := findSimpleDividers(i)
		for _, v := range dividersSlice {
			currentNumberDivider[v] += 1
		}
		for i := range currentNumberDivider {
			if currentNumberDivider[i] > totalDividers[i] {
				totalDividers[i] = currentNumberDivider[i]
				//fmt.Println("Содержимое totalDividers = ", totalDividers)
			}
		}
		for i := range currentNumberDivider {
			currentNumberDivider[i] = 0
		}
	}
	for i, v := range totalDividers {
		product *= int(math.Pow(float64(i), float64(v)))
	}
	return product
}
