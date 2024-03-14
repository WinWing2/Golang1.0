package main

// Считает простое число номера numberOfPrimeNumber, предел поиска searchingLimit
func euler_7(numberOfPrimeNumber, searchingLimit int) (numberValue, primeSerialNumber int) {
	currentNumberOfPrimeNumber := 0
	primeNumberValue := 0
	for i := 1; currentNumberOfPrimeNumber < numberOfPrimeNumber && i < searchingLimit; i++ {
		if len(findSimpleDividers(i)) == 2 && i > primeNumberValue {
			primeNumberValue = i
			currentNumberOfPrimeNumber++
		}
	}
	return primeNumberValue, currentNumberOfPrimeNumber
}
