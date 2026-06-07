package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	// 1. Перевод из baseFrom в десятичное число (int)
	n := 0
	for _, char := range nbr {
		val := findIndex(baseFrom, char)
		n = n*len(baseFrom) + val
	}

	// 2. Перевод из десятичного в baseTo
	if n == 0 {
		return string(baseTo[0])
	}

	res := ""
	for n > 0 {
		res = string(baseTo[n%len(baseTo)]) + res
		n /= len(baseTo)
	}
	return res
}

// Вспомогательная функция для поиска индекса символа в базе
func findIndex(base string, char rune) int {
	for i, c := range base {
		if c == char {
			return i
		}
	}
	return 0
}
