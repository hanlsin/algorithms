package data_structures

import (
	"fmt"
	"testing"
)

func TestDoublyLinkedList(t *testing.T) {
	dll := NewYPDllist()
	dll.PushHead("Hello")
	dll.PushHead(123)
	dll.PushTail(true)

	for dll.Head(); dll.Next() != nil; {
		fmt.Println("Value: ", dll.GetValue())
	}
}