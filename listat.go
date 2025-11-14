package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	if l == nil || pos < 0 {
		return nil
	}

	index := 0
	current := l

	for current != nil {
		if index == pos {
			return current
		}

		current = current.Next
		index++
	}
	return nil
}
