package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func ListSort(l *NodeI) *NodeI {
	if l == nil || l.Next == nil {
		return l
	}

	for i := l; i != nil; i = i.Next {
		for j := i.Next; j != nil; j = j.Next {
			if i.Data > j.Data {
				i.Data, j.Data = j.Data, i.Data
			}
		}
	}

	return l
}
