package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {
	// 1. Создаем новый узел с переданными данными
	newNode := &NodeL{Data: data}

	// 2. Проверяем, не пустой ли список
	if l.Head == nil {
		// Если список пуст, новый узел становится одновременно и началом, и концом
		l.Head = newNode
		l.Tail = newNode
		return
	}

	// 3. Если список не пуст, связываем текущий хвост с новым узлом
	l.Tail.Next = newNode

	// 4. Обновляем указатель Tail в структуре списка, делая новый узел последним
	l.Tail = newNode
}
