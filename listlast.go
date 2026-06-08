package piscine

func ListLast(l *List) interface{} {
	// 1. Проверяем, не пустой ли список
	if l.Tail == nil {
		return nil
	}

	// 2. Если список не пуст, возвращаем интерфейс Data последнего узла
	return l.Tail.Data
}
