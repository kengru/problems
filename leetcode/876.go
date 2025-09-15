package leetcode

// Given the head of a singly linked list, return the middle node of the linked list.

// If there are two middle nodes, return the second middle node.

var Examples876 = Example[*ListNode[int]]{
	Tests: []*ListNode[int]{
		CreateSinglyLinkedList([]int{1, 2, 3, 4, 5}),
		CreateSinglyLinkedList([]int{1, 2, 3, 4, 5, 6}),
	},
}

func MiddleNode876(head *ListNode[int]) *ListNode[int] {
	if head == nil || head.Next == nil {
		return head
	}

	middle := head
	iter := 1

	for head.Next != nil {
		if iter&1 != 0 {
			middle = middle.Next
		}
		head = head.Next
		iter++
	}
	return middle
}

// Old approach
// func MiddleNode876(head *ListNode[int]) *ListNode[int] {
// 	if head == nil || head.Next == nil {
// 		return head
// 	}

// 	length := 1
// 	pointer := head
// 	for pointer.Next != nil {
// 		pointer = pointer.Next
// 		length++
// 	}
// 	for i := 0; i < length/2; i++ {
// 		head = head.Next
// 	}
// 	return head
// }
