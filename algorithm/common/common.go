package common

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewList(val ...int) *ListNode {
	l := make([]*ListNode, len(val))
	for i := 0; i < len(val); i++ {
		l[i] = NewListNode(val[i], nil)
	}
	for i := 0; i < len(l)-1; i++ {
		l[i].Next = l[i+1]
	}
	return l[0]
}

func NewListNode(val int, next *ListNode) *ListNode {
	return &ListNode{Val: val, Next: next}
}

func PrintListNode(node *ListNode) {
	counter := 0
	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
		counter++
		if counter > 20 {
			return
		}
	}
}
