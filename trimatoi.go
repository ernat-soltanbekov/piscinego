package piscine

func TrimAtoi(s string) int {
	res := 0
	sign := 1
	started := false
	foundDigit := false

	for _, r := range s {
		// 1. Обработка знака (только если еще не начали собирать число)
		if (r == '-' || r == '+') && !started {
			if r == '-' {
				sign = -1
			}
			started = true
			continue
		}

		// 2. Если символ — цифра, добавляем её
		if r >= '0' && r <= '9' {
			res = res*10 + int(r-'0')
			started = true
			foundDigit = true
		}
		// Всё остальное (буквы и прочее) просто пропускаем
	}

	if !foundDigit {
		return 0
	}
	return res * sign
}
