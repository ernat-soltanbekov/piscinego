package piscine

func Split(s, sep string) []string {
	var res []string
	start := 0

	// Идем по строке до тех пор, пока индекс не выйдет за границы
	for i := 0; i <= len(s)-len(sep); i++ {
		// Проверяем, совпадает ли кусок строки с разделителем
		if s[i:i+len(sep)] == sep {
			res = append(res, s[start:i])
			start = i + len(sep)
			i = start - 1 // Сдвигаем индекс
		}
	}

	// Добавляем остаток строки после последнего разделителя
	res = append(res, s[start:])
	return res
}
