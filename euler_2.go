package main

// Суммирует все чётные элементы ряда Фибоначи, до указанного числа, проверять на 4_000_000
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
