package piscine

func ReverseMenuIndex(menu []string) []string {
	// 1. Узнаем длину исходного среза
	length := len(menu)

	// 2. Создаем новый срез такого же размера с помощью make
	result := make([]string, length)

	// 3. Проходим циклом по исходному меню
	for i := 0; i < length; i++ {
		// 4. Записываем элементы в новый срез с конца в начало
		// Индекс конца: length - 1 - i
		result[length-1-i] = menu[i]
	}

	// Возвращаем перевернутый срез
	return result
}
