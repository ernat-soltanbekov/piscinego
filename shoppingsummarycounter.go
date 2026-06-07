package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	summary := make(map[string]int)
	word := ""

	for _, char := range str {
		if char == ' ' {
			// Мы убрали проверку if word != ""
			// Теперь, если встретился пробел, мы слепо добавляем текущее содержимое word
			// в мапу. Даже если это пустая строка "".
			summary[word]++
			word = "" // Сбрасываем для следующего слова
		} else {
			word += string(char)
		}
	}

	// Обязательно добавляем последний кусок (даже если он пустой)
	summary[word]++

	return summary
}
