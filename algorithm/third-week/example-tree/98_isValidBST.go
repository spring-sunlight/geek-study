package example_tree

import "geek-study/algorithm/common"

func isValidBST(root *common.TreeNode) bool {
	return check(root, -1<<63, 1<<63-1)
}

func check(root *common.TreeNode, rangeLeft, rangeRight int) bool {
	if root == nil {
		return true
	}
	if root.Val <= rangeLeft || root.Val >= rangeRight {
		return false
	}
	return check(root.Left, rangeLeft, root.Val) && check(root.Right, root.Val, rangeRight)
}
