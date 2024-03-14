package main

import "strconv"

// Ищет самое большое число, являющееся палиндромом среди произведения чисел указанной длинны, проверять на 3.
func euler_4(limit int) int {
	var iterationLimit int
	var palindrom int
	var biggestPalindrom = 0
	for i := 0; i < limit; i++ {
		iterationLimit = iterationLimit*10 + 9
	}
	for i := iterationLimit; i > 0; i-- {
		for k := iterationLimit; k > 0; k-- {
			palindrom = i * k
			if palindromCheck(strconv.Itoa(palindrom)) && palindrom > biggestPalindrom {
				biggestPalindrom = palindrom
			}
		}
	}
	return biggestPalindrom
}
