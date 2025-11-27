package structures

type Node struct {
	Val  string
	Next *Node
}

type LinkedList struct {
	Head *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (l *LinkedList) PushBack(val string) {
	newNode := &Node{Val: val}
	if l.Head == nil {
		l.Head = newNode
		return
	}
	curr := l.Head
	for curr.Next != nil {
		curr = curr.Next
	}
	curr.Next = newNode
}

func (l *LinkedList) GetHead() string {
	if l.Head == nil {
		return ""
	}
	return l.Head.Val
}
