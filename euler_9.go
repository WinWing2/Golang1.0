package main

import (
	"math"
)

// Ищет произведение трёх чисел пифагора, сумма которых равна sum, вплоть до лимита limit
func euler_9(sum int) (multiplication int) {
	for a := 1; a <= sum; a++ {
		for b := 1; b <= sum; b++ {
			for c := 1; c <= sum; c++ {
				if a+b+c == sum && math.Pow(float64(a), 2)+math.Pow(float64(b), 2) == math.Pow(float64(c), 2) {
					return a * b * c
				}
			}
		}
	}
	return 0
}
