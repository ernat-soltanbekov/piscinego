package piscine

func RockAndRoll(n int) string {
	// 1. Сначала отсекаем отрицательные числа
	if n < 0 {
		return "error: number is negative\n"
	}

	// 2. Затем проверяем самое строгое условие: деление и на 2, и на 3
	if n%2 == 0 && n%3 == 0 {
		return "rock and roll\n"
	}

	// 3. Проверяем деление только на 2
	if n%2 == 0 {
		return "rock\n"
	}

	// 4. Проверяем деление только на 3
	if n%3 == 0 {
		return "roll\n"
	}

	// 5. Если ничего не подошло, возвращаем ошибку
	return "error: non divisible\n"
}
