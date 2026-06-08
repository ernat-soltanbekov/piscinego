package piscine

func ListClear(l *List) {
	// Присваиваем указателям на начало и конец списка значение nil
	l.Head = nil
	l.Tail = nil
}
