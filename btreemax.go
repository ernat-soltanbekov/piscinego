package piscine

func BTreeMax(root *TreeNode) *TreeNode {
	// Если дерево пустое, возвращаем nil
	if root == nil {
		return nil
	}

	// Идем вправо до самого конца
	current := root
	for current.Right != nil {
		current = current.Right
	}

	// Тот узел, у которого нет правого ребенка, и является максимальным
	return current
}
