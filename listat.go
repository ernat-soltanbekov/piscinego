package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	// 1. Проверяем на отрицательный индекс
	if pos < 0 {
		return nil
	}

	current := l

	// 2. Смещаем указатель вперед pos раз
	for i := 0; i < pos; i++ {
		// Если мы уперлись в nil до того, как закончился цикл,
		// значит, pos больше, чем длина самого списка
		if current == nil {
			return nil
		}
		current = current.Next
	}

	// 3. Возвращаем найденный узел (или nil, если список закончился ровно на этом шаге)
	return current
}
