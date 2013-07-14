package structures

type ListNode struct {
	next    *ListNode
	element int
}

type List struct {
	head *ListNode
}

func CreateList(i int) List {
	return List{&ListNode{element: i}}
}

func (l *List) Append(n *ListNode) {
	head := l.head

	for head.next != nil {
		head = head.next
	}

	head.next = n
}

func (l *List) Find(i int) *ListNode {
	head := l.head

	for head != nil {
		if head.element == i {
			return head
		}
		head = head.next
	}
	return nil
}

func (l *List) Insert(n *ListNode, idx int) {
	curIdx := 0
	head := l.head

	// Special case for head
	if idx == 0 {
		n.next = head
		l.head = n
	}

	for idx != curIdx+1 || head.next == nil {
		head = head.next
	}

	nextNode := head.next
	head.next = n
	n.next = nextNode
}
