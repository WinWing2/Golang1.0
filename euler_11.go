package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Получает строку, делает из неё 'матрицу' размером size, вычисляет максимальное произведение 4х последовательных чисел в матрице.
func euler_11(squareSize, lineLength int, s string) int {

	var slisedString = make([][]int, squareSize)
	splitString := strings.Split(s, " ")

	if int(math.Pow(float64(squareSize), 2)) != len(splitString) {
		fmt.Printf("\nInput string is not for a square with sides == %v, Len of string = %v\n\n", squareSize, len(splitString))
	}
	for i := 0; i < len(splitString) && i < int(math.Pow(float64(squareSize), 2)); i++ {
		vertical := i / squareSize
		newElement, _ := strconv.Atoi(splitString[i])
		slisedString[vertical] = append(slisedString[vertical], newElement)
	}

	maxsumHorizontal, _, _ := horizontalMaxSum(lineLength, slisedString)
	maxsumVertical, _, _ := verticalMaxSum(lineLength, slisedString)
	maxsumtopDownDiagonall, _, _ := topDownDiagonalMaxSum(lineLength, slisedString)
	maxsumbotomUpDiagonall, _, _ := botomUpDiagonallMaxSum(lineLength, slisedString)
	return max(maxsumHorizontal, maxsumVertical, maxsumtopDownDiagonall, maxsumbotomUpDiagonall)

	/* Если хочется подробный вывод сумм, а не одного числа, как в задаче.
	maxsumHorizontal, yHorizontal, xHorizontal := horizontalMaxSum(lineLength, slisedString)
	maxsumVertical, yVertical, xVertical := verticalMaxSum(lineLength, slisedString)
	maxsumtopDownDiagonall, ytopDownDiagonall, xtopDownDiagonall := topDownDiagonalMaxSum(lineLength, slisedString)
	maxsumbotomUpDiagonall, ybotomUpDiagonall, xbotomUpDiagonall := botomUpDiagonallMaxSum(lineLength, slisedString)

	message := fmt.Sprintf("MaxsumHorizontal = %v, y=%v, x=%v\n", maxsumHorizontal, yHorizontal, xHorizontal)
	message += fmt.Sprintf("MaxsumVertical = %v, y=%v, x=%v\n", maxsumVertical, yVertical, xVertical)
	message += fmt.Sprintf("MaxsumtopDownDiagonall = %v, y=%v, x=%v\n", maxsumtopDownDiagonall, ytopDownDiagonall, xtopDownDiagonall)
	message += fmt.Sprintf("MaxsumbotomUpDiagonall = %v, y=%v, x=%v\n", maxsumbotomUpDiagonall, ybotomUpDiagonall, xbotomUpDiagonall)
	return message
	*/
}

func horizontalMaxSum(lineLength int, square [][]int) (maxSum, y, x int) {
	var maxsum, verticalMax, horizontalMax int
	for vertical := 0; vertical < len(square); vertical++ {
		for horizontal := 0; horizontal <= len(square[vertical])-lineLength; horizontal++ {
			sum := 0
			for i := 0; i < lineLength; i++ {
				sum += square[vertical][horizontal+i]
			}
			if maxsum < sum {
				maxsum = sum
				verticalMax = vertical
				horizontalMax = horizontal
			}
		}
	}
	return maxsum, verticalMax + 1, horizontalMax + 1
}

func verticalMaxSum(lineLength int, square [][]int) (maxSum, y, x int) {
	var maxsum, verticalMax, horizontalMax int
	for vertical := 0; vertical <= len(square)-lineLength; vertical++ {
		for horizontal := 0; horizontal < len(square[vertical]); horizontal++ {
			sum := 0
			for i := 0; i < lineLength; i++ {
				sum += square[vertical+i][horizontal]
			}
			if maxsum < sum {
				maxsum = sum
				verticalMax = vertical
				horizontalMax = horizontal
			}
		}
	}
	return maxsum, verticalMax + 1, horizontalMax + 1
}

func topDownDiagonalMaxSum(lineLength int, square [][]int) (maxSum, y, x int) {
	var maxsum, verticalMax, horizontalMax int
	for vertical := 0; vertical <= len(square)-lineLength; vertical++ {
		for horizontal := 0; horizontal <= len(square[vertical])-lineLength; horizontal++ {
			sum := 0
			for i := 0; i < lineLength; i++ {
				sum += square[vertical+i][horizontal+i]
			}
			if maxsum < sum {
				maxsum = sum
				verticalMax = vertical
				horizontalMax = horizontal
			}
		}
	}
	return maxsum, verticalMax + 1, horizontalMax + 1
}

func botomUpDiagonallMaxSum(lineLength int, square [][]int) (maxSum, y, x int) {
	var maxsum, verticalMax, horizontalMax int
	for vertical := lineLength - 1; vertical < len(square); vertical++ {
		for horizontal := 0; horizontal <= len(square[vertical])-lineLength; horizontal++ {
			sum := 0
			for i := 0; i < lineLength; i++ {
				sum += square[vertical-i][horizontal+i]
			}
			if maxsum < sum {
				maxsum = sum
				verticalMax = vertical
				horizontalMax = horizontal
			}
		}
	}
	return maxsum, verticalMax + 1, horizontalMax + 1
}
