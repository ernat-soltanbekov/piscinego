package piscine

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	// Базовый случай 1: если дошли до пустого узла, значит элемента в дереве нет
	if root == nil {
		return nil
	}

	// Базовый случай 2: если нашли совпадение, возвращаем текущий узел
	if elem == root.Data {
		return root
	}

	// Если искомое меньше текущего, по правилам BST идем искать в левое поддерево
	if elem < root.Data {
		return BTreeSearchItem(root.Left, elem)
	}

	// В противном случае (если искомое больше), идем искать в правое поддерево
	return BTreeSearchItem(root.Right, elem)
}
