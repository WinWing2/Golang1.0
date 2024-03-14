package main

import (
	"fmt"
	"math"
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

// Суммирует все числа, кратные 3 и 5, до указанного числа, проверять на 1000
func euler_1(limit int) int {
	sum := 0
	for i := 1; i < limit; i++ {
		if i%3 == 0 {
			sum += i
		}
		if i%5 == 0 {
			sum += i
		}
	}
	return sum
}

// Суммирует все чётные элементы ряда Фибоначи, до указанного числа, проверять на 4000000
func euler_2(limit int) int {
	sum := 0
	f1, f2 := 0, 1
	for f1 < limit {
		if f1%2 == 0 {
			sum += f1
			//fmt.Println("Чётные числа фибоначи: = ", f1)
		}
		f1, f2 = f2, f1+f2
	}
	return sum
}

// Ищет наибольший простой делитель, проверять на 600851475143
func euler_3(x int) int {
	return slices.Max(findSimpleDividers(x))
}

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
			if palindromCheck(fmt.Sprintf("%v", palindrom)) && palindrom > biggestPalindrom {
				biggestPalindrom = palindrom
			}
		}
	}
	return biggestPalindrom
}

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

// Считает разницу меж квадратом суммы и суммой квадратов, проверять на 10, = 2640
func euler_6(limit int) int {
	squareOfSums := 0
	sumOfSquares := 0
	for i := 0; i <= limit; i++ {
		squareOfSums += int(math.Pow(float64(i), float64(2)))
		sumOfSquares += i
	}
	result := int(math.Pow(float64(sumOfSquares), 2)) - squareOfSums
	return result
}

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

// Считает самое большое из произведений N чисел, идущих подряд в строке с числами.
func euler_8(n int, s string) int {
	return 1
}

func main() {
	/* Позже изучить
	(fmt.Println() \r \a
	big.int для задачи 8
	*/

	//fmt.Printf("\n\t Решение задачек: ")
	//fmt.Printf("\n\t euler_1 | sum = %v", euler_1(1000))
	//fmt.Printf("\n\t euler_2 | sum = %v", euler_2(4_000_000))
	//fmt.Printf("\n\t euler_3 | biggestSimpleDivider = %v", euler_3(600851475143))
	//fmt.Printf("\n\t euler_4 | biggestPalindrome = %v", euler_4(3))
	//fmt.Printf("\n\t euler_5 | minimalDividedNumber = %v", euler_5(10))
	//fmt.Printf("\n\t euler_6 | sumOfSquares-squareOfSums = %v", euler_6(10))
	//euler_7 := strings.Split(fmt.Sprint(euler_7(1001, 1_000_000)), " ")
	//fmt.Printf("\n\t euler_7 | primeNumberValue=%v, primeSerialNumber = %v", euler_7[0], euler_7[1])
	euler_8(13, "7316717653133062491922511967442657474235534919493496983520312774506326239578318016984801869478851843858615607891129494954595017379583319528532088055111254069874715852386305071569329096329522744304355766896648950445244523161731856403098711121722383113622298934233803081353362766142828064444866452387493035890729629049156044077239071381051585930796086670172427121883998797908792274921901699720888093776657273330010533678812202354218097512545405947522435258490771167055601360483958644670632441572215539753697817977846174064955149290862569321978468622482839722413756570560574902614079729686524145351004748216637048440319989000889524345065854122758866688116427171479924442928230863465674813919123162824586178664583591245665294765456828489128831426076900422421902267105562632111110937054421750694165896040807198403850962455444362981230987879927244284909188845801561660979191338754992005240636899125607176060588611646710940507754100225698315520005593572972571636269561882670428252483600823257530420752963450") //
	fmt.Println("")
}
