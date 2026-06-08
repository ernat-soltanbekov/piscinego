package piscine

// BTreeIsBinary проверяет, является ли дерево бинарным деревом поиска
func BTreeIsBinary(root *TreeNode) bool {
	return isBSTHelper(root, nil, nil)
}

// isBSTHelper рекурсивно проверяет границы значений для каждого узла
func isBSTHelper(node *TreeNode, min, max *string) bool {
	if node == nil {
		return true
	}

	// Проверяем, не выходит ли текущий узел за рамки допустимых границ
	if (min != nil && node.Data < *min) || (max != nil && node.Data >= *max) {
		return false
	}

	// Рекурсивно проверяем детей, сужая границы
	// Для левого узла верхняя граница (max) — это текущий узел
	// Для правого узла нижняя граница (min) — это текущий узел
	return isBSTHelper(node.Left, min, &node.Data) &&
		isBSTHelper(node.Right, &node.Data, max)
}
