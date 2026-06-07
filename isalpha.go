package piscine

func IsAlpha(s string) bool {
	// Если строка пустая, тест чаще всего ожидает false
	if len(s) == 0 {
		return false
	}

	for _, r := range s {
		// Проверяем: если символ НЕ является буквой И НЕ является цифрой
		if !((r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9')) {
			return false
		}
	}

	return true
}
