package practise

import "geek-study/algorithm/common"

func mergeKLists(lists []*common.ListNode) *common.ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	if len(lists) == 2 {
		return mergeList(lists[0], lists[1])
	}

	return mergeList(mergeKLists(lists[0:len(lists)/2]), mergeKLists(lists[len(lists)/2:]))

}

func mergeList(l1, l2 *common.ListNode) *common.ListNode {
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

func main() {
	l1 := common.NewList(1, 5, 9, 13, 17)
	l2 := common.NewList(2, 6, 10, 14, 18)
	l3 := common.NewList(3, 7, 11, 15, 19)
	l4 := common.NewList(4, 8, 12, 16, 20)
	l := []*common.ListNode{l1, l2, l3, l4}
	common.PrintListNode(mergeKLists(l))
}
