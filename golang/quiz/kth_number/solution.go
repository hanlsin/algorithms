package main

import (
	"fmt"
	"time"
)

type LLNode struct {
	v int
	prev *LLNode
	next *LLNode
}

func findKthNum1(numbers []int, k int) int {
	head := LLNode{}
	head.next = &LLNode{
		v:    numbers[0],
		prev: &head,
		next: nil,
	}

	for i := 1; i < len(numbers); i++ {
		n := numbers[i]
		node := head.next

		for {
			if node.v >= n {
				if node.next == nil {
					new := &LLNode {
						v: n,
						prev: node,
					}
					node.next = new
					break
				} else {
					node = node.next
				}
			} else {
				new := &LLNode{
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
	n := 0
	for i := 0; i < k; i++ {
		n = node.v
		node = node.next
		if node == nil {
			break
		}
	}
	return n
}

func findKthNum2(numbers []int, k int) int {
	result := make([]int, k)

	for _, n := range numbers {
		for j := 0; j < k; j++ {
			if n > result[j] {
				tmp := n
				n = result[j]
				result[j] = tmp
			}
		}
	}
	return result[k-1]
}

type findKthNum func([]int, int) int

func calculateElapsed(fn findKthNum, numbers []int, k int) (int, int64) {
	start := time.Now()
	ret := fn(numbers, k)
	elapsed := time.Since(start)
	return ret, elapsed.Nanoseconds()
}

func main() {
	numbers := []int{1, 4, 3, 4, 2, 5}
	k := 3
	fmt.Println("Numbers are ", numbers)
	fmt.Printf("What is the %d (nd/rd/th) largest number?\n", k)

	fmt.Println("Solution 1: ", findKthNum1(numbers, k))
	fmt.Println("Solution 2: ", findKthNum2(numbers, k))

	fmt.Println("Elapsed time of Solution 1: ", fmt.Sprint(calculateElapsed(findKthNum1, numbers, k)))
	fmt.Println("Elapsed time of Solution 2: ", fmt.Sprint(calculateElapsed(findKthNum2, numbers, k)))
}
