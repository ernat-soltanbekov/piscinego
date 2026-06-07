package piscine

func Abort(a, b, c, d, e int) int {
	// 1. Собираем все аргументы в один срез (slice)
	nums := []int{a, b, c, d, e}

	// 2. Сортируем срез методом "пузырька" (Bubble Sort)
	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				// Меняем элементы местами, если предыдущий больше следующего
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}

	// 3. Возвращаем центральный элемент
	// Так как элементов всего 5, индексы: 0, 1, 2, 3, 4. Центр — это индекс 2.
	return nums[2]
}
