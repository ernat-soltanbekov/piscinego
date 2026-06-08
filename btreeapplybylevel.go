package piscine

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	// Если дерево пустое, нечего применять
	if root == nil {
		return
	}

	// Создаем очередь и кладем туда корень
	queue := []*TreeNode{root}

	// Пока в очереди есть узлы
	for len(queue) > 0 {
		// Берем первый узел из очереди
		node := queue[0]
		queue = queue[1:]

		// Применяем функцию f к данным узла
		f(node.Data)

		// Если есть левый ребенок, добавляем его в очередь
		if node.Left != nil {
			queue = append(queue, node.Left)
		}

		// Если есть правый ребенок, добавляем его в очередь
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
}
