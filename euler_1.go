package main

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
