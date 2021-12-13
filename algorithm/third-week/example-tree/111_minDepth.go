package example_tree

import (
	"fmt"
	"geek-study/algorithm/common"
)

func minDepth(root *common.TreeNode) int {
	return getMinDepth(root)
}

func getMinDepth(root *common.TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil {
		return getMinDepth(root.Right) + 1
	}
	if root.Right == nil {
		return getMinDepth(root.Left) + 1
	}
	return min(getMinDepth(root.Left), getMinDepth(root.Right)) + 1
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	tree := common.NewTree(1, 2, nil, 3, nil)
	common.PrintTree(tree)
	fmt.Println(minDepth(tree))
}
