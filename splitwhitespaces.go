package piscine

func SplitWhiteSpaces(s string) []string {
	var res []string
	var word string

	for _, char := range s {
		// Проверяем, является ли символ разделителем
		if char == ' ' || char == '\t' || char == '\n' {
			if word != "" {
				res = append(res, word)
				word = "" // Сбрасываем слово после добавления
			}
		} else {
			// Накапливаем буквы в текущем слове
			word += string(char)
		}
	}

	// Не забываем добавить последнее слово, если оно было
	if word != "" {
		res = append(res, word)
	}

	return res
}
