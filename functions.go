package main

import (
	"math"
	"slices"
	"strings"
)

// Дробит число на простые множители.
func findSimpleDividers(x int) []int {
	var dividers []int
	for i := 1; i < x && x != 1; i++ {
		if x%i == 0 {
			dividers = append(dividers, i)
			x = x / i
			i = 1
		}
	}
	return dividers
}

// Ищет все простые числа от 1 до N (включительно)
func findPrimeNumbers(limit int) []int {
	limit++
	numbersBeforeCheck := make([]bool, limit)
	var numbersAfterCheck []int
	checklimit := int(math.Sqrt(float64(limit))) + 1
	for i := 2; i < checklimit; i++ {
		if !numbersBeforeCheck[i] {
			for k := 2 * i; k < limit; k += i {
				numbersBeforeCheck[k] = true
			}
		}
	}
	for i := range numbersBeforeCheck {
		if !numbersBeforeCheck[i] {
			numbersAfterCheck = append(numbersAfterCheck, i)
		}
	}
	numbersAfterCheck = numbersAfterCheck[2:]
	return numbersAfterCheck
}

func palindromCheck(s string) bool {
	slisedString := make([]string, len(s))
	slisedString = strings.Split(s, "")
	reversedSlisedString := make([]string, len(s))
	copy(reversedSlisedString, slisedString)
	slices.Reverse(reversedSlisedString)
	return slices.Equal(slisedString, reversedSlisedString)
}

// Ищет все числа формата a^2 + b^2 = c^2, вплоть до лимита N
// Способ вывести значения на экран
// pifagors := pifagorNumbersGenerator(100)
//
//	for i := 0; i < len(pifagors); i++ {
//		fmt.Printf("Numbers form i=%v:", i)
//		for k := 0; k < len(pifagors[i]); k++ {
//			fmt.Printf("\t%v", pifagors[i][k])
//		}
//		fmt.Printf("\n")
//	}
func pifagorNumbersGenerator(limit int) (numbers [][]int) {
	var pifagors [][]int
	for a := 1; a <= limit; a++ {
		for b := 1; b <= limit; b++ {
			for c := 1; c <= limit; c++ {
				if math.Pow(float64(a), 2)+math.Pow(float64(b), 2) == math.Pow(float64(c), 2) {
					pifagors = append(pifagors, []int{a, b, c})
				}
			}
		}
	}
	ilim := len(pifagors)
	for i := 0; i < ilim; i++ {
		slices.Sort(pifagors[i])
		if i != 0 && slices.Equal(pifagors[i], pifagors[i-1]) {
			pifagors = slices.Delete(pifagors, i-1, i)
			ilim--
			i--
		}
	}
	return pifagors
}
