package data_structures

/*
https://en.wikipedia.org/wiki/Linked_data_structure
 */

type LLNode struct {
	value *Node
	prev *LLNode
	next *LLNode
}

type LinkedList interface {
	PushHead(interface{}) *LLNode
	PushTail(interface{}) *LLNode
	/*
	PopHead() interface{}
	PopTail() interface{}
	*/
	Head() LinkedList
	Tail() LinkedList
	Next() LinkedList
	Prev() LinkedList
	GetValue() interface{}
	/*
	PutNext(*LLNode, interface{})
	PutPrev(interface{})
	*/
	Len() int
}

/*
Doubly linked list
https://en.wikipedia.org/wiki/Doubly_linked_list
 */
type YPDllist struct {
	size int
	head *LLNode
	tail *LLNode
	current *LLNode
}

func NewYPDllist() *YPDllist {
	head := &LLNode {
		value: nil,
		prev: nil,
		next: nil,
	}
	tail := &LLNode {
		value: nil,
		prev: nil,
		next: nil,
	}

	head.next = tail
	head.prev = tail
	tail.prev = head
	tail.next = head

	return &YPDllist{
		size: 0,
		head: head,
		tail: tail,
	}
}

func (dll *YPDllist) PushHead(obj interface{}) *LLNode {
	nextNode := dll.head.next
	dll.head.next = &LLNode{
		value: &Node{
			object: obj,
		},
		prev:  dll.head,
		next:  nextNode,
	}
	nextNode.prev = dll.head.next
	dll.size++
	dll.current = dll.head.next
	return dll.current
}

func (dll *YPDllist) PushTail(obj interface{}) *LLNode {
	prevNode := dll.tail.prev
	dll.tail.prev = &LLNode{
		value: &Node{
			object: obj,
		},
		prev: prevNode,
		next: dll.tail,
	}
	prevNode.next = dll.tail.prev
	dll.size++
	dll.current = dll.tail.prev
	return dll.current
}

func (dll *YPDllist) Head() LinkedList {
	dll.current = dll.head
	return dll
}

func (dll *YPDllist) Tail() LinkedList {
	dll.current = dll.tail
	return dll
}

func (dll *YPDllist) Next() LinkedList {
	dll.current = dll.current.next
	if dll.current == dll.tail {
		return nil
	}
	return dll
}

func (dll *YPDllist) Prev() LinkedList {
	dll.current = dll.current.prev
	if dll.current == dll.head {
		return nil
	}
	return dll
}

func (dll *YPDllist) GetValue() interface{} {
	return dll.current.value
}

func (dll *YPDllist) Len() int {
	return dll.size
}
