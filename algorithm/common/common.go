package common

import (
	"fmt"
	"strconv"
)

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

func PrintListNode(node *ListNode, name ...string) {
	counter := 0
	str := ""
	if len(name) == 0 {
		name = append(name, " ")
	}
	for node != nil {
		str += strconv.Itoa(node.Val) + "->"
		node = node.Next
		counter++
		if counter > 20 {
			break
		}
	}
	if str != "" {
		fmt.Println("链表", name[0], "为: "+str[:len(str)-2])
	} else {
		fmt.Println("当前链表为nil ")
	}

}
