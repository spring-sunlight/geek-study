package tree

import (
	"fmt"
	"geek-study/algorithm/common"
)

func inorderTraversal(root *common.TreeNode) []int {
	ans := []int{}
	var inorder func(root *common.TreeNode)
	inorder = func(root *common.TreeNode) {
		if root == nil {
			return
		}
		inorder(root.Left)
		ans = append(ans, root.Val)
		inorder(root.Right)
	}
	inorder(root)
	return ans
}

func main() {
	root := common.NewTree(1, nil, 2, nil, nil, nil, 3)
	common.PrintTree(root)
	fmt.Println(inorderTraversal(root))
}
