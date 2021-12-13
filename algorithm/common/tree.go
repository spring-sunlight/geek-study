package common

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(val int, left, right *TreeNode) *TreeNode {
	return &TreeNode{
		Val:   val,
		Left:  left,
		Right: right,
	}
}

func NewTree(val ...interface{}) *TreeNode {
	return buildTree(0, val)

}

func buildTree(i int, val []interface{}) *TreeNode {
	if val[i] == nil {
		return nil
	}
	var node = NewTreeNode(val[i].(int), nil, nil)
	if i < len(val) && 2*i+1 < len(val) {
		node.Left = buildTree(2*i+1, val)
	}
	if i < len(val) && 2*i+2 < len(val) {
		node.Right = buildTree(2*i+2, val)
	}
	return node
}

func PrintTree(tree *TreeNode) {
	nodes := []*TreeNode{tree}
	dep := 2
	fmt.Println("===============================")
	fmt.Println("第", 1, "层:[", nodes[0].Val, "]")
	for len(nodes) != 0 {
		s := ""
		temp := []*TreeNode{}
		for i := 0; i < len(nodes); i++ {

			if nodes[i].Left != nil {
				temp = append(temp, nodes[i].Left)
				s += strconv.Itoa(nodes[i].Left.Val) + ", "

			} else {
				s += "nil, "
			}
			if nodes[i].Right != nil {
				temp = append(temp, nodes[i].Right)
				s += strconv.Itoa(nodes[i].Right.Val) + ", "
			} else {
				s += "nil, "
			}
		}
		if len(temp) != 0 {
			fmt.Println("第", dep, "层:[", s[:len(s)-2], "]")
		}
		dep++
		nodes = temp
	}
	fmt.Println("===============================")
}
