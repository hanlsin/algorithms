package main

import "fmt"

type LLNode struct {
	v int
	prev *LLNode
	next *LLNode
}

func main() {
	num := []int{1, 4, 3, 4, 2, 5}

	head := LLNode{}
	head.next = &LLNode{
		v:    num[0],
		prev: &head,
		next: nil,
	}

	for i := 1; i < len(num); i++ {
		n := num[i]
		node := head.next

		for {
			if node.v >= n {
				if node.next == nil {
					new := &LLNode {
						v:    n,
						prev: node,
					}
					node.next = new
					break
				} else {
					node = node.next
				}
			} else {
				new := &LLNode {
					v:    n,
					prev: node.prev,
					next: node,
				}
				node.prev = new
				if new.prev != nil {
					new.prev.next = new
				}
				break
			}
		}
	}

	node := head.next
	k := 3
	n := 0
	for i := 0; i < k; i++ {
		n = node.v
		node = node.next
		if node == nil {
			break
		}
	}
	fmt.Println(n)
}
