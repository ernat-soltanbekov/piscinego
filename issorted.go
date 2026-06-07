package piscine

// IsSorted проверяет, отсортирован ли срез a согласно функции f
func IsSorted(f func(a, b int) int, a []int) bool {
	if len(a) <= 1 {
		return true
	}

	// Определяем направление сортировки по первым двум элементам
	// или просто проверяем монотонность всей последовательности
	ascending := true
	descending := true

	for i := 0; i < len(a)-1; i++ {
		res := f(a[i], a[i+1])
		if res > 0 {
			ascending = false
		} else if res < 0 {
			descending = false
		}
	}

	return ascending || descending
}
