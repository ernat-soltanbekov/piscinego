package piscine

// Map применяет функцию f к каждому элементу среза a и возвращает новый срез bool.
func Map(f func(int) bool, a []int) []bool {
	// Создаем новый срез той же длины, что и исходный
	result := make([]bool, len(a))

	// Проходим циклом по элементам исходного среза
	for i, val := range a {
		// Применяем функцию f к текущему значению и сохраняем результат
		result[i] = f(val)
	}

	return result
}
