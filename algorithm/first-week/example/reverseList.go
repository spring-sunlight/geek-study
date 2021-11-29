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

	common.PrintListNode(head, "初始值")
	fmt.Println("--------------")
	var pre, nextHead *common.ListNode

	for head != nil {
		nextHead = head.Next
		head.Next = pre
		pre = head
		head = nextHead
		common.PrintListNode(pre)
		fmt.Println("--------------")
	}
	return pre
}

func TestReverseList() {
	L := common.NewList(1, 2, 3, 4, 5)
	common.PrintListNode(reverseList(L))
}
