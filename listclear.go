package piscine

func ListClear(l *List) {
	if l != nil {
		l.Head = nil
		l.Tail = nil
	}
}
