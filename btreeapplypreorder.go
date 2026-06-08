package piscine

func BTreeApplyPreorder(root *TreeNode, f func(...interface{}) (int, error)) {
	// Базовый случай: если узел пустой, прекращаем выполнение ветки
	if root == nil {
		return
	}

	// Шаг 1: Сразу обрабатываем текущий узел
	f(root.Data)

	// Шаг 2: Рекурсивно уходим в левую ветку
	BTreeApplyPreorder(root.Left, f)

	// Шаг 3: Рекурсивно уходим в правую ветку
	BTreeApplyPreorder(root.Right, f)
}
