package piscine

func BTreeLevelCount(root *TreeNode) int {
	// Базовый случай: если узел пустой (дерево или ветка закончились),
	// его высота равна 0.
	if root == nil {
		return 0
	}

	// Шаг 1: Рекурсивно вычисляем глубину левого поддерева
	leftLevel := BTreeLevelCount(root.Left)

	// Шаг 2: Рекурсивно вычисляем глубину правого поддерева
	rightLevel := BTreeLevelCount(root.Right)

	// Шаг 3: Сравниваем высоты веток и возвращаем наибольшую + 1 (за текущий узел)
	if leftLevel > rightLevel {
		return leftLevel + 1
	}
	return rightLevel + 1
}
