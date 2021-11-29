package example

import (
	"geek-study/algorithm/common"
)

func reverseKGroup(head *common.ListNode, k int) *common.ListNode {
	protected := common.NewListNode(0, head)
	var pre = protected
	for head != nil {
		//1.分组:每组 k 个
		end := getEnd(head, k)
		if end == nil {
			break
		}
		//2.将每组进行翻转,从 head 到 nextGroupHead 处停止
		nextGroupHead := end.Next
		//reverse 里只有对 head.Next 操作才会改变 外部head 的值, head = nextHead 只会改变函数内部 head 的值
		//reverse 前,head:1->2->3->4->5,end:3->4->5
		//reverse 后,head:1,end:3->2->1
		reverse(head, nextGroupHead)

		//3.将每组连接
		pre.Next = end
		head.Next = nextGroupHead

		pre = head
		head = nextGroupHead
		common.PrintListNode(protected, "protected")
	}
	return protected.Next
}

func getEnd(head *common.ListNode, k int) *common.ListNode {
	for head != nil {
		k--
		if k == 0 {
			return head
		}
		head = head.Next
	}
	return nil
}

func reverse(head *common.ListNode, stop *common.ListNode) {
	var pre, nextHead *common.ListNode
	for head != stop {
		nextHead = head.Next
		head.Next = pre
		pre = head
		head = nextHead
	}
}

func TestReverseKGroup() {
	l := common.NewList(1, 2, 3, 4, 5)
	reverseKGroup(l, 3)
}
