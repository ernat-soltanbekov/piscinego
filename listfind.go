package piscine

func CompStr(a, b interface{}) bool {
	return a == b
}

func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	current := l.Head

	for current != nil {
		// Передаем данные узла и эталон в функцию сравнения
		if comp(current.Data, ref) {
			// Возвращаем адрес поля Data внутри этого узла
			return &current.Data
		}
		current = current.Next
	}

	return nil
}
