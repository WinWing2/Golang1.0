package main

// Считает простое число номера serialNumber, предел поиска searchingLimit
func euler_7(serialNumber, searchingLimit int) (primeNumberValue, primesFound int) {
	primeNumbersFound := 0
	biggestPrimeNumberFound := 0
	for i := 1; primeNumbersFound < serialNumber && i < searchingLimit; i++ {
		if len(findSimpleDividers(i)) == 1 {
			biggestPrimeNumberFound = i
			primeNumbersFound++
		}
	}
	return biggestPrimeNumberFound, primeNumbersFound
}
