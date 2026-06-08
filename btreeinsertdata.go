package piscine

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	// Если дерево вообще пустое (корень nil), создаем и возвращаем сам корень
	if root == nil {
		return &TreeNode{Data: data}
	}

	// Если новые данные меньше данных текущего узла, идем НАЛЕВО
	if data < root.Data {
		if root.Left == nil {
			// Если слева пусто — это наше место! Создаем узел и связываем с родителем
			root.Left = &TreeNode{Data: data, Parent: root}
		} else {
			// Если место занято, спускаемся глубже по левой ветке (рекурсия)
			BTreeInsertData(root.Left, data)
		}
	} else {
		// Если новые данные больше или равны, идем НАПРАВО
		if root.Right == nil {
			// Если справа пусто — создаем узел здесь
			root.Right = &TreeNode{Data: data, Parent: root}
		} else {
			// Если занято, спускаемся глубже по правой ветке
			BTreeInsertData(root.Right, data)
		}
	}

	// Функция должна вернуть указатель на корень дерева
	return root
}
