package main

import (
	"slices"
	"strings"
)

// Дробит число на простые множители.
func findSimpleDividers(x int) []int {
	var dividers []int
	for i := 1; i <= x && x != 1; i++ {
		if x%i == 0 {
			dividers = append(dividers, i)
			x = x / i
			i = 1
		}
	}
	return dividers
}

func palindromCheck(s string) bool {
	slisedString := make([]string, len(s))
	slisedString = strings.Split(s, "")
	reversedSlisedString := make([]string, len(s))
	copy(reversedSlisedString, slisedString)
	slices.Reverse(reversedSlisedString)
	return slices.Equal(slisedString, reversedSlisedString)
}
