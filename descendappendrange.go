package piscine

func DescendAppendRange(max, min int) []int {
	// 1. Проверка на некорректный или пустой диапазон
	if max <= min {
		// Возвращаем инициализированный пустой срез
		return []int{}
	}

	// 2. Объявляем пустой срез (nil-срез).
	// Память под него пока не выделена.
	var result []int

	// 3. Цикл идет от max до min (не включая min, так как стоит строгое условие i > min)
	for i := max; i > min; i-- {
		// 4. Добавляем текущее значение i в конец среза
		result = append(result, i)
	}

	// 5. Возвращаем заполненный срез
	return result
}
