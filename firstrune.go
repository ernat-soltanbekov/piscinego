package piscine

func FirstRune(s string) rune {
	// Превращаем строку в массив рун и берем нулевой элемент
	runes := []rune(s)
	return runes[0]
}
