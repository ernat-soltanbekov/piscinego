package piscine

// Any возвращает true, если функция f вернула true хотя бы для одного элемента среза.
func Any(f func(string) bool, a []string) bool {
	for _, s := range a {
		// Если нашли хотя бы один подходящий элемент — сразу выходим с true
		if f(s) {
			return true
		}
	}
	// Если цикл закончился и ничего не нашли — возвращаем false
	return false
}
