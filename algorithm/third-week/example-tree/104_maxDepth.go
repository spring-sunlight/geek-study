package example_tree

import (
	"fmt"
	"geek-study/algorithm/common"
)

func maxDepth(root *common.TreeNode) int {
	return getMaxDepth(root)
}

func getMaxDepth(root *common.TreeNode) int {
	if root == nil {
		return 0
	}
	return max(getMaxDepth(root.Left), getMaxDepth(root.Right)) + 1
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func test() {
	tree := common.NewTree(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	common.PrintTree(tree)
	fmt.Println(maxDepth(tree))

}
