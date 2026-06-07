package piscine

func IsNumeric(s string) bool {
	// Если строка пустая, возвращаем false (как и в предыдущих задачах)
	if len(s) == 0 {
		return false
	}

	for _, r := range s {
		// Если символ НЕ является цифрой (вне диапазона от '0' до '9')
		if !(r >= '0' && r <= '9') {
			return false
		}
	}

	return true
}
