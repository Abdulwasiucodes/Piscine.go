package piscine

func ListReverse(l *List) {
	if l == nil || l.Head == nil {
		return
	}
	var prev *NodeL
	curr := l.Head
	l.Tail = l.Head

	for curr != nil {
		nextTemp := curr.Next
		curr.Next = prev
		prev = curr
		curr = nextTemp
	}
	l.Head = prev
}
