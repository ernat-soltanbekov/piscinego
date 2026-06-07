package piscine

func BasicJoin(elems []string) string {
	result := ""

	// Проходим по каждому элементу массива
	for _, str := range elems {
		// Присоединяем текущую строку к общему результату
		result += str
	}

	return result
}
