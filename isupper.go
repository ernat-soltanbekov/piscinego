package piscine

func IsUpper(s string) bool {
	// Если строка пустая, тест скорее всего ждет false (исходя из твоей ошибки)
	if len(s) == 0 {
		return false
	}

	for _, r := range s {
		// Если символ - строчная латинская буква, возвращаем false
		if r >= 'a' && r <= 'z' {
			return false
		}
		// Если символ - НЕ заглавная буква, возвращаем false
		// (То есть, если это цифра, спецсимвол или что-то другое)
		if !(r >= 'A' && r <= 'Z') {
			return false
		}
	}

	return true
}
