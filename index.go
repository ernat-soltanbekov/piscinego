package piscine

func Index(s string, toFind string) int {
	// Превращаем обе строки в слайсы рун для корректной работы с символами
	runesS := []rune(s)
	runesFind := []rune(toFind)

	lenS := len(runesS)
	lenFind := len(runesFind)

	// Если искомая строка пустая, вернем 0 (по логике большинства систем)
	if lenFind == 0 {
		return 0
	}

	// Идем по строке s до той точки, где еще может уместиться toFind
	for i := 0; i <= lenS-lenFind; i++ {
		// Проверяем, совпадает ли кусок строки S с toFind
		match := true
		for j := 0; j < lenFind; j++ {
			if runesS[i+j] != runesFind[j] {
				match = false
				break
			}
		}
		// Если кусок совпал, возвращаем индекс начала
		if match {
			return i
		}
	}

	// Если цикл завершился и ничего не нашлось
	return -1
}
