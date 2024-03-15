package main

import (
	"math"
	"strconv"
	"strings"
)

// Получает строку, делает из неё 'матрицу' размером size, вычисляет максимальное произведение 4х последовательных чисел в матрице.
func euler_11(size int, s string) int {
	var maxsum int
	var slisedString = make([][]int, size)
	splitString := strings.Split(s, " ")
	for i := 0; i < len(splitString) && i < math.Pow(float64(size), 2); i++ {
		j := i % size
		k := i / size
		slisedString[k][j], _ = strconv.Atoi(splitString[i])
	}
	for i := range slisedString {
		for k := range slisedString[i] {
			if len(slisedString)-i < 3 {

			}
		}
	}
	return slisedString
}
