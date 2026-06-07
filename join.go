package piscine

func Join(elems []string, sep string) string {
	// Если массив пустой, возвращаем пустую строку
	if len(elems) == 0 {
		return ""
	}

	// Начинаем с первого элемента
	result := elems[0]

	// Проходим по остальным элементам (начиная со второго)
	for i := 1; i < len(elems); i++ {
		// Сначала добавляем разделитель, потом следующий элемент
		result += sep + elems[i]
	}

	return result
}
