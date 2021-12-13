package example_tree

import "geek-study/algorithm/common"

func invertTree(root *common.TreeNode) *common.TreeNode {
	invert(root)
	return root
}

func invert(root *common.TreeNode) {
	if root == nil {
		return
	}
	root.Left, root.Right = root.Right, root.Left
	invert(root.Left)
	invert(root.Right)
}

func main() {
	tree := common.NewTree(4, 2, 7, 1, 3, 6, 9)
	common.PrintTree(tree)
	common.PrintTree(invertTree(tree))
}
