package example

import (
	"geek-study/algorithm/common"
)

func reverseKGroup(head *common.ListNode, k int) *common.ListNode {
	return common.NewListNode(1, nil)
}

func reverse(start, end *common.ListNode) (head, tail *common.ListNode) {
	if head == end || end == nil || end.Next == nil {
		return start, end
	}
	cur := start
	next := start.Next
	for next.Next != end {
		cur.Next = next.Next
		next.Next = head
		head = next
		next = cur.Next
	}
	tail = cur
	common.PrintListNode(head)
	return head, tail

}

func main() {
	l := common.NewList(1, 2, 3, 4, 5, 6)
	reverse(l.Next, l.Next.Next.Next)
	reverseKGroup(l, 2)
}
