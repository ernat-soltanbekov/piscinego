package piscine

func LoafOfBread(str string) string {
	// 1. Если строка абсолютно пустая
	if str == "" {
		return "\n"
	}

	// 2. Если строка меньше 5 символов
	if len(str) < 5 {
		return "Invalid Output\n"
	}

	result := ""
	count := 0

	// 3. Основная логика
	for i := 0; i < len(str); i++ {
		// Игнорируем пробелы, пока собираем 5 букв
		if str[i] == ' ' && count < 5 {
			continue
		}

		// Если уже собрали 5 букв
		if count == 5 {
			result += " " // добавляем пробел вместо следующего символа
			count = 0     // сбрасываем счетчик
			continue      // пропускаем текущий символ
		}

		result += string(str[i])
		count++
	}

	// 4. НОВОЕ: Отсекаем висящий пробел в самом конце, если он там оказался
	// (срабатывает, если пропущенный символ был самым последним в строке)
	if len(result) > 0 && result[len(result)-1] == ' ' {
		// Берем срез строки от начала и до предпоследнего символа
		result = result[:len(result)-1]
	}

	return result + "\n"
}
