package piscine

func BTreeMin(root *TreeNode) *TreeNode {
	// Если дерево пустое, возвращаем nil
	if root == nil {
		return nil
	}

	// Спускаемся по левым ссылкам до тех пор, пока это возможно
	current := root
	for current.Left != nil {
		current = current.Left
	}

	// Узел, у которого нет левого ребенка, — это минимум
	return current
}
