package piscine

func AlphaCount(s string) int {
	count := 0
	// Превращаем строку в массив рун, чтобы корректно читать каждый символ
	for _, r := range s {
		// Проверяем, является ли символ заглавной или строчной латинской буквой
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
			count++
		}
	}
	return count
}
