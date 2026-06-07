package piscine

func Capitalize(s string) string {
	runes := []rune(s)
	isNewWord := true // Флаг: следующий символ должен быть заглавным

	for i, r := range runes {
		if isAlphaNumeric(r) {
			if isNewWord {
				// Если это начало слова — делаем заглавной
				if r >= 'a' && r <= 'z' {
					runes[i] = r - 32
				}
				isNewWord = false
			} else {
				// Если это середина слова — делаем строчной
				if r >= 'A' && r <= 'Z' {
					runes[i] = r + 32
				}
			}
		} else {
			// Если это не буква/цифра, следующий символ будет началом нового слова
			isNewWord = true
		}
	}
	return string(runes)
}

// Вспомогательная функция для проверки, является ли символ буквой или цифрой
func isAlphaNumeric(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9')
}
