package piscine

func BTreeApplyInorder(root *TreeNode, f func(...interface{}) (int, error)) {
	// Базовый случай: если пришли в пустой узел (конец ветки), просто выходим
	if root == nil {
		return
	}

	// Шаг 1: Идем в самую глубь левого поддерева
	BTreeApplyInorder(root.Left, f)

	// Шаг 2: Обрабатываем текущий узел (применяем функцию f)
	f(root.Data)

	// Шаг 3: Идем вглубь правого поддерева
	BTreeApplyInorder(root.Right, f)
}
