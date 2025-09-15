package leetcode

import "fmt"

type ListNode[V any] struct {
	Val  V
	Next *ListNode[V]
}

func CreateSinglyLinkedList[V any](values []V) *ListNode[V] {
	head := &ListNode[V]{
		Val:  values[0],
		Next: nil,
	}
	pointer := head
	for i := 1; i < len(values); i++ {
		ln := &ListNode[V]{
			Val:  values[i],
			Next: nil,
		}
		pointer.Next = ln
		pointer = ln
	}
	return head
}

func PrintSinglyLinkedList[V any](head *ListNode[V]) {
	fmt.Printf("[%v,", head.Val)
	for head.Next != nil {
		head = head.Next
		if head != nil {
			fmt.Printf("%v,", head.Val)
		}
	}
	fmt.Printf("]\n")
}
