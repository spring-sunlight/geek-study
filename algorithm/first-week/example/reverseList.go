package example

import (
	"fmt"
	"geek-study/algorithm/common"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

//  head
//  cur  next
//   ↓    ↓
//1: 1 -> 2 -> 3 -> 4 -> 5

//   cur.next = next.next
//2: 1    2 -> 3 -> 4 -> 5
//	 |_____↑

//   next.next = head
//3: 1 <-2  3 -> 4 -> 5
//   |_____↑

//  head = next
//  next = cur.next
//  head cur  next
//   ↓    ↓    ↓
//4: 2 -> 1 -> 3 -> 4 -> 5

func reverseList(head *common.ListNode) *common.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	cur := head
	next := head.Next
	common.PrintListNode(head)
	fmt.Println("--------------")
	for next != nil {
		cur.Next = next.Next
		next.Next = head
		head = next
		next = cur.Next
		common.PrintListNode(head)
		fmt.Println("--------------")
	}
	return head
}

func testReverseList() {
	L := common.NewList(1, 2, 3, 4, 5)
	common.PrintListNode(reverseList(L))
}
