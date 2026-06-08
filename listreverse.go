package piscine

func ListReverse(l *List) {
	if l == nil || l.Head == nil || l.Head == l.Tail {
		return
	}

	var prev *NodeL
	current := l.Head
	l.Tail = l.Head

	for current != nil {
		nextTemp := current.Next
		current.Next = prev
		prev = current
		current = nextTemp
	}

	l.Head = prev
}
