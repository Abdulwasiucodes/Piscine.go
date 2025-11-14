package piscine

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	newNode := &NodeI{Data: data_ref}
	if l == nil || data_ref < l.Data {
		newNode.Next = l
		return newNode
	}
	iterator := l
	for iterator.Next != nil && iterator.Next.Data < data_ref {
		iterator = iterator.Next
	}
	newNode.Next = iterator.Next
	iterator.Next = newNode
	return l
}
