package piscine

// CountIf возвращает количество элементов среза tab, для которых функция f вернула true.
func CountIf(f func(string) bool, tab []string) int {
	count := 0
	for _, s := range tab {
		// Если условие выполняется, увеличиваем счетчик
		if f(s) {
			count++
		}
	}
	return count
}
