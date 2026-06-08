package piscine

func BTreeApplyPostorder(root *TreeNode, f func(...interface{}) (int, error)) {
	// Базовый случай: если пришли в пустой узел, выходим из рекурсии
	if root == nil {
		return
	}

	// Шаг 1: Сначала полностью обходим левую ветку
	BTreeApplyPostorder(root.Left, f)

	// Шаг 2: Затем полностью обходим правую ветку
	BTreeApplyPostorder(root.Right, f)

	// Шаг 3: Только после этого обрабатываем текущий узел
	f(root.Data)
}
