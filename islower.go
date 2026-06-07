package piscine

func IsLower(s string) bool {
	// Если строка пустая, тест, скорее всего, ожидает false
	if len(s) == 0 {
		return false
	}

	for _, r := range s {
		// Если символ - НЕ маленькая латинская буква,
		// значит условие "все строчные" нарушено.
		if !(r >= 'a' && r <= 'z') {
			return false
		}
	}

	return true
}
