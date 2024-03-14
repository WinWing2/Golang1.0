package main

import (
	"strconv"
	"strings"
)

// Считает самое большое из произведений N чисел, идущих подряд в строке с числами.
func euler_8(n int, s string) (product, position int) {
	biggestProduct := 0
	startingPosition := 0
	slisedString := strings.Split(s, "")
	for i := 0; i <= len(s)-n; i++ {
		currentProduct := 1
		for k := i; k < i+n; k++ {
			currentNumber, _ := strconv.Atoi(slisedString[k])
			currentProduct *= currentNumber
			if currentProduct > biggestProduct {
				biggestProduct = currentProduct
				startingPosition = k
			}
		}
	}
	return biggestProduct, startingPosition
}
