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
	var l = common.NewListNode(0, nil)
	for l1 != nil || l2 != nil {
		if l1 == nil || l1.Val > l2.Val {
			l.Next = l2
			l2 = l2.Next
		}
		if l2 == nil || l1.val < l2.val {
			l.Next == l1
			l1 = l1.Next
		}
	}
	common.PrintListNode(l)
	return l
}

func testMergeTwoLists() {
	l1 := common.NewList(1, 3, 5, 7, 9)
	l2 := common.NewList(2, 4, 6, 8, 10)
	mergeTwoLists(l1, l2)
}
