package piscine

func StringToIntSlice(str string) []int {
	// Если строка пустая, вернется пустой срез (nil), что корректно обрабатывается при выводе
	var result []int

	// Цикл range автоматически разбивает строку на руны (символы)
	for _, char := range str {
		// char имеет тип rune (эквивалент int32),
		// поэтому мы явно преобразуем его в int, как того требует функция
		result = append(result, int(char))
	}

	return result
}
