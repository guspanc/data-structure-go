package list

// Node represents node struct
type Node struct {
	Data interface{}
	next *Node
}

// NewNode node ctor
func NewNode(data interface{}) *Node {
	return &Node{Data: data}
}

// LinkedList represents singly linked list struct
type LinkedList struct {
	head   *Node
	length int
}

// Len returns list length
func (l *LinkedList) Len() int {
	return l.length
}

// Prepend adds a node item
func (l *LinkedList) Prepend(n *Node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

// DeleteWithValue removes from the list by value
func (l *LinkedList) DeleteWithValue(value interface{}) {
	if l.length == 0 {
		return
	}

	if l.head.Data == value {
		l.head = l.head.next
		l.length--
		return
	}

	previousToDelete := l.head
	for previousToDelete.next.Data != value {
		if previousToDelete.next.next == nil {
			return
		}
		previousToDelete = previousToDelete.next
	}
	previousToDelete.next = previousToDelete.next.next
	l.length--
}
