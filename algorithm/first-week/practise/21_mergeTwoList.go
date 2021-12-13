package practise

import "geek-study/algorithm/common"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *common.ListNode, l2 *common.ListNode) *common.ListNode {
	l := &common.ListNode{Val: 0, Next: nil}
	res := l
	for l1 != nil || l2 != nil {
		if l1 == nil || (l2 != nil && l1.Val >= l2.Val) {
			l.Next = l2
			l2 = l2.Next
		} else if l2 == nil || (l1 != nil && l2.Val > l1.Val) {
			l.Next = l1
			l1 = l1.Next
		}
		l = l.Next
	}
	return res.Next
}

func testMergeTwoLists() {
	l1 := common.NewList(1, 3, 5, 7, 9)
	l2 := common.NewList(2, 4, 6, 8, 10)
	mergeTwoLists(l1, l2)
}
