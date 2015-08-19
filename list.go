package linkedList

type List struct {
	first, current, last *Node
	len            int
}

func (l *List) Append(value interface{}) {
	if l.current == nil {
		l.init(value)
	} else {
		newNode := &Node{Value: value}
		newNode.prev = l.last
		newNode.next = l.first
		l.last = newNode
		l.current = l.last
	}
	l.len++
}

func (l *List) init(value interface{}) {
	l.current = &Node{Value: value}
	l.first = l.current
	l.last = l.current
}

func (l *List) Len() int {
	return l.len
}

func (l *List) Get() *Node {
	return l.current
}

func (l *List) Next() *Node {
	if (l.current.next != nil) {
		l.current = l.current.next
	}
	return l.current
}

func (l *List) Prev() *Node {
	if (l.current.prev != nil) {
		l.current = l.current.prev
	}
	return l.current
}

func (l *List) Add(value interface{}) *Node {
	if l.current == nil {
		l.init(value)
		l.len++
	} else {
		newNode := &Node{Value: value}
		newNode.next = l.current.next
		l.current.next = newNode
		newNode.prev = l.current		
		l.len++
	}
	return l.current
}

func (l *List) GetFirst() *Node {
	return l.first
}

func (l *List) GetLast() *Node {
	return l.last
}

func (l *List) Remove() bool {
	return l.last == l.current
}