package tree

import (
	"fmt"
	"geek-study/algorithm/common"
)

func isSameTree(p *common.TreeNode, q *common.TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if (p == nil && q != nil) || (p != nil && q == nil) || p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func main() {
	root := common.NewTree(1, 2, 3)
	root1 := common.NewTree(1, 2, 3, 5)
	common.PrintTree(root)
	fmt.Println(isSameTree(root, root1))
}
