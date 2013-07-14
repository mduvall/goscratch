package structures

import (
	"testing"
)

func Test_CreateList(t *testing.T) {
	list := CreateList(1)
	if list.head.element != 1 {
		t.Errorf("The list dun goofed")
	}
}

func Test_FindInList(t *testing.T) {
	list := CreateList(1)
	list.Append(&ListNode{element: 2})

	if list.Find(2) == nil {
		t.Errorf("Couldn't find the element in the list")
	}
}

func Test_InsertInList(t *testing.T) {
	list := CreateList(1)
	list.Append(&ListNode{element: 3})
	list.Insert(&ListNode{element: 2}, 1)

	if list.Find(2).next.element != 3 {
		t.Errorf("Does insert even work?")
	}
}
