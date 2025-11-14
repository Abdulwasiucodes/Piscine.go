package piscine

func ListSize(l *List) int {
	count := 0
	currentNode := l.Head
	for currentNode != nil {
		count++
		currentNode = currentNode.Next
	}
	return count
}
