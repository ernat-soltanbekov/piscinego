package piscine

func LastRune(s string) rune {
	// 1. Превращаем строку в массив рун
	runes := []rune(s)

	// 2. Находим индекс последнего элемента (длина минус 1)
	lastIndex := len(runes) - 1

	// 3. Возвращаем элемент по этому индексу
	return runes[lastIndex]
}
