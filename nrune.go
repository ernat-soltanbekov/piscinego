package piscine

func NRune(s string, n int) rune {
	runes := []rune(s)

	// Проверка: индекс 'n' должен быть больше 0 (так как мы считаем с 1 по заданию)
	// И 'n' должен быть меньше или равен количеству символов в строке
	if n > 0 && n <= len(runes) {
		return runes[n-1]
	}

	// Если n не входит в диапазон, возвращаем 0
	return 0
}
